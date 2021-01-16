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

    // Complete the countTriplets function below.
    static long countTriplets(List<Long> arr, long r) {
        long ret = 0;
        HashMap<Long, Long> rightMap = new HashMap<>();
        HashMap<Long, Long> leftMap = new HashMap<>();

        for (long a : arr) {
            Long count = rightMap.get(a);
            if (count == null) count = 0l;
            rightMap.put(a, count + 1);
        }

        for(long a : arr) {
            rightMap.put(a, rightMap.get(a) - 1l);

            long next = a * r;
            long prev = a / r;

            if(a % r == 0){
                Long nextCount = rightMap.get(next);
                Long prevCount = leftMap.get(prev);

                if(nextCount != null && prevCount != null){
                    ret += (nextCount * prevCount);
                }
            }

            Long count = leftMap.get(a);
            if (count == null) count = 0l;
            leftMap.put(a, count + 1);
        }

        return ret;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nr = bufferedReader.readLine().replaceAll("\\s+$", "").split(" ");

        int n = Integer.parseInt(nr[0]);

        long r = Long.parseLong(nr[1]);

        List<Long> arr = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
            .map(Long::parseLong)
            .collect(toList());

        long ans = countTriplets(arr, r);

        bufferedWriter.write(String.valueOf(ans));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
