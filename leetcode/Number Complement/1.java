class Solution {
    public int findComplement(int num) {
        int result = 0;
        int idx = 0;
        while (num > 0) {
            int digit = num % 2;
            if (digit == 0) {
                result += Math.pow(2, idx);
            }
            num = num >> 1;
            idx++;
        }
        return result;
    }
}