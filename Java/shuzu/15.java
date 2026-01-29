class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
//排序
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();
        int len = nums.length;
        for (int i = 0; i < len - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;//重复
            if(nums[i]+nums[i+1]+nums[i+2] >0){
                break;
            }
            if (nums[i] + nums[len-1] + nums[len-2] < 0){
                continue;
            }
            int l=i+1;
            int r=len-1;
            while (l<r){
                int t = nums[i] + nums[l] + nums[r];//总和
                if (t == 0){
                    result.add(List.of(nums[i], nums[l], nums[r]));
                    for (l++; l < r && nums[l] == nums[l - 1]; l++);
                    for (r--; r > l && nums[r] == nums[r + 1]; r--);
                }else if (t > 0){
                    r--;
                }else {
                    l++;
                }
            }
        }
        return result;
    }
}