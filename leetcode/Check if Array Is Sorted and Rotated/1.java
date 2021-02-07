class Solution {
    public boolean check(int[] nums) {
        if (nums.length <= 1) return true;

        int min = nums[0];
        int minIdx = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < min) {
                min = nums[i];
                minIdx = i;
            }
        }

        for (int i = 0; i < nums.length - 1; i++) {
            int idx = (minIdx + i) % nums.length;
            int nextIdx = (minIdx + i + 1) % nums.length;
            if (nums[idx] > nums[nextIdx] && nums[nextIdx] != min) return false;
        }

        return true;
    }
}
