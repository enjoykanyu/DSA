class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        //枚举最大长度计算
        int maxLen = 0;
        for (String word:wordDict) {
            maxLen = Math.max(maxLen,word.length());
        }
        //新建单词哈希集合用于判断遍历过程是否存在于单词中
        Set<String> word = new HashSet<>(wordDict);
        //新建记忆化数组计算是否计算过
        int n = s.length();
        int[] memo = new int[n+1];
        Arrays.fill(memo,-1);
        return dfs(maxLen,s,word,n,memo)==1;
    }
    public int dfs(int maxLen,String s,Set<String> word,int n, int[] memo){
        //临界跳出循环
        if (n==0){
            return 1;
        }
        //计算过
        if (memo[n] !=-1){
            return memo[n];
        }
        //dp
        for (int l = n-1; l >=Math.max(n-maxLen,0) ; l--) {
            if (word.contains(s.substring(l,n)) && dfs(maxLen,s,word,l,memo) ==1){
                memo[n] = 1;
                return memo[n];
            }
        }
        memo[n] = 0;//当前计算过但不符合
        return memo[n];
    }
}