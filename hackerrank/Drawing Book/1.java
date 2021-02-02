import java.io.*;
import java.math.*;
import java.text.*;
import java.util.*;
import java.util.regex.*;

public class Solution {

    static int openFront(int n, int p){
        if(p == 1) return 0;
        if(p % 2 == 1) p--;
        return p / 2;
    }

    static int openBack(int n, int p) {
        if (n % 2 == 1 && (p == n || p == n - 1)) return 0;
        if (n % 2 == 0 && n == p) return 0;

        int start = 0;
        if (n % 2 == 0) {
            start++;
            n--;
        }
        return ((n - p) / 2) + start;
    }

    static int pageCount(int n, int p) {
        return Math.min(openFront(n, p), openBack(n, p));
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])*");

        int p = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])*");

        int result = pageCount(n, p);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
