class Solution {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (l, t) -> l[0] - t[0]); //左端点排序
        List<int[]> value = new ArrayList<>();//单个数组
        for (int i = 0; i < intervals.length; i++) {
            int len = value.size(); //当前合并后的大小
            if(len >0 && value.get(len-1)[1] >= intervals[i][0]){
                value.get(len-1)[1] = Math.max(intervals[i][1],value.get(len-1)[1]);//更新右端节点
            }else {//第一次
                value.add(intervals[i]);
            }
        }
        return value.toArray(new int[value.size()][]);
    }
}