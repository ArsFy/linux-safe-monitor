package main

import "strings"

func isStringArrayNotContains(str string, arr []string) bool {
	for _, item := range arr {
		if strings.Contains(item, str) {
			return false
		}
	}
	return true
}

func containsString(arr *[]string, target string) bool {
	for _, str := range *arr {
		if str == target {
			return true
		}
	}
	return false
}

func compareMaps(map1, map2 map[string]string) (addedKeys, removedKeys []string) {
	for key := range map1 {
		if _, ok := map2[key]; !ok {
			removedKeys = append(removedKeys, key)
		}
	}

	for key := range map2 {
		if _, ok := map1[key]; !ok {
			addedKeys = append(addedKeys, key)
		}
	}

	return addedKeys, removedKeys
}
