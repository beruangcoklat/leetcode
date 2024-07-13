class Solution {
    public int maximumUniqueSubarray(int[] nums) {
        int l = 0;
        int r = 0;
        int total = nums[0];
        Map<Integer, Integer> m = new HashMap<>();
        m.put(nums[0], 1);
        int max = total;
        boolean isValid = true;
        int invalidNum = 0;

        while (l < nums.length && r < nums.length) {
            if (!isValid) {
                int c = m.getOrDefault(nums[l], 0);
                if (c == 1) {
                    total -= nums[l];
                }
                m.put(nums[l], c - 1);
                if (nums[l] == invalidNum) {
                    isValid = true;
                }
                l++;
                continue;
            }

            r++;
            if (r >= nums.length) {
                break;
            }

            int c = m.getOrDefault(nums[r], 0);
            if (c == 0) {
                total += nums[r];
            } else {
                isValid = false;
                invalidNum = nums[r];
            }
            m.put(nums[r], c + 1);
            max = Math.max(max, total);
        }

        return max;
    }
}