class Solution {
    int solve(int leafID, int n, int headID, int[] manager, int[] informTime) {
        int ret = 0;
        int currID = leafID;
        while (manager[currID] != -1) {
            currID = manager[currID];
            ret += informTime[currID];
        }
        return ret;
    }

    public int numOfMinutes(int n, int headID, int[] manager, int[] informTime) {
        int max = 0;
        for (int i = 0; i < informTime.length; i++) {
            if (informTime[i] != 0) continue;

            int curr = solve(i, n, headID, manager, informTime);
            if (curr > max) {
                max = curr;
            }
        }
        return max;
    }
}
