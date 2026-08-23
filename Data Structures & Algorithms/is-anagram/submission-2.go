func isAnagram(s string, t string) bool {
      if len(s)!=len(t){
		return false;
	  }

	  maps := make(map[rune]int)
	  mapt := make(map[rune]int)
	  for _,v := range []rune(s){
		maps[v]++
	  }

	  for _,v := range []rune(t){
		mapt[v]++
	  }

	  for key,_ := range maps{
			if maps[key]!=mapt[key]{
				return false;
			}
	  }

	  return true
}
