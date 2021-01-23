import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static BigInteger solve(int x, HashMap<Integer, BigInteger> memo) {
        BigInteger mem = memo.get(x);
        if (mem != null) return mem;
        
        BigInteger a = solve(x - 1, memo);
        BigInteger b = solve(x - 2, memo);
        BigInteger res = a.multiply(a).add(b);
        memo.put(x, res);
        return res;
    }

    static BigInteger fibonacciModified(int t1, int t2, int n) {
        HashMap<Integer, BigInteger> memo = new HashMap<>();
        memo.put(1, new BigInteger(t1 + ""));
        memo.put(2, new BigInteger(t2 + ""));
        return solve(n, memo);
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] t1T2n = scanner.nextLine().split(" ");

        int t1 = Integer.parseInt(t1T2n[0]);

        int t2 = Integer.parseInt(t1T2n[1]);

        int n = Integer.parseInt(t1T2n[2]);

        BigInteger result = fibonacciModified(t1, t2, n);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
