func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i,v := range nums{
		cpt := target-v
		if j,ok:= seen[cpt]; ok{
			return []int{j,i}
		}
		seen[v]=i
	}

	return nil;
}
