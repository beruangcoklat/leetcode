class Solution {
    public boolean isBipartite(int[][] graph) {
        final int size = graph.length;
        int color[] = new int[size];
        boolean visited[] = new boolean[size];

        final int RED = 1;
        final int BLUE = 2;

        for (int src = 0; src < graph.length; src++) {
            if (color[src] != 0) continue;

            List<Integer> list = new ArrayList<>();
            list.add(src);
            color[src] = RED;

            while (!list.isEmpty()) {
                int curr = list.get(0);
                list.remove(0);

                if (visited[curr]) continue;
                visited[curr] = true;

                int currColor = color[curr];
                int nextColor = currColor == RED ? BLUE : RED;

                for (int i = 0; i < graph[curr].length; i++) {
                    int neighborNode = graph[curr][i];
                    if (color[neighborNode] == currColor) return false;
                    color[neighborNode] = nextColor;
                    list.add(neighborNode);
                }
            }
        }

        return true;
    }
}
