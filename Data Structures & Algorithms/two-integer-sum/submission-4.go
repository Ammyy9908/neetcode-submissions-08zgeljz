func twoSum(nums []int, target int) []int {
    mymap := make(map[int]int);

    for i,v := range nums{
        cpt := target-v;
        if _,value := mymap[cpt]; value{
            return []int{mymap[cpt],i};
        }
        mymap[v]=i;
    }

    return []int{}
}
