func longestCommonPrefix(strs []string) string {
     if len(strs) == 0 {
        return ""
    }
    res :=""
   sort.Strings(strs)
   first := strs[0]
   last := strs[len(strs)-1]
   
   for i := 0; i < len(first); i++{
    if first[i]!=last[i]{
        break;
    }
    res +=string(first[i])
   }
   return res
}
