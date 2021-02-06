import java.io.*;
import java.util.*;

public class Solution {
    static HashMap<Integer, Integer> getCostMap(int[][] graph, int[]connectedNode, int src) {
        List<Integer> list = new ArrayList<>();
        list.add(src);

        HashMap<Integer, Boolean> visited = new HashMap<>();
        HashMap<Integer, Integer> costMap = new HashMap<>();

        for (int moves = 0; !list.isEmpty(); moves++) {
            int currListSize = list.size();
            for (int i = 0; i < currListSize; i++) {
                int curr = list.get(0);
                list.remove(0);

                if (visited.containsKey(curr)) continue;
                visited.put(curr, true);

                costMap.put(curr, moves * 6);

                for (int j = 0; j < connectedNode[curr]; j++) {
                    list.add(graph[curr][j]);
                }
            }
        }

        return costMap;
    }

    static int[] bfs(int n, int m, int[][] edges, int s) {
        int[][] graph = new int[n + 1][n + 1];
        int[] connectedNode = new int[n + 1];

        HashMap<String, Boolean> marker = new HashMap<>();

        int src = 0, dst = 0, min = 0, max = 0;
        for (int i = 0; i < m; i++) {
            src = edges[i][0];
            dst = edges[i][1];
            if (src == dst) continue;
            min = Math.min(src, dst);
            max = Math.max(src, dst);
            String key = min + "_" + max;
            if (marker.containsKey(key)) continue;
            marker.put(key, true);

            graph[src][connectedNode[src]] = dst;
            graph[dst][connectedNode[dst]] = src;
            connectedNode[src]++;
            connectedNode[dst]++;
        }

        HashMap<Integer, Integer> costMap = getCostMap(graph, connectedNode, s);
        int ret[] = new int[n - 1];
        int idx = 0;

        for (int i = 1; i <= n; i++) {
            if (i == s) {
                continue;
            }

            Integer cost = costMap.get(i);
            if (cost == null) cost = -1;
            ret[idx++] = cost;
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int q = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int qItr = 0; qItr < q; qItr++) {
            String[] nm = scanner.nextLine().split(" ");

            int n = Integer.parseInt(nm[0]);

            int m = Integer.parseInt(nm[1]);

            int[][] edges = new int[m][2];

            for (int i = 0; i < m; i++) {
                String[] edgesRowItems = scanner.nextLine().split(" ");
                scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

                for (int j = 0; j < 2; j++) {
                    int edgesItem = Integer.parseInt(edgesRowItems[j]);
                    edges[i][j] = edgesItem;
                }
            }

            int s = scanner.nextInt();
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            int[] result = bfs(n, m, edges, s);

            for (int i = 0; i < result.length; i++) {
                bufferedWriter.write(String.valueOf(result[i]));

                if (i != result.length - 1) {
                    bufferedWriter.write(" ");
                }
            }

            bufferedWriter.newLine();
        }

        bufferedWriter.close();

        scanner.close();
    }
}