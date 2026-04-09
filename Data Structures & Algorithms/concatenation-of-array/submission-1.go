func getConcatenation(nums []int) []int {
    ans := []int{}
    for i:=0;i<2;i++{
        for _,v := range nums{
            ans = append(ans,v)
        }
    }
    return ans;
}
