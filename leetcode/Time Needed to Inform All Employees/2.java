class Solution {
    int dfs(HashMap<Integer, ArrayList<Integer>> graph, int node, int[] informTime) {
        if (!graph.containsKey(node)) {
            return 0;
        }

        ArrayList<Integer> arr = graph.get(node);
        int max = 0;
        for (int i = 0; i < arr.size(); i++) {
            int curr = dfs(graph, arr.get(i), informTime) + informTime[arr.get(i)];
            if (curr > max) {
                max = curr;
            }
        }
        return max;
    }

    public int numOfMinutes(int n, int headID, int[] manager, int[] informTime) {
        HashMap<Integer, ArrayList<Integer>> graph = new HashMap<>();
        for (int i = 0; i < manager.length; i++) {
            if (graph.containsKey(manager[i])) {
                graph.get(manager[i]).add(i);
                continue;
            }
            ArrayList<Integer> arr = new ArrayList<>();
            arr.add(i);
            graph.put(manager[i], arr);
        }

        return dfs(graph, -1, informTime);
    }
}