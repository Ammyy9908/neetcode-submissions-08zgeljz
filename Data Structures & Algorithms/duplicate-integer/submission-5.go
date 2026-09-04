func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _,v := range nums {
        if ok,_ := seen[v];ok{
            return true;
        }
        seen[v]=true;
    }
    return false;
}
