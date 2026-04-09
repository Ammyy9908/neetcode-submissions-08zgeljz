func majorityElement(nums []int) int {

    hashmap := make(map[int]int)
    for _,v := range nums{
        hashmap[v]++
    }

    for v,count := range hashmap {
       if count>(len(nums)/2){
        return v;
       }
    }

    return 0;
    
}
