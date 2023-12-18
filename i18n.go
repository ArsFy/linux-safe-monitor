package main

var i18n = map[string]map[string]string{
	"en": {
		"process_running": "New Process %s is running\nCmdline: %s\nUser: %s\nPid: %s",
		"process_killed":  "Process %s is killed\nCmdline: %s\nUser: %s\nPid: %s",
	},
	"cn": {
		"process_running": "新进程 %s 正在运行\n命令行: %s\n用户: %s\nPid: %s",
		"process_killed":  "进程 %s 被杀死\n命令行: %s\n用户: %s\nPid: %s",
	},
}
