import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static int[] maxSubarray(int[] arr) {
        int ret[] = new int[2];
        int maxSubsequence = 0;
        int max = Integer.MIN_VALUE;
        int maxSubarray = Integer.MIN_VALUE;
        int localMax = 0;

        for(int i=0 ; i<arr.length ; i++) {
            if (arr[i] > 0) maxSubsequence += arr[i];
            if (arr[i] > max) max = arr[i];

            localMax += arr[i];
            if (localMax > maxSubarray) maxSubarray = localMax;
            if (localMax < 0) localMax = 0;
        }

        if(maxSubsequence == 0) {
            ret[1] = max;
            ret[0] = max;
        }else {
            ret[1] = maxSubsequence;
            ret[0] = maxSubarray;
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int t = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int tItr = 0; tItr < t; tItr++) {
            int n = scanner.nextInt();
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            int[] arr = new int[n];

            String[] arrItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int i = 0; i < n; i++) {
                int arrItem = Integer.parseInt(arrItems[i]);
                arr[i] = arrItem;
            }

            int[] result = maxSubarray(arr);

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
