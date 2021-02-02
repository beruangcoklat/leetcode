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

                while (parent[source] != 0) {
                    if(parent[parent[source]] != 0)
                        parent[source] = parent[parent[source]];
                    source = parent[source];
                }

                while (parent[destination] != 0) {
                    if(parent[parent[destination]] != 0)
                        parent[destination] = parent[parent[destination]];
                    destination = parent[destination];
                }

                if (source == destination) continue;

                parent[source] = destination;
                count[parent[source]] += count[source];
                count[source] = 0;
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
