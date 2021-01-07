class Solution {
    public double solve(double x, int n) {
        if (n == 1) {
            return x;
        }

        double ret = solve(x, n / 2);
        ret *= ret;

        if (n % 2 == 1) {
            ret *= x;
        }

        return ret;
    }

    public double myPow(double x, int n) {
        if (n == 0) {
            return 1;
        }
        if (n > 0) {
            return solve(x, n);
        }

        if (n == Integer.MIN_VALUE) {
            return 1 / (solve(x, Integer.MAX_VALUE) * x);
        }

        return 1 / solve(x, -n);
    }
}
