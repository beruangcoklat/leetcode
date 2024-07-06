class Solution {
    public int passThePillow(int n, int time) {
        if (time < n) {
            return time + 1;
        }
        int pos = time % (n - 1);
        int direction = (time / (n - 1)) % 2;
        if (direction == 0) {
            return pos + 1;
        }
        return n - pos;
    }
}