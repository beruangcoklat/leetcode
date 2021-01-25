class Solution {
    public int maxRepeating(String sequence, String word) {
        int ret = 0;
        String str = word;
        while (true) {
            if (sequence.indexOf(str) == -1) {
                return ret;
            }
            ret++;
            str += word;
        }
    }
}
