func topKFrequent(nums []int, k int) []int {
    hashMap := make(map[int]int)
    for _,v := range nums {
        hashMap[v]++
    }
    type kv struct {
        Key int
        Value int
    }

    var pairs []kv
    for k,v := range hashMap{
        pairs = append(pairs,kv{k,v})
    }
    sort.Slice(pairs,func(i,j int)bool{
        return pairs[i].Value > pairs[j].Value
    })

    var result []int
    for i:=0;i<k;i++{
        result = append(result,pairs[i].Key);
    }
    return result;
}
