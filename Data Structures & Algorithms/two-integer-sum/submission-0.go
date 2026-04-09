func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)

    for i,v := range nums{
        diff := target-v
        index, ok := hashmap[diff]
    if ok {
    return []int{index, i}
    }
        hashmap[v]=i
    }

    return nil;


}
