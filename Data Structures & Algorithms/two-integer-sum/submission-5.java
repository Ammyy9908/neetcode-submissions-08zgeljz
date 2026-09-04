class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer,Integer> mymap = new HashMap<>();

        for (int i=0;i<nums.length;i++){
            int cpt = target-nums[i];
            boolean exist = mymap.containsKey(cpt);
            if (exist){
                return new int[]{mymap.get(cpt),i};
            }
            mymap.put(nums[i],i);
        }

        return new int[]{};
    }
}
