import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {
    static class FastReader {
        BufferedReader br;
        StringTokenizer st;

        public FastReader() {
            br = new BufferedReader(new InputStreamReader(System.in));
        }

        String next() {
            while (st == null || !st.hasMoreElements()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }

        long nextLong() {
            return Long.parseLong(next());
        }

        double nextDouble() {
            return Double.parseDouble(next());
        }

        String nextLine() {
            String str = "";
            try {
                str = br.readLine();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return str;
        }
    }

    static class MinHeap {
        Node[] arr;
        int idx;

        MinHeap(int n) {
            arr = new Node[n];
            idx = 1;
        }

        boolean isEmpty() {
            return idx <= 1;
        }

        void add(Node val) {
            arr[idx++] = val;
            int curr = idx - 1;
            int parent = curr / 2;

            while (parent > 0 && arr[parent].cost > arr[curr].cost) {
                Node temp = arr[parent];
                arr[parent] = arr[curr];
                arr[curr] = temp;
                curr = parent;
                parent /= 2;
            }
        }

        Node get() {
            Node ret = arr[1];
            arr[1] = arr[idx - 1];
            idx--;
            int curr = 1;
            while (true) {
                int leftIdx = curr * 2;
                int rightIdx = curr * 2 + 1;

                if (leftIdx >= idx) break;

                int chosenIdx = -1;
                if (leftIdx < idx) chosenIdx = leftIdx;
                if (rightIdx < idx && arr[rightIdx].cost < arr[leftIdx].cost) chosenIdx = rightIdx;
                if (chosenIdx != -1 && arr[chosenIdx].cost >= arr[curr].cost) chosenIdx = -1;

                if (chosenIdx == -1) break;

                Node temp = arr[curr];
                arr[curr] = arr[chosenIdx];
                arr[chosenIdx] = temp;

                curr = chosenIdx;
            }
            return ret;
        }
    }

    static class Node {
        int dst, cost;

        Node(int dst, int cost) {
            this.dst = dst;
            this.cost = cost;
        }
    }

    static int[] shortestReach(int n, int m, Node[][] graph, int[] counter, int[][] costTable, int s) {
        MinHeap heap = new MinHeap(m + 1);
        heap.add(new Node(s, 0));

        HashMap<Integer, Boolean> visited = new HashMap<>();
        HashMap<Integer, Integer> result = new HashMap<>();

        while (!heap.isEmpty()) {
            Node curr = heap.get();
            if (!result.containsKey(curr.dst)) {
                result.put(curr.dst, curr.cost);
            }

            if (visited.containsKey(curr.dst)) continue;
            visited.put(curr.dst, true);

            for (int i = 0; i < counter[curr.dst]; i++) {
                if (visited.containsKey(graph[curr.dst][i].dst)) {
                    continue;
                }
                graph[curr.dst][i].cost += (curr.cost + costTable[curr.dst][graph[curr.dst][i].dst]);
                heap.add(graph[curr.dst][i]);
            }
        }

        int ret[] = new int[n - 1];
        int idx = 0;
        for (int i = 1; i <= n; i++) {
            if (i == s) continue;
            Integer cost = result.get(i);
            if (cost == null) cost = -1;
            ret[idx++] = cost;
        }

        return ret;
    }

    public static void main(String[] args) {
        FastReader in = new FastReader();
        int t = in.nextInt();
        for (int tItr = 0; tItr < t; tItr++) {
            int n = in.nextInt();
            int m = in.nextInt();

            Node[][] graph = new Node[n + 1][n + 1];
            int[][] costTable = new int[n + 1][n + 1];
            int[] counter = new int[n + 1];

            for (int i = 0; i < m; i++) {
                int src = in.nextInt();
                int dst = in.nextInt();
                int cost = in.nextInt();

                if (costTable[src][dst] != 0) {
                    if (costTable[src][dst] <= cost) continue;
                    costTable[src][dst] = cost;
                    costTable[dst][src] = cost;
                    continue;
                }

                costTable[src][dst] = cost;
                costTable[dst][src] = cost;

                graph[src][counter[src]] = new Node(dst, 0);
                counter[src]++;
                graph[dst][counter[dst]] = new Node(src, 0);
                counter[dst]++;
            }

            int s = in.nextInt();
            int[] result = shortestReach(n, m, graph, counter, costTable, s);
            for (int i = 0; i < result.length; i++) {
                System.out.print(result[i]);
                if (i != result.length - 1) {
                    System.out.print(" ");
                }
            }
            System.out.println();
        }
    }
}
