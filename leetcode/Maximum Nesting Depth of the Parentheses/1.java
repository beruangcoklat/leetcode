class Solution {
    public int maxDepth(String s) {
        int curr = 0;
        int total = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                curr++;
            } else if (c == ')') {
                curr--;
            }
            if (curr > total) {
                total = curr;
            }
        }
        return total;
    }
}