import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static char[] cloneChars(char[] chars) {
        char ret[] = new char[chars.length];
        for (int i = 0; i < chars.length; i++) {
            ret[i] = chars[i];
        }
        return ret;
    }

    static int solve(char[] str, int idx) {
        if (idx >= str.length) return 0;

        if (idx + 2 < str.length && str[idx] == '0' && str[idx + 1] == '1' && str[idx + 2] == '0') {
            char[] cloneA = cloneChars(str);
            char[] cloneB = cloneChars(str);
            char[] cloneC = cloneChars(str);
            cloneA[idx] = '1';
            cloneB[idx + 1] = '0';
            cloneC[idx + 2] = '1';

            int a = solve(cloneA, idx + 1) + 1;
            int b = solve(cloneB, idx + 2) + 1;
            int c = solve(cloneB, idx + 3) + 1;
            return Math.min(a, Math.min(b, c));
        }
        return solve(str, idx + 1);
    }

    static int beautifulBinaryString(String b) {
        return solve(b.toCharArray(), 0);
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        String b = scanner.nextLine();

        int result = beautifulBinaryString(b);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
