class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        myhashmap = {}
        for s in strs:
            key = ''.join(sorted(s))
            if key in myhashmap:
                myhashmap[key].append(s)
            else:
                myhashmap[key]=[s]
        
        return list(myhashmap.values())