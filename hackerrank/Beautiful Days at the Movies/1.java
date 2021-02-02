import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static int reverse(int x){
        int mult = 10;
        int ret = 0;
        while((mult / 10) <= x){
            int digit = (x % mult) / (mult / 10);
            ret *= 10;
            ret += digit;
            mult *= 10;
        }
        return ret;
    }

    // Complete the beautifulDays function below.
    static int beautifulDays(int a, int b, int k) {
        int ret = 0;
        for(int i=a ; i<=b ; i++){
            if((i - reverse(i)) % k == 0){
                ret++;
            }
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] ijk = scanner.nextLine().split(" ");

        int i = Integer.parseInt(ijk[0]);

        int j = Integer.parseInt(ijk[1]);

        int k = Integer.parseInt(ijk[2]);

        int result = beautifulDays(i, j, k);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
