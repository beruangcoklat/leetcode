class Solution {
	public int countKDifference(int[] nums, int k) {
		Map<Integer, Integer> m = new HashMap<>();
		int result = 0;
		for (int i = 0; i < nums.length; i++) {
			int diff = Math.abs(k - nums[i]);
			int diff2 = Math.abs(k + nums[i]);

			if (Math.abs(nums[i] - diff) == k) {
				int count = m.getOrDefault(diff, 0);
				result += count;
			}

			if (Math.abs(nums[i] - diff2) == k) {
				int count = m.getOrDefault(diff2, 0);
				result += count;
			}
			
			m.put(nums[i], m.getOrDefault(nums[i], 0) + 1);
		}
		return result;
	}
}