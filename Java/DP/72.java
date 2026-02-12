class Solution {
    public int minDistance(String word1, String word2) {
        char[] l = word1.toCharArray();
        char[] r = word2.toCharArray();
        int len1 = l.length;
        int len2 = r.length;
        int memo[][] = new int[len1][len2];
        //赋值-1未算过
        for (int[] cur: memo) {
            Arrays.fill(cur,-1);
        }
        return dfs(len1-1,len2-1,memo,l,r);
    }
    public int dfs(int l,int r,int[][] memo,char[] charL,char[] charR){
        if (l<0){
            return  r+1;
        }//l字符数量=0则需插入r数量+1
        if (r<0){
            return l+1;
        }
        if (memo[l][r]!=-1){//算过了
            return memo[l][r];
        }
        if (charL[l]==charR[r]){//两个结尾字符相同
            memo[l][r]=dfs(l-1,r-1,memo,charL,charR);
            return memo[l][r];
        }
        memo[l][r] = Math.min(Math.min(dfs(l,r-1,memo,charL,charR),dfs(l-1,r,memo,charL,charR)),dfs(l-1,r-1,memo,charL,charR))+1;
        return memo[l][r];
    }
}