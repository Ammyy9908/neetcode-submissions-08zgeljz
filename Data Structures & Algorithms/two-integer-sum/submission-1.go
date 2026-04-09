func twoSum(nums []int, target int) []int {
    
mp := make(map[int]int)

for i,j := range nums {
	diff := target-j;
	if val,ok := mp[j];ok{
		return []int{val,i}
	}
	mp[diff] = i;
}

return nil

}
