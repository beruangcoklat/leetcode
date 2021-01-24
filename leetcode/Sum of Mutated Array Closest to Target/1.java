class Solution {
    int getMax(int[] arr) {
        int max = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
        }
        return max;
    }

    public int findBestValue(int[] arr, int target) {
        Arrays.sort(arr);

        int left = 0;
        int right = getMax(arr);

        int minDiff = Integer.MAX_VALUE;
        int minSum = Integer.MAX_VALUE;
        int bestValue = -1;

        while (left <= right) {
            int mid = (left + right) / 2;

            int currSum = 0;
            for (int a : arr) {
                if (a > mid) currSum += mid;
                else currSum += a;
            }

            int diff = Math.abs(target - currSum);
            if (diff < minDiff) {
                minDiff = diff;
                minSum = currSum;
                bestValue = mid;
            } else if (diff == minDiff && currSum < minSum) {
                minSum = currSum;
                bestValue = mid;
            }

            if (currSum < target) left = mid + 1;
            else right = mid - 1;
        }

        return bestValue;
    }
}