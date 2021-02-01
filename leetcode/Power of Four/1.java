class Solution {
    public boolean isPowerOfFour(int n) {
        HashMap<Integer, Boolean> map = new HashMap();
        map.put(1, true);
        int a = 1;
        while (true) {
            a *= 4;
            if (a <= 0) break;
            map.put(a, true);
        }
        return map.containsKey(n);
    }
}
