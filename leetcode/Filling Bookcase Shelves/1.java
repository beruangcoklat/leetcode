class Solution {
    Map<String, Integer> cache;

    public int minHeightShelves(int[][] books, int shelfWidth) {
        cache = new HashMap<>();
        return solve(books, shelfWidth, 0, shelfWidth, 0);
    }

    int solve(int[][] books, int shelfWidth, int bookIdx, int spaceLeft, int maxHeight) {
        String key = String.format("%d-%d-%d", bookIdx, spaceLeft, maxHeight);
        Integer cacheResult = cache.get(key);
        if (cacheResult != null) {
            return cacheResult;
        }

        if (bookIdx == books.length) {
            return maxHeight;
        }

        int thickness = books[bookIdx][0];
        int height = books[bookIdx][1];

        int moveFloor = solve(books, shelfWidth, bookIdx + 1, shelfWidth - thickness, height) + maxHeight;
        int sameFloor = moveFloor;

        if (bookIdx == 0) {
            cache.put(key, moveFloor);
            return moveFloor;
        }

        if (thickness <= spaceLeft) {
            sameFloor = solve(books, shelfWidth, bookIdx + 1, spaceLeft - thickness, Math.max(maxHeight, height));
        }

        int min = Math.min(sameFloor, moveFloor);
        cache.put(key, min);
        return min;
    }
}