import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static boolean solve(String strA, String strB, HashMap<String, Boolean> memo) {
        String key = strA + "-" + strB;
        Boolean mem = memo.get(key);
        if (mem != null) return mem;

        boolean temp = false;
        
        int lenA = strA.length();
        int lenB = strB.length();
        
        if(lenA < lenB){
            memo.put(key, false);
            return false;
        }

        if (lenB == 0) {
            for (char c : strA.toCharArray()) {
                if (Character.isUpperCase(c)) {
                    memo.put(key, false);
                    return false;
                }
            }
            memo.put(key, true);
            return true;
        }

        char a = strA.charAt(0);
        char b = strB.charAt(0);

        if (a == b) {
            temp = solve(strA.substring(1), strB.substring(1), memo);
            memo.put(key, temp);
            return temp;
        }

        if (Character.toUpperCase(a) == b) {
            temp = solve(strA.substring(1), strB.substring(1), memo);
            if (temp) {
                memo.put(key, true);
                return true;
            }
            temp = solve(strA.substring(1), strB, memo);
            if (temp) {
                memo.put(key, true);
                return true;
            }
            memo.put(key, false);
            return false;
        }

        if (Character.isUpperCase(a)) {
            memo.put(key, false);
            return false;
        }

        temp = solve(strA.substring(1), strB, memo);
        memo.put(key, temp);
        return temp;
    }

    static String abbreviation(String a, String b) {
        return solve(a, b, new HashMap<String, Boolean>()) ? "YES" : "NO";
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int q = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int qItr = 0; qItr < q; qItr++) {
            String a = scanner.next();

            String b = scanner.next();

            String result = abbreviation(a, b);

            bufferedWriter.write(result);
            bufferedWriter.newLine();
        }

        bufferedWriter.close();

        scanner.close();
    }
}
