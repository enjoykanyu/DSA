class Solution {
    public int numSquares(int n) {
        int[] dp = new int[n+1];//构成当前数字最少需的完全平方数量
        for (int i = 1; i <=n ; i++) {//遍历递推凑成从1到n所需的数量
            dp[i] = Integer.MAX_VALUE;
            for (int j = 1; j*j<=i ; j++) {
                dp[i] = Math.min(dp[i-j*j]+1,dp[i]);
            }
        }
        return dp[n];
    }
}