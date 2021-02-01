class Solution {
    public boolean isPowerOfTwo(int n) {
        HashMap<Integer, Boolean> map = new HashMap();
        map.put(1, true);
        int a = 1;
        for (int i = 1; i <= 30; i++) {
            a *= 2;
            map.put(a, true);
        }

        return map.containsKey(n);
    }
}
