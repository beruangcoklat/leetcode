import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static String add(String num1, String num2) {
        String ret = "";
        int idx1 = num1.length() - 1;
        int idx2 = num2.length() - 1;
        int carrier = 0;
        while (idx1 >= 0 && idx2 >= 0) {
            int a = num1.charAt(idx1) - '0';
            int b = num2.charAt(idx2) - '0';

            int total = a + b + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }

            ret = (total + "") + ret;

            idx1--;
            idx2--;
        }

        while (idx1 >= 0) {
            int a = num1.charAt(idx1) - '0';
            int total = a + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }
            ret = (total + "") + ret;
            idx1--;
        }

        while (idx2 >= 0) {
            int a = num2.charAt(idx2) - '0';
            int total = a + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }
            ret = (total + "") + ret;
            idx2--;
        }

        return (carrier == 1 ? "1" : "") + ret;
    }

    static boolean validNumber(String str) {
        if (str.length() > 1 && str.charAt(0) == '0') {
            return false;
        }
        return true;
    }

    static void separateNumbers(String s) {
        for (int i = 1; i < s.length(); i++) {
            String firstNumber = s.substring(0, i);

            if (!validNumber(firstNumber)) {
                continue;
            }

            String expectedNumber = firstNumber;
            int idx = firstNumber.length();
            boolean valid = true;

            while (true) {
                expectedNumber = add(expectedNumber, "1");

                int expectedNumberDigit = expectedNumber.length();
                if (idx + expectedNumberDigit > s.length()) {
                    valid = false;
                    break;
                }
                String actualNumber = s.substring(idx, idx + expectedNumberDigit);

                if (!expectedNumber.equals(actualNumber)) {
                    valid = false;
                    break;
                }

                if (idx + expectedNumberDigit == s.length()) {
                    break;
                }

                idx += expectedNumberDigit;
            }

            if (valid) {
                System.out.println("YES " + firstNumber);
                return;
            }
        }

        System.out.println("NO");
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        int q = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int qItr = 0; qItr < q; qItr++) {
            String s = scanner.nextLine();

            separateNumbers(s);
        }

        scanner.close();
    }
}
