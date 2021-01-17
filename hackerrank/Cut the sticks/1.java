import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the cutTheSticks function below.
    static int[] cutTheSticks(int[] arr) {
        int min = Integer.MAX_VALUE;
        List<Integer> list = new ArrayList<>();
        List<Integer> ret = new ArrayList<>();

        for (int a : arr) {
            list.add(a);
            if (a < min) min = a;
        }

        while (!list.isEmpty()) {
            int cut = 0;
            int nextMin = Integer.MAX_VALUE;
            for (int i = 0; i < list.size(); i++) {
                int curr = list.get(i);
                curr -= min;
                cut++;
                if (curr <= 0) {
                    list.remove(i);
                    i--;
                } else {
                    list.set(i, curr);
                    if(curr<nextMin) {
                        nextMin = curr;
                    }
                }
            }
            min = nextMin;
            ret.add(cut);
        }

        int retArr[] = new int[ret.size()];
        for (int i = 0; i < ret.size(); i++) {
            retArr[i] = ret.get(i);
        }
        return retArr;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int[] arr = new int[n];

        String[] arrItems = scanner.nextLine().split(" ");
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int i = 0; i < n; i++) {
            int arrItem = Integer.parseInt(arrItems[i]);
            arr[i] = arrItem;
        }

        int[] result = cutTheSticks(arr);

        for (int i = 0; i < result.length; i++) {
            bufferedWriter.write(String.valueOf(result[i]));

            if (i != result.length - 1) {
                bufferedWriter.write("\n");
            }
        }

        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
