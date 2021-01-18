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

class Result {

    /*
     * Complete the 'nonDivisibleSubset' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER k
     *  2. INTEGER_ARRAY s
     */

    public static int nonDivisibleSubset(int k, List<Integer> s) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int a : s) {
            int key = a % k;
            Integer count = map.get(key);
            if (count == null) count = 0;
            map.put(key, count + 1);
        }

        int ret = 0;
        if (map.containsKey(0)) {
            ret++;
        }

        for (int i = 1; i < k; i++) {
            if (i == k - i) {
                ret++;
                map.remove(i);
                continue;
            }

            Integer a = map.get(i);
            Integer b = map.get(k - i);

            if (a == null && b == null) continue;
            if (a == null && b != null) {
                ret += b;
                map.remove(k - i);
            } else if (a != null && b == null) {
                ret += a;
                map.remove(i);
            } else if (a != null && b != null) {
                ret += Math.max(a, b);
                map.remove(i);
                map.remove(k - i);
            }
        }

        return ret;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] firstMultipleInput = bufferedReader.readLine().replaceAll("\\s+$", "").split(" ");

        int n = Integer.parseInt(firstMultipleInput[0]);

        int k = Integer.parseInt(firstMultipleInput[1]);

        List<Integer> s = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
            .map(Integer::parseInt)
            .collect(toList());

        int result = Result.nonDivisibleSubset(k, s);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
