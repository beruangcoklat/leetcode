class Solution {
    List<String> result;
    
    void recursive(String s, int idx) {
        if (idx == s.length()) {
            result.add(s);
            return;
        }

        if (!Character.isAlphabetic(s.charAt(idx))) {
            recursive(s, idx + 1);
            return;
        }

        char word1[] = s.toCharArray();
        char word2[] = s.toCharArray();

        word1[idx] = Character.toUpperCase(word1[idx]);
        word2[idx] = Character.toLowerCase(word2[idx]);

        recursive(new String(word1), idx + 1);
        recursive(new String(word2), idx + 1);
    }
    
    public List<String> letterCasePermutation(String S) {
        result = new ArrayList<>();
        recursive(S, 0);
        return result;
    }
}
