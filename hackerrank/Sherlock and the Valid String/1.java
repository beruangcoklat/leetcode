import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the isValid function below.
    static String isValid(String s) {
        Map<Character, Integer> map = new HashMap<>();
        for (char c : s.toCharArray()) {
            Integer count = map.get(c);
            map.put(c, count == null ? 1 : count + 1);
        }

        Map<Integer, Integer> countMap = new HashMap<>();
        int map2Size = 0;

        for (Map.Entry<Character, Integer> e : map.entrySet()) {
            Integer count = countMap.get(e.getValue());
            if (count == null) {
                count = 0;
                map2Size++;

                if (map2Size > 2) {
                    return "NO";
                }
            }
            countMap.put(e.getValue(), count + 1);
        }

        if (map2Size == 1) {
            return "YES";
        }

        boolean first = true;
        Map.Entry<Integer, Integer> a = null, b = null;

        for (Map.Entry<Integer, Integer> e : countMap.entrySet()) {
            if (first) {
                first = false;
                a = e;
            } else {
                b = e;
            }
        }

        if (a.getValue() != 1 && b.getValue() != 1) {
            return "NO";
        }

        if (a.getValue() == 1) {
            if (a.getKey() != 1 && a.getKey() - 1 != b.getKey()) {
                return "NO";
            }
        }

        if (b.getValue() == 1) {
            if (b.getKey() != 1 && b.getKey() - 1 != a.getKey()) {
                return "NO";
            }
        }

        return "YES";
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = scanner.nextLine();

        String result = isValid(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
