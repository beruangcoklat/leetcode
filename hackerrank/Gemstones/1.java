import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static int gemstones(String[] arr) {
        List<HashMap<Character, Byte>> listOfMap = new ArrayList<>();
        HashMap<Character, Byte> mapOfChar = new HashMap<>();
        for (String s : arr) {
            HashMap<Character, Byte> map = new HashMap<>();
            for (char c : s.toCharArray()) {
                map.put(c, null);
                mapOfChar.put(c, null);
            }
            listOfMap.add(map);
        }

        int gemCount = 0;
        for (Map.Entry<Character, Byte> e : mapOfChar.entrySet()) {
            char key = e.getKey();
            boolean isGem = true;
            for (HashMap<Character, Byte> map : listOfMap) {
                if (!map.containsKey(key)) {
                    isGem = false;
                    break;
                }
            }
            if (isGem) gemCount++;
        }
        return gemCount;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        String[] arr = new String[n];

        for (int i = 0; i < n; i++) {
            String arrItem = scanner.nextLine();
            arr[i] = arrItem;
        }

        int result = gemstones(arr);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
