import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static long arrayManipulation(int n, int[][] queries) {
        HashMap<Integer, Long> increase = new HashMap();
        HashMap<Integer, Long> decrease = new HashMap();

        for (int i = 0; i < queries.length; i++) {
            Long a = increase.get(queries[i][0]);
            Long b = decrease.get(queries[i][1] + 1);
            if (a == null) a = 0l;
            if (b == null) b = 0l;
            increase.put(queries[i][0], a + queries[i][2]);
            decrease.put(queries[i][1] + 1, b + queries[i][2]);
        }

        long max = 0l;
        long curr = 0l;
        Long incr = 0l;
        Long decr = 0l;

        for (int i = 1; i <= n; i++) {
            incr = increase.get(i);
            decr = decrease.get(i);

            if (incr != null) curr += incr;
            if (decr != null) curr -= decr;

            if (curr > max) max = curr;
        }

        return max;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nm = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nm[0]);

        int m = Integer.parseInt(nm[1]);

        int[][] queries = new int[m][3];

        for (int i = 0; i < m; i++) {
            String[] queriesRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < 3; j++) {
                int queriesItem = Integer.parseInt(queriesRowItems[j]);
                queries[i][j] = queriesItem;
            }
        }

        long result = arrayManipulation(n, queries);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
