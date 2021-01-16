import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the sherlockAndAnagrams function below.
    static boolean isAnagram(String a, String b) {
        HashMap<Character, Integer> map = new HashMap<>();

        for (char c : a.toCharArray()) {
            Integer count = map.get(c);
            if (count == null) count = 0;
            map.put(c, count + 1);
        }

        for (char c : b.toCharArray()) {
            Integer count = map.get(c);
            if (count == null || count == 0) {
                return false;
            }
            map.put(c, count - 1);
        }

        return true;
    }

    static int sherlockAndAnagrams(String s) {
        int len = 0;
        int ret = 0;

        while (true) {
            len++;
            boolean stop = true;
            for (int i = 0; i <= s.length() - len; i++) {
                String a = s.substring(i, i + len);
                for (int j = i + 1; j <= s.length() - len; j++) {
                    stop = false;
                    String b = s.substring(j, j + len);
                    if (isAnagram(a, b)) {
                        ret++;
                    }
                }
            }
            if (stop) break;
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int q = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int qItr = 0; qItr < q; qItr++) {
            String s = scanner.nextLine();

            int result = sherlockAndAnagrams(s);

            bufferedWriter.write(String.valueOf(result));
            bufferedWriter.newLine();
        }

        bufferedWriter.close();

        scanner.close();
    }
}
