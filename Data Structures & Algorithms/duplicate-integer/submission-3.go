func hasDuplicate(nums []int) bool {
    mymap := make(map[int]bool)
	for _,value := range nums{
		_,exist := mymap[value]
		if exist{
			return true
		}
		mymap[value] = true
	}

	return false
}
