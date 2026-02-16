class Solution {
    public int longestCommonSubsequence(String text1, String text2) {
        char[] l = text1.toCharArray();
        char[] r = text2.toCharArray();
        int[][] dp = new int[l.length+1][r.length+1];
        for (int i = 0; i <l.length ; i++) {
            for (int j = 0; j <r.length ; j++) {
                if (l[i]==r[j]){
                    dp[i+1][j+1] = dp[i][j]+1;
                }else {
                    dp[i+1][j+1] = Math.max(dp[i][j+1],dp[i+1][j]);
                }
            }
        }
        return dp[l.length][r.length];
    }
}