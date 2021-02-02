import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static void merge(int[] serveTime, int[]customers, int left, int mid, int right) {
        int idxA = 0;
        int idxB = 0;
        int idx = left;

        int[] stA = new int[mid - left + 1];
        int[] cA = new int[mid - left + 1];
        int[] stB = new int[right - mid];
        int[] cB = new int[right - mid];

        for (int i = left; i < mid + 1; i++) {
            stA[i - left] = serveTime[i];
            cA[i - left] = customers[i];
        }
        for (int i = mid + 1; i < right + 1; i++) {
            stB[i - mid - 1] = serveTime[i];
            cB[i - mid - 1] = customers[i];
        }

        while (idxA < stA.length && idxB < stB.length) {
            int min = 0;
            int cust = 0;
            if (stA[idxA] < stB[idxB]) {
                min = stA[idxA];
                cust = cA[idxA];
                idxA++;
            } else if (stA[idxA] > stB[idxB]) {
                min = stB[idxB];
                cust = cB[idxB];
                idxB++;
            } else if (cA[idxA] < cB[idxB]) {
                min = stA[idxA];
                cust = cA[idxA];
                idxA++;
            } else {
                min = stB[idxB];
                cust = cB[idxB];
                idxB++;
            }

            serveTime[idx] = min;
            customers[idx] = cust;
            idx++;
        }

        while (idxA < stA.length) {
            serveTime[idx] = stA[idxA];
            customers[idx] = cA[idxA];
            idx++;
            idxA++;
        }

        while (idxB < stB.length) {
            serveTime[idx] = stB[idxB];
            customers[idx] = cB[idxB];
            idx++;
            idxB++;
        }
    }

    static void mergeSort(int[] serveTime, int[]customers, int left, int right) {
        if (left == right) return;

        int mid = (left + right) / 2;
        mergeSort(serveTime, customers, left, mid);
        mergeSort(serveTime, customers, mid + 1, right);
        merge(serveTime, customers, left, mid, right);
    }

    static int[] jimOrders(int[][] orders) {
        int serveTime[] = new int[orders.length];
        int customers[] = new int[orders.length];

        for (int i = 0; i < orders.length; i++) {
            int st = orders[i][0] + orders[i][1];
            int c = i + 1;
            customers[i] = c;
            serveTime[i] = st;
        }

        mergeSort(serveTime, customers, 0, orders.length - 1);

        return customers;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int[][] orders = new int[n][2];

        for (int i = 0; i < n; i++) {
            String[] ordersRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < 2; j++) {
                int ordersItem = Integer.parseInt(ordersRowItems[j]);
                orders[i][j] = ordersItem;
            }
        }

        int[] result = jimOrders(orders);

        for (int i = 0; i < result.length; i++) {
            bufferedWriter.write(String.valueOf(result[i]));

            if (i != result.length - 1) {
                bufferedWriter.write(" ");
            }
        }

        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
