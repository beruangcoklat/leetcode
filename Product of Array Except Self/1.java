class Solution {
    public int[] productExceptSelf(int[] nums) {
        int left[] = new int[nums.length];
        int right[] = new int[nums.length];
        int ret[] = new int[nums.length];

        int curr = 1;
        left[0] = 1;
        for (int i = 1; i < nums.length; i++) {
            curr *= nums[i - 1];
            left[i] = curr;
        }
        
        curr = 1;
        right[nums.length - 1] = 1;
        for (int i = nums.length - 2; i >= 0; i--) {
            curr *= nums[i + 1];
            right[i] = curr;
        }

        for (int i = 0; i < nums.length; i++) {
            ret[i] = left[i] * right[i];
        }

        return ret;
    }
}
