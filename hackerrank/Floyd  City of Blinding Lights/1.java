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
        int id, cost;

        Node(int id) {
            this.id = id;
            this.cost = 0;
        }
    }

    static int[] shortestPath(int src, int n, int m, Node[][] graph, int[][] costTable, int[] counter) {
        MinHeap heap = new MinHeap(m + 1);
        boolean visited[] = new boolean[n + 1];
        int ret[] = new int[n + 1];
        Arrays.fill(ret, -1);

        heap.add(new Node(src));

        while (!heap.isEmpty()) {
            Node curr = heap.get();
            if (ret[curr.id] == -1) {
                ret[curr.id] = curr.cost;
            }
            if (visited[curr.id]) continue;
            visited[curr.id] = true;

            for (int i = 0; i < counter[curr.id]; i++) {
                Node next = graph[curr.id][i];
                if (visited[next.id]) continue;
                next.cost = (curr.cost + costTable[curr.id][next.id]);
                heap.add(next);
            }
        }

        return ret;
    }

    public static void main(String[] args) {
        FastReader in = new FastReader();

        int n = in.nextInt();
        int m = in.nextInt();

        Node[][] graph = new Node[n + 1][n + 1];
        int[][] costTable = new int[n + 1][n + 1];
        boolean[][] filled = new boolean[n + 1][n + 1];
        int[] counter = new int[n + 1];
        for (int i = 0; i < m; i++) {
            int src = in.nextInt();
            int dst = in.nextInt();
            int cost = in.nextInt();
            
            if (filled[src][dst]) {
                costTable[src][dst] = cost;
                continue;
            }

            filled[src][dst] = true;
            costTable[src][dst] = cost;
            graph[src][counter[src]] = new Node(dst);
            counter[src]++;
        }

        int resultMap[][] = new int[n + 1][n + 1];
        for (int i = 0; i < n + 1; i++) {
            resultMap[i] = shortestPath(i, n, m, graph, costTable, counter);
        }
        int q = in.nextInt();
        for (int qItr = 0; qItr < q; qItr++) {
            int src = in.nextInt();
            int dst = in.nextInt();
            System.out.println(resultMap[src][dst]);
        }
    }
}
