func topKFrequent(nums []int, k int) []int {
hashmap := make(map[int]int)

for _, num := range nums{
	hashmap[num]++
}

res := []int{}
for num := range hashmap{
	res = append(res,num)
}
sort.Slice(res, func(i, j int) bool {
        return hashmap[res[i]] > hashmap[res[j]]
    })
return res[:k]
}
