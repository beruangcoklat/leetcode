class Solution {
    public long countSubarrays(int[] nums, int minK, int maxK) {
        List<Integer> list = new ArrayList<>();
        long result = 0;
        for (int i = 0; i < nums.length; i++) {
            int curr = nums[i];
            if (curr < minK || curr > maxK) {
                if (!list.isEmpty()) {
                    result += solve(list, minK, maxK);
                    list.clear();
                }
                continue;
            }
            list.add(curr);
        }
        if (!list.isEmpty()) {
            result += solve(list, minK, maxK);
        }
        return result;
    }

    public long solve(List<Integer> nums, int minK, int maxK) {
        int l = 0;
        int r = 0;
        long result = 0;

        Map<Integer, Integer> m = new HashMap<>();
        m.put(nums.get(l), 1);

        while (true) {
            if (valid(minK, maxK, m)) {
                result += nums.size() - r;
                int c = m.getOrDefault(nums.get(l), 0);
                if (c > 0) {
                    m.put(nums.get(l), c - 1);
                }
                l++;
                continue;
            }

            r++;
            if (r >= nums.size()) {
                break;
            }
            m.put(nums.get(r), m.getOrDefault(nums.get(r), 0) + 1);
        }
        return result;
    }

    private boolean valid(int minK, int maxK, Map<Integer, Integer> m) {
        return m.getOrDefault(minK, 0) > 0 &&
            m.getOrDefault(maxK, 0) > 0;
    }
}
