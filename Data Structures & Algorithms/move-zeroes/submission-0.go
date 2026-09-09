func moveZeroes(nums []int) {
	insertPos:=0;
	for i,_ := range nums{
		if nums[i]!=0{
			nums[insertPos],nums[i] = nums[i],nums[insertPos]
			insertPos+=1;
		}
	}
}
