func hasDuplicate(nums []int) bool {
    hashmap := make(map[int]bool)
   

    for _, v := range nums{
        key := hashmap[v]
        if key{
            return true;
        }
        hashmap[v]=true;
    }

    return false;
}
