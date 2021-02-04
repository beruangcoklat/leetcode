import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static class Edge {
        int src, dst, weight;

        public Edge(int src, int dst, int weight) {
            this.src = src;
            this.dst = dst;
            this.weight = weight;
        }
    }

    static int findParent(int[]parent, int curr) {
        while(parent[curr] != 0){
            if(parent[parent[curr]] != 0) {
                parent[curr] = parent[parent[curr]];
            }
            curr = parent[curr];
        }
        return curr;
    }

    static int prims(int n, int[][] edges, int start) {
        Edge[] edgeList = new Edge[edges.length];
        int[] parent = new int[n + 1];

        for (int i = 0; i < edges.length; i++) {
            int a = edges[i][0];
            int b = edges[i][1];
            int w = edges[i][2];
            edgeList[i] = new Edge(a, b, w);
        }

        Arrays.sort(edgeList, (Edge a, Edge b) -> {
            return a.weight - b.weight;
        });

        int ret = 0;

        for (int i = 0; i < edges.length; i++) {
            Edge curr = edgeList[i];
            int a = findParent(parent, curr.src);
            int b = findParent(parent, curr.dst);
            if (a == b) continue;
            parent[a] = b;
            ret += curr.weight;
        }

        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nm = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nm[0]);

        int m = Integer.parseInt(nm[1]);

        int[][] edges = new int[m][3];

        for (int i = 0; i < m; i++) {
            String[] edgesRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < 3; j++) {
                int edgesItem = Integer.parseInt(edgesRowItems[j]);
                edges[i][j] = edgesItem;
            }
        }

        int start = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int result = prims(n, edges, start);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
