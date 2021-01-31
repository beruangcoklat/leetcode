class Solution {
    
    int[] getLPS(String pat) {
        int lps[] = new int[pat.length()];

        int i = 1;
        int len = 0;

        while (i < pat.length()) {
            if (pat.charAt(i) == pat.charAt(len)) {
                len++;
                lps[i] = len;
                i++;
                continue;
            }

            if (len == 0) {
                lps[i] = 0;
                i++;
                continue;
            }

            len = lps[len - 1];
        }

        return lps;
    }

    int kmp(String txt, String pat) {
        int[] lps = getLPS(pat);
        int i = 0, j = 0;

        while (i < txt.length()) {
            if (txt.charAt(i) == pat.charAt(j)) {
                i++;
                j++;
                if (j == pat.length()) {
                    return i - pat.length();
                }
                continue;
            }
            
            if (j == 0) {
                i++;
                continue;
            }

            j = lps[j - 1];
        }
        return -1;
    }
    
    public int strStr(String haystack, String needle) {
        if (needle.length() == haystack.length()) {
            return needle.equals(haystack) ? 0 : -1;
        }
        if (needle.length() > haystack.length()) {
            return -1;
        }
        if (needle.length() == 0) {
            return 0;
        }
        return kmp(haystack, needle);
    }
}