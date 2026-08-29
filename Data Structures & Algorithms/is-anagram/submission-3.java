class Solution {
    public boolean isAnagram(String s, String t) {
		if (s.length()!=t.length()){
			return false;
		}

		Map<Character,Integer> maps= new HashMap<>();
		Map<Character,Integer> mapt = new HashMap<>();

		for(char ch: s.toCharArray()){
			maps.put(ch,maps.getOrDefault(ch,0)+1);
		}

		for(char ch: t.toCharArray()){
			mapt.put(ch,mapt.getOrDefault(ch,0)+1);
		}

		for(char key:maps.keySet()){
			if(!maps.get(key).equals(mapt.get(key))){
				return false;
			}
		}

		return true;
    }
}
