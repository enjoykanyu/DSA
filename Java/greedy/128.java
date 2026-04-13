class Solution {
    public int longestConsecutive(int[] nums) {
        int result = 0;
        int count = 1;
        HashSet<Integer> cur = new HashSet<>();
        for (int num: nums) {
            cur.add(num);
        }
        for (Integer value: cur) {
            if (cur.contains(value-1)){
                continue;
            }
            int cur_value = value+1;
            while (cur.contains(cur_value)){
                cur_value++;
                count++;
            }
            if (count>result){
                result=count;
            }
            count=1;
        }
        return nums.length==0? 0:Math.max(count,result);
    }
}