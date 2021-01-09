class Solution {

    void remove(int[] arr, int idx, int len) {
        for (int i = idx; i < len - 1; i++) {
            arr[i] = arr[i + 1];
        }
    }

    public int removeElement(int[] nums, int val) {
        int idx = 0;
        int len = nums.length;
        while (idx < len) {
            if (nums[idx] == val) {
                remove(nums, idx, len);
                len--;
            } else {
                idx++;
            }
        }
        return len;
    }
}
