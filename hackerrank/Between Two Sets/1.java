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
     * Complete the 'getTotalX' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER_ARRAY a
     *  2. INTEGER_ARRAY b
     */

    static int gcd(int a, int b) {
        if (b == 0) return a;
        return gcd(b, a % b);
    }

    static int lcm(int a, int b) {
//        a*b = lcm(a,b) * gcd(a,b)
//        lcm(a,b) = (a*b) / gcd(a,b)
        return (a * b) / gcd(a, b);
    }

    static int findGcd(List<Integer> list){
        if(list.size() == 1)
            return list.get(0);

        int fpb = gcd(list.get(0),list.get(1));
        for (int i = 2; i < list.size(); i++) {
            fpb = gcd(fpb, list.get(i));
        }

        return fpb;
    }

    static int findLcm(List<Integer> list) {
        int kpk = 1;
        for (int c : list) {
            kpk = lcm(kpk, c);
        }
        return kpk;
    }

    public static int getTotalX(List<Integer> a, List<Integer> b) {
        int kpk = findLcm(a);
        int fpb = findGcd(b);
        int ret = 0;
        for (int i = kpk; i <= fpb; i += kpk) {
            if (fpb % i == 0) ret++;
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

        int m = Integer.parseInt(firstMultipleInput[1]);

        List<Integer> arr = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
            .map(Integer::parseInt)
            .collect(toList());

        List<Integer> brr = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
            .map(Integer::parseInt)
            .collect(toList());

        int total = Result.getTotalX(arr, brr);

        bufferedWriter.write(String.valueOf(total));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
