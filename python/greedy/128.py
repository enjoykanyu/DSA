class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums:
            return 0
        count = 1
        result = 0
        map = {}
        # hashset 存储所有元素
        for num in nums:
            map[num] = True

        for num in map:
            if num-1 in map:
                continue
            cur = num+1
            while cur in map:
                cur += 1
                count += 1
            result = max(result,count)
            count = 1

        return result