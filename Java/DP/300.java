class Solution {
    public int lengthOfLIS(int[] nums) {
        int len = nums.length;
        int[] dp = new int[len];
        Arrays.fill(dp,1);//以当前结尾最长序列长度
        int value = 0;
        for (int i = 0; i <nums.length ; i++) {
            for (int j = 0; j <i ; j++) {
                if (nums[j]<nums[i]){
                    dp[i] = Math.max(dp[i],dp[j]+1);
                }
            }
            value = Math.max(value,dp[i]);
        }
        return value;
    }
}