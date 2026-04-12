class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] result = new int[2];
        Map<Integer,Integer> map_cur = new HashMap<>();
        for (int i = 0; i <nums.length ; i++) {
            int cur= target-nums[i];
            if (map_cur.containsKey(cur)){
                result[0] = i;
                result[1] = map_cur.get(cur);
            }
            map_cur.put(nums[i],i);
        }
        return result;

    }
}