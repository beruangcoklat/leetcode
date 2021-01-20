import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

public class Solution {

    static String getString(String s, char a, char b) {
        String ret = "";
        char curr = ' ';
        boolean first = true;
        for (char c : s.toCharArray()) {
            if (c != a && c != b) continue;
            if (c == curr && !first) {
                return "";
            }
            first = false;
            curr = c;
            ret += c;
        }
        return ret;
    }

    static int alternate(String s) {
        HashMap<Character, Byte> map = new HashMap<>();
        List<Character> chars = new ArrayList<>();
        for (char a : s.toCharArray()) {
            if (map.containsKey(a)) {
                continue;
            }
            map.put(a, null);
            chars.add(a);
        }

        if (chars.size() == 1) {
            return 0;
        }

        int max = Integer.MIN_VALUE;
        for (int i = 0; i < chars.size(); i++) {
            for (int j = i + 1; j < chars.size(); j++) {
                String c = getString(s, chars.get(i), chars.get(j));
                max = Math.max(max, c.length());
            }
        }
        return max;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int l = Integer.parseInt(bufferedReader.readLine().trim());

        String s = bufferedReader.readLine();

        int result = alternate(s);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
