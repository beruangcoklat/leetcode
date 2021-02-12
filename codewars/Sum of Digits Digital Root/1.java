public class DRoot {
    public static int getSum(int n) {
        int i = 10;
        int ret = 0;
        while (i / 10 <= n) {
            int digit = (n % i) / (i / 10);
            ret += digit;
            i *= 10;
        }
        return ret;
    }

    public static int digital_root(int n) {
        while (n >= 10){
            n = getSum(n);
        }
        return n;
    }
}