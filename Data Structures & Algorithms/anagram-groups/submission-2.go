func groupAnagrams(strs []string) [][]string {
        mymap := make(map[[26]int][]string)

        for _,s := range strs{
            count := [26]int{}
            for _,c := range s{
                count[c-'a']++
            }

            mymap[count] = append(mymap[count],s)
        }

        result := [][]string{}

        for _,v := range mymap{
            result = append(result,v)
        }

        return result;
}


// func sortString(s string) string{
//     bytes := []byte(s)
//     sort.Slice(bytes,func(i,j int)bool{
//         return bytes[i] < bytes[j]
//     })

//     return string(bytes)
// }
