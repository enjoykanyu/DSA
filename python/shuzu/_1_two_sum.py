class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = {}
        for index,cur in enumerate(nums):
            if target-cur in map:
                return [index,map[target-cur]]
            map[cur] = index # 若存在相同的数值 得先判断才存储map，否则比如[3,3]会 存入了{3:0}
            #判断target=6 3存在 直接返回了0 0
        return []

