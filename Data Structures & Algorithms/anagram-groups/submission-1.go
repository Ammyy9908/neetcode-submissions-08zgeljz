func groupAnagrams(strs []string) [][]string {
        mymap := make(map[string][]string)

        for _,s := range strs{
            key:= sortString(s)
            mymap[key] = append(mymap[key],s)
        }

        result:= make([][]string,0,len(mymap))

        for _,value:= range mymap{
            result = append(result,value)
        }

        return result;
}


func sortString(s string) string{
    bytes := []byte(s)
    sort.Slice(bytes,func(i,j int)bool{
        return bytes[i] < bytes[j]
    })

    return string(bytes)
}
