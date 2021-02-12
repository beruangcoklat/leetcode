class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int a = nums[i];
            Integer idx = map.get(target - a);
            if (idx != null) {
                return new int[]{idx, i};
            }
            map.put(a, i);
        }
        return null;
    }
}