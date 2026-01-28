class Solution {
    public int lengthOfLongestSubstring(String s) {
        char[] chars = s.toCharArray();//字符串
        int result = 0;//最大值
        int left = 0;//窗口左边
        int[] cnt = new int[128]; //字符计数
        for (int right = 0; right < chars.length; right++) {
            int t = chars[right];
            cnt[t]++;
            while (cnt[t] > 1) {//有重复 窗口左边-1
                int l = chars[left];
                cnt[l]--;
                left++;
            }
            result = Math.max(result, right - left + 1);
        }
        return result;
    }
}