class Solution {
    public int minDays(int[] bloomDay, int m, int k) {
        if (bloomDay.length < ((long) m * (long) k)) {
            return -1;
        }

        int minDay = bloomDay[0];
        int maxDay = bloomDay[1];
        for (int day : bloomDay) {
            minDay = Math.min(minDay, day);
            maxDay = Math.max(maxDay, day);
        }

        int l = minDay;
        int r = maxDay;
        int last = 0;
        while (l <= r) {
            int mid = (l + r) / 2;
            if (valid(bloomDay, mid, m, k)) {
                r = mid - 1;
                last = mid;
            } else {
                l = mid + 1;
            }
        }

        return last;
    }

    private boolean valid(int[] bloomDay, int day, int m, int k) {
        for (int i = 0; i < bloomDay.length; i++) {
            if (i + k > bloomDay.length) {
                break;
            }

            boolean valid = true;
            for (int j = i; j < i + k; j++) {
                if (bloomDay[j] > day) {
                    valid = false;
                    break;
                }
            }

            if (valid) {
                m--;
                i += k - 1;
            }
            if (m == 0) {
                return true;
            }
        }
        return false;
    }
}