import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static String getKey(int a, int b) {
        int max = Math.max(a, b);
        int min = Math.min(a, b);
        return min + "_" + max;
    }

    static int pairs(int k, int[] arr) {
        HashMap<Integer, Integer> arrMap = new HashMap<>();
        HashMap<String, Boolean> markMap = new HashMap<>();

        for (int a : arr) {
            Integer count = arrMap.get(a);
            if (count == null) count = 0;
            arrMap.put(a, count + 1);
        }

        int ret = 0;
        for (int a : arr) {
            String key = getKey(a, a - k);
            if (arrMap.containsKey(a - k) && !markMap.containsKey(key)) {
                ret++;
                markMap.put(key, true);
            }

            key = getKey(a, a + k);
            if (arrMap.containsKey(a + k) && !markMap.containsKey(key)) {
                ret++;
                markMap.put(key, true);
            }
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nk = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nk[0]);

        int k = Integer.parseInt(nk[1]);

        int[] arr = new int[n];

        String[] arrItems = scanner.nextLine().split(" ");
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int i = 0; i < n; i++) {
            int arrItem = Integer.parseInt(arrItems[i]);
            arr[i] = arrItem;
        }

        int result = pairs(k, arr);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
