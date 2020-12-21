class Solution {
    public int findKthPositive(int[] arr, int k) {
        int prev = 0;
        int totalRange = 0;
        for (int i = 0; i < arr.length; i++) {
            int range = arr[i] - prev - 1;
            if (range == 0) {
                prev = arr[i];
                continue;
            }

            if (totalRange + range >= k) {
                return prev + (k - totalRange);
            }
            totalRange += range;
            prev = arr[i];
        }

        if (totalRange == 0) {
            return arr[arr.length - 1] + k;
        }

        return arr[arr.length - 1] + k - totalRange;
    }
}
