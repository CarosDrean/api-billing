package model

func UniqueUints(list []uint) []uint {
	keys := make(map[uint]struct{})
	var uniqueList []uint

	for _, item := range list {
		if _, ok := keys[item]; !ok {
			keys[item] = struct{}{}
			uniqueList = append(uniqueList, item)
		}
	}

	return uniqueList
}
