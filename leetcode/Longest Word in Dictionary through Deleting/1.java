class Solution {
    
    boolean possible(String str, String dictionary) {
        int idxStr = 0;
        int idxDict = 0;
        while (idxStr < str.length() && idxDict < dictionary.length()) {
            char a = str.charAt(idxStr);
            char b = dictionary.charAt(idxDict);

            if (a == b) {
                idxStr++;
                idxDict++;
            } else {
                idxStr++;
            }
        }

        return idxDict == dictionary.length();
    }
    
    public String findLongestWord(String s, List<String> d) {
        int longestCharacterCount = 0;
        String ret = "";

        for (String str : d) {
            if (!possible(s, str)) continue;

            if (str.length() > longestCharacterCount) {
                longestCharacterCount = str.length();
                ret = str;
            } else if (str.length() == longestCharacterCount && str.compareTo(ret) < 0) {
                ret = str;
            }
        }
        return ret;
    }
}
