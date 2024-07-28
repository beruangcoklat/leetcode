class Solution {
    private Map<String, Boolean> cache;

    public boolean isInterleave(String s1, String s2, String s3) {
        if (s3.length() != (s1.length() + s2.length())) {
            return false;
        }

        cache = new HashMap<>();
        return solve(s1, s2, s3, 0, 0, 0);
    }

    private boolean solve(String s1, String s2, String s3, int idx1, int idx2, int idx3) {
        if (idx3 == s3.length()) {
            return true;
        }

        String key = String.format("%d-%d-%d", idx1, idx2, idx3);
        if (cache.containsKey(key)) {
            return cache.get(key);
        }

        if (idx1 < s1.length() &&
            s1.charAt(idx1) == s3.charAt(idx3) &&
            solve(s1, s2, s3, idx1 + 1, idx2, idx3 + 1)
        ) {
            cache.put(key, true);
            return true;
        }
        if (idx2 < s2.length() &&
            s2.charAt(idx2) == s3.charAt(idx3) &&
            solve(s1, s2, s3, idx1, idx2 + 1, idx3 + 1)
        ) {
            cache.put(key, true);
            return true;
        }

        cache.put(key, false);
        return false;
    }
}