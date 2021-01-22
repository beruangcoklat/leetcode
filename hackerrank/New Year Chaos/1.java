import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static void swap(int[] q, int a, int b) {
        int temp = q[a];
        q[a] = q[b];
        q[b] = temp;
    }

    static void minimumBribes(int[] q) {
        HashMap<Integer, Integer> map = new HashMap<>();
        int arr[] = new int[q.length];
        for (int i = 0; i < q.length; i++) {
            map.put(i + 1, i);
            arr[i] = i + 1;
        }

        int allSwapCount = 0;
        for (int i = 0; i < arr.length; i++) {
            int idx = map.get(q[i]);
            int swapCount = 0;
            while (idx != i) {
                swapCount++;
                if (swapCount > 2) {
                    System.out.println("Too chaotic");
                    return;
                }
                swap(arr, idx, idx - 1);
                map.put(arr[idx], idx);
                map.put(arr[idx - 1], idx - 1);
                idx--;
            }

            allSwapCount += swapCount;
        }
        System.out.println(allSwapCount);
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        int t = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int tItr = 0; tItr < t; tItr++) {
            int n = scanner.nextInt();
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            int[] q = new int[n];

            String[] qItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int i = 0; i < n; i++) {
                int qItem = Integer.parseInt(qItems[i]);
                q[i] = qItem;
            }

            minimumBribes(q);
        }

        scanner.close();
    }
}
