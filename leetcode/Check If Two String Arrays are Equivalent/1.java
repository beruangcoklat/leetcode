class Solution {
    public boolean arrayStringsAreEqual(String[] word1, String[] word2) {
        String ret = "";
        for (String w : word1) {
            ret += w;
        }

        Arrays.sort(word2, (String a, String b) -> {
            return b.length() - a.length();
        });

        for (String w : word2) {
            ret = ret.replace(w, "");
        }

        return ret.length() == 0;
    }
}
