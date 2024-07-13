class Solution {
    public int minOperations(String[] logs) {
        int pos = 0;
        for (String log : logs) {
            if (log.equals("./")) {
                continue;
            }

            if (log.equals("../")) {
                pos--;
            } else {
                pos++;
            }
            
            pos = Math.max(pos, 0);
        }
        return pos;
    }
}