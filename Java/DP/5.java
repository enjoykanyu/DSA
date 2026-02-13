class Solution {
    public String longestPalindrome(String s) {
        boolean dp[][] = new boolean[s.length()][s.length()];
        for (int i = 0; i <s.length(); i++) {
                dp[i][i] = true;
        }
        int begin =0;
        int value =1;
        //从长度=2开始动态递推
        for (int L = 2; L <=s.length(); L++) {
            for (int i = 0; i <s.length(); i++) {
                int j =i+L-1; //右端点
                if (j>=s.length()){
                    break;
                }
                if (s.charAt(i)!=s.charAt(j)){
                    dp[i][j] = false;
                }else {
                    if (j-i<3){//偶数
                        dp[i][j] = true; //比如lbbcltf
                    }else {
                        dp[i][j] = dp[i+1][j-1];
                    }
                }
                if (dp[i][j] && j-i+1>value){
                    value  = j-i+1;
                    begin = i;
                }
            }
        }
        return s.substring(begin,begin+value);
    }
}