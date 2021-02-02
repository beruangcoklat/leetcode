import java.io.*;
import java.util.*;

public class Solution {
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int q = scanner.nextInt();

        int[] parent = new int[n + 1];
        int[] count = new int[n + 1];
        Arrays.fill(count, 1);

        for (int i = 0; i < q; i++) {
            char type = scanner.next().charAt(0);
            if (type == 'M') {
                int source = scanner.nextInt();
                int destination = scanner.nextInt();

                if (source == destination) continue;

                int max = Math.max(source, destination);
                int min = Math.min(source, destination);

                while (parent[min] != 0) {
                    if(parent[parent[min]] != 0)
                        parent[min] = parent[parent[min]];
                    min = parent[min];
                }

                while (parent[max] != 0) {
                    max = parent[max];
                }

                if (min == max) continue;

                parent[min] = max;
                count[parent[min]] += count[min];
                count[min] = 0;
            } else {
                int node = scanner.nextInt();
                while (parent[node] != 0) {
                    node = parent[node];
                }
                System.out.println(count[node]);
            }
        }
    }
}
