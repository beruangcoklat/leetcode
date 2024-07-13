class Solution {
    public int maximumGain(String s, int x, int y) {
        String str1, str2;
        int score1, score2;
        if (x > y) {
            str1 = "ab";
            score1 = x;
            str2 = "ba";
            score2 = y;
        } else {
            str1 = "ba";
            score1 = y;
            str2 = "ab";
            score2 = x;
        }

        List<Character> stack = new ArrayList<>();
        int result = 0;
        
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            char top = ' ';
            if (stack.size() > 0) {
                top = stack.get(stack.size() - 1);
            }

            if (top == str1.charAt(0) && c == str1.charAt(1)) {
                result += score1;
                stack.remove(stack.size() - 1);
                continue;
            }

            stack.add(c);
        }

        List<Character> nextString = new ArrayList<>(stack);
        stack.clear();

        for (int i = 0; i < nextString.size(); i++) {
            char c = nextString.get(i);

            char top = ' ';
            if (stack.size() > 0) {
                top = stack.get(stack.size() - 1);
            }

            if (top == str2.charAt(0) && c == str2.charAt(1)) {
                result += score2;
                stack.remove(stack.size() - 1);
                continue;
            }

            stack.add(c);
        }

        return result;
    }
}