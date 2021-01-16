class Solution {
    public int getMaximumGenerated(int n) {
        if (n < 2) return n;

        int arr[] = new int[n + 1];
        arr[0] = 0;
        arr[1] = 1;
        int max = Integer.MIN_VALUE;

        for (int i = 2; i < n + 1; i++) {
            if (i % 2 == 0) {
                arr[i] = arr[i / 2];
            } else {
                int a = (i - 1) / 2;
                int b = a + 1;
                arr[i] = arr[a] + arr[b];
            }

            if (arr[i] > max) {
                max = arr[i];
            }
        }

        return max;
    }
}
