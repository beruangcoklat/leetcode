import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static void swap(int[] arr, int a, int b) {
        int temp = arr[a];
        arr[a] = arr[b];
        arr[b] = temp;
    }

    static int solve(int[] arr) {
        int sorted[] = cloneArr(arr, false);
        System.arraycopy(arr, 0, sorted, 0, arr.length);

        Arrays.sort(sorted);

        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            map.put(arr[i], i);
        }

        int swapCount = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == sorted[i]) continue;

            int trueIndex = map.get(sorted[i]);
            swap(arr, i, trueIndex);
            map.remove(arr[i]);

            map.put(arr[trueIndex], trueIndex);
            swapCount++;
        }

        return swapCount;
    }

    static int lilysHomework(int[] arr) {
        int[] cloned = cloneArr(arr, true);
        int a = solve(arr);
        int b = solve(cloned);
        return Math.min(a, b);
    }

    static int[] cloneArr(int[] arr, boolean isReversed) {
        int ret[] = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            int target = i;
            if (isReversed) target = arr.length - 1 - i;
            ret[i] = arr[target];
        }
        return ret;
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

        int result = lilysHomework(arr);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
