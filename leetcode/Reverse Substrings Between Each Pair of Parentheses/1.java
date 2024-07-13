class Solution {
    public String reverseParentheses(String s) {
        List<Integer> stack = new ArrayList<>();
        List<String> result = new ArrayList<>();

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            if (c == '(') {
                stack.add(result.size());
            } else if (c == ')') {
                reverse(result, stack.get(stack.size() - 1));
                stack.remove(stack.size() - 1);
            } else {
                result.add(String.valueOf(c));
            }
        }

        String resp = "";
        for (String r : result) {
            resp += r;
        }
        return resp;
    }

    private void reverse(List<String> result, int start) {
        int end = result.size() - 1;
        while (start < end) {
            String s = result.get(start);
            String e = result.get(end);
            result.set(start, e);
            result.set(end, s);
            start++;
            end--;
        }
    }
}
