class Solution {
    public int rob(int[] nums) {
        int len = nums.length;
        int[] account = new int[len];//计算没有计算过的金额
        Arrays.fill(account, -1);
        return dfs(nums, len-1, account);
    }
    public int dfs(int[] nums,int l, int[] account){
        if (l<0){
            return 0;
        }
        if (account[l]!=-1){//计算过
            return account[l];
        }
        int noChoose = dfs(nums,l-1,account);
        int choose = dfs(nums,l-2,account)+nums[l];
        account[l]=Math.max(choose,noChoose);//记忆化存储
        return account[l];
    }
}