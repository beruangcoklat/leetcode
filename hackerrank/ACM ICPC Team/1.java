import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the acmTeam function below.
    static int getSubjects(String a, String b) {
        int ret = 0;
        for (int i = 0; i < a.length(); i++) {
            if (a.charAt(i) == '1' || b.charAt(i) == '1') {
                ret++;
            }
        }
        return ret;
    }

    static int[] acmTeam(String[] topic) {
        int max = Integer.MIN_VALUE;
        HashMap<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < topic.length; i++) {
            for (int j = i + 1; j < topic.length; j++) {
                int subjects = getSubjects(topic[i], topic[j]);
                Integer count = map.get(subjects);
                if (count == null) count = 0;
                if (subjects > max) max = subjects;

                map.put(subjects, count + 1);
            }
        }

        return new int[]{max, map.get(max)};
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nm = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nm[0]);

        int m = Integer.parseInt(nm[1]);

        String[] topic = new String[n];

        for (int i = 0; i < n; i++) {
            String topicItem = scanner.nextLine();
            topic[i] = topicItem;
        }

        int[] result = acmTeam(topic);

        for (int i = 0; i < result.length; i++) {
            bufferedWriter.write(String.valueOf(result[i]));

            if (i != result.length - 1) {
                bufferedWriter.write("\n");
            }
        }

        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
