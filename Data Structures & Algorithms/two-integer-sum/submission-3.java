class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer,Integer> seen = new HashMap<>();

		for(int i=0;i<nums.length;i++){
			int cpt = target-nums[i];
			if(seen.containsKey(cpt)){
				return new int[]{seen.get(cpt),i};
			}
			seen.put(nums[i],i);
		}

		return new int[]{};
    }
}
