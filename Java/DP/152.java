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
//优化的版本 采用动态规划以右端点
class Solution {
    public int maxProduct3(int[] nums) {
        int dpMax[] = new int[nums.length];
        int dpMin[] = new int[nums.length];

        dpMax[0] = nums[0];
        dpMin[0] = nums[0];
        int value = Integer.MIN_VALUE;
        for (int i = 1; i <nums.length ; i++) {
            dpMax[i] = Math.max(Math.max(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i]),nums[i]);
            dpMin[i] = Math.min(Math.min(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i]),nums[i]);
            value = Math.max(value,dpMax[i]);
        }
        return Math.max(value,nums[0]);
    }
}