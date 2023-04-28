package tool

func MergeArrayStr(arr1 []string, arr2 []string) []string {
	arr1 = append(arr1, arr2...)
	return arr1
}

func UniqueArrStr(arr []string) []string {
	list := make([]string, 0)
	temp := map[string]string{}
	for _, val := range arr {
		if _, ok := temp[val]; !ok {
			temp[val] = ""
			list = append(list, val)
		}
	}
	return list
}
