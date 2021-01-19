import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static String highestValuePalindrome(String s, int n, int k) {
        HashMap<Integer, Byte> map = new HashMap<>();
        char[] res = s.toCharArray();

        for (int i = 0; i <= n - 1 - i; i++) {
            char a = s.charAt(i);
            char b = s.charAt(n - 1 - i);
            if (a == b) continue;
            if (a > b) {
                res[n - 1 - i] = a;
                k--;
                map.put(n - 1 - i, null);
            } else {
                res[i] = b;
                k--;
                map.put(i, null);
            }
            if (k < 0) {
                return "-1";
            }
        }

        int i = 0;
        while (i <= n - 1 - i) {
            if (i == n - 1 - i) {
                if (k >= 1) {
                    res[i] = '9';
                }
                break;
            }

            if (res[i] == '9') {
                i++;
                continue;
            }

            if (!map.containsKey(i) && !map.containsKey(n - 1 - i) && k >= 2) {
                res[i] = '9';
                res[n - 1 - i] = '9';
                k -= 2;
            } else if ((map.containsKey(i) || map.containsKey(n - 1 - i)) && k >= 1) {
                res[i] = '9';
                res[n - 1 - i] = '9';
                k--;
            }

            i++;
        }

        return new String(res);
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nk = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nk[0]);

        int k = Integer.parseInt(nk[1]);

        String s = scanner.nextLine();

        String result = highestValuePalindrome(s, n, k);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
