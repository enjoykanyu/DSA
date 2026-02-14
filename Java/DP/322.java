class Solution {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount+1];//凑成当前这个金额的最小硬币数量
        Arrays.fill(dp,amount+1);//最小硬币数量最多=amount amount+1表示默认可以凑出来 但当amount+1 即经过计算dp[amount]的值肯定会经过修改，当没有被修改则无法凑出判断为-1 且需由最多推倒出来 理论上这里可以设置int上线也可以设置amount+1
        dp[0] = 0; //凑成0元无需硬币
        //从硬币数组从左到右遍历判断硬币数量是否<= 组成金额数量 符合则判断凑齐数量最小值 凑齐小数额的dp算过了可以由这里递推过来
        for (int c = 1; c <=amount ; c++) {//遍历硬币金额
            for (int i = 0; i <coins.length ; i++) {
                if (coins[i]<=c){
                    dp[c] = Math.min(dp[c],dp[c-coins[i]]+1);
                }
            }
        }
        return dp[amount] > amount ?-1:dp[amount];
    }
}