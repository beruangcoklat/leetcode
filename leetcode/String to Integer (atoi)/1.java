class Solution {
    public int myAtoi(String s) {
        boolean start = false;
        boolean isMin = false;
        long ret = 0;

        for (char c : s.toCharArray()) {
            if (!start) {
                if (c == '-') {
                    isMin = true;
                    start = true;
                } else if (c == '+') {
                    start = true;
                } else if (c == ' ') {
                } else if (Character.isDigit(c)) {
                    start = true;
                    ret = c - '0';
                } else {
                    return 0;
                }
                continue;
            }

            if (c < '0' || c > '9') {
                break;
            }

            ret = (ret * 10) + (c - '0');
            if (isMin && ret * -1 <= Integer.MIN_VALUE) {
                return Integer.MIN_VALUE;
            }
            if (!isMin && ret >= Integer.MAX_VALUE) {
                return Integer.MAX_VALUE;
            }
        }

        return (int) ret * (isMin ? -1 : 1);
    }
}