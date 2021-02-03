import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

class Result {

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

    public static int kruskals(int gNodes, List<Integer> gFrom, List<Integer> gTo, List<Integer> gWeight) {
        int[] parent = new int[gNodes + 1];
        Edge[] edges = new Edge[gFrom.size()];
        for (int i = 0; i < gFrom.size(); i++) {
            edges[i] = new Edge(gFrom.get(i), gTo.get(i), gWeight.get(i));
        }

        Arrays.sort(edges, (Edge a, Edge b) -> {
            return a.weight - b.weight;
        });

        int ret = 0;

        for (int i = 0; i < edges.length; i++) {
            Edge curr = edges[i];
            int a = findParent(parent, curr.src);
            int b = findParent(parent, curr.dst);
            if (a == b) continue;
            parent[a] = b;
            ret += curr.weight;
        }

        return ret;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] gNodesEdges = bufferedReader.readLine().replaceAll("\\s+$", "").split(" ");

        int gNodes = Integer.parseInt(gNodesEdges[0]);
        int gEdges = Integer.parseInt(gNodesEdges[1]);

        List<Integer> gFrom = new ArrayList<>();
        List<Integer> gTo = new ArrayList<>();
        List<Integer> gWeight = new ArrayList<>();

        IntStream.range(0, gEdges).forEach(i -> {
            try {
                String[] gFromToWeight = bufferedReader.readLine().replaceAll("\\s+$", "").split(" ");

                gFrom.add(Integer.parseInt(gFromToWeight[0]));
                gTo.add(Integer.parseInt(gFromToWeight[1]));
                gWeight.add(Integer.parseInt(gFromToWeight[2]));
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        int res = Result.kruskals(gNodes, gFrom, gTo, gWeight);

        // Write your code here.
        bufferedWriter.write(String.valueOf(res));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
