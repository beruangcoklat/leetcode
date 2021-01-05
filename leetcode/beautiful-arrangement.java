class Solution {
    
    int res = 0;

    void perm(int n, int visited, ArrayList<Integer> arr) {
        int size = arr.size();
        if (size > 0) {
            int x = arr.get(size - 1);
            if (x % size != 0 && size % x != 0) {
                return;
            }
        }

        if (size == n) {
            res++;
            return;
        }

        for (int i = 0; i < n; i++) {
            if ((visited & (1 << i)) != 0) continue;

            int newVisited = visited | (1 << i);
            ArrayList<Integer> cloneArr = (ArrayList<Integer>) arr.clone();
            cloneArr.add(i + 1);

            perm(n, newVisited, cloneArr);
        }
    }

    public int countArrangement(int n) {
        res = 0;
        perm(n, 0, new ArrayList());
        return res;
    }
}
