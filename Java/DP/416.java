class Solution {
    public boolean canPartition(int[] nums) {
        int count =0;
        for (int i = 0; i <nums.length ; i++) {
            count+=nums[i];
        }
        if (count%2!=0){//奇数无法拆分两组数
            return false;
        }
        count /=2;
        boolean[][] dp = new boolean[nums.length+1][count+1];
        dp[0][0]=true;
        for (int i = 0; i <nums.length ; i++) {
            int cur = nums[i];
            for (int j = 0; j <=count ; j++) {
                dp[i+1][j] = j>=cur && dp[i][j-cur] || dp[i][j];
            }
        }
        return dp[nums.length][count];
    }
}