class Solution {
    //版本一 会超出内存限制
    public int maxProduct(int[] nums) {
        int value = Integer.MIN_VALUE;
        int[][] dp = new int[nums.length][nums.length];
        for (int i = 0; i < nums.length; i++) {
            dp[i][i] = nums[i];
        }
        for (int i = 1; i <nums.length ; i++) {
            for (int j = 0; j <=i ; j++) {
                dp[i][j] = dp[i-1][j]*nums[j];
                value = Math.max(dp[i][j],value);
            }
        }
        return value;
    }
}