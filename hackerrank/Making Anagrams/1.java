import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the makingAnagrams function below.
    static int makingAnagrams(String s1, String s2) {
        HashMap<Character, Integer> map1 = new HashMap<>();
        HashMap<Character, Integer> map2 = new HashMap<>();
        HashMap<Character, Boolean> mapAll = new HashMap<>();

        for(char c : s1.toCharArray()) {
            Integer count = map1.get(c);
            if (count == null) count = 0;
            map1.put(c, count + 1);
            mapAll.put(c, true);
        }

        for(char c : s2.toCharArray()) {
            Integer count = map2.get(c);
            if (count == null) count = 0;
            map2.put(c, count + 1);
            mapAll.put(c, true);
        }

        int ret = 0;
        for(Map.Entry<Character,Boolean> e : mapAll.entrySet()) {
            char key = e.getKey();
            Integer count1 = map1.get(key);
            Integer count2 = map2.get(key);

            if (count1 == null) count1 = 0;
            if (count2 == null) count2 = 0;

            ret += Math.abs(count1 - count2);
        }
        
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s1 = scanner.nextLine();

        String s2 = scanner.nextLine();

        int result = makingAnagrams(s1, s2);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
