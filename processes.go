package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

var ProcessMap = make(map[string]string)

func getAllProcesses() error {
	var processes = make(map[string]string)
	var wg sync.WaitGroup
	files, err := os.ReadDir("/proc")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			pid := file.Name()
			if _, err := strconv.Atoi(pid); err == nil {
				wg.Add(1)
				go func(pid string) {
					defer wg.Done()
					stat, err := os.ReadFile(filepath.Join("/proc", pid, "status"))
					if err != nil {
						log.Fatal(pid, " Cannot read status: ", err)
					}
					statusLines := strings.Split(string(stat), "\n")
					name := strings.Replace(statusLines[0], "Name:\t", "", 1)
					user, err := user.LookupId(strings.Split(strings.Replace(statusLines[8], "Uid:\t", "", 1), "\t")[0])
					if err != nil {
						log.Fatal(pid, " Cannot lookup user id: ", err)
					}
					cmdline, err := os.ReadFile(filepath.Join("/proc", pid, "cmdline"))
					if err != nil {
						log.Fatal(pid, " Cannot read cmdline: ", err)
					}

					mu.Lock()
					processes[user.Username+" "+name+" "+string(cmdline)] = pid
					mu.Unlock()
				}(pid)
			}
		}
	}

	wg.Wait()

	if len(ProcessMap) != 0 {
		add, del := compareMaps(ProcessMap, processes)
		for _, key := range add {
			fmt.Println("Added:", key, processes[key])

			switch Config.KillMode {
			case 1:
				if !containsString(&WhiteList, strings.Split(key, " ")[1]) {
					if err := exec.Command("kill", "-9", processes[key]).Run(); err != nil {
						log.Println("Failed to kill process: ", err)
						return err
					}
					delete(processes, key)
				}
			case 2:
				if err := exec.Command("kill", "-9", processes[key]).Run(); err != nil {
					log.Println("Failed to kill process: ", err)
					return err
				}
				delete(processes, key)
			}

			pInfo := strings.Split(key, " ")
			go push(1, pInfo[1], key, processes[key], pInfo[0])
		}
		for _, key := range del {
			fmt.Println("Removed:", key, ProcessMap[key])

			pInfo := strings.Split(key, " ")
			go push(0, pInfo[1], key, processes[key], pInfo[0])
		}
	}

	ProcessMap = processes

	return nil
}
