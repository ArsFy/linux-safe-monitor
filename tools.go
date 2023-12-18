package main

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
