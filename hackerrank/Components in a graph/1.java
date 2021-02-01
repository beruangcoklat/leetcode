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

    static int getConnectedComponent(HashMap<Integer, List<Integer>> graph, int node, HashMap<Integer, Boolean> visited) {
        int total = 0;
        List<Integer> list = new ArrayList<>();
        list.add(node);
        while (!list.isEmpty()) {
            int curr = list.get(0);
            list.remove(0);
            if (visited.containsKey(curr)) continue;
            visited.put(curr, true);
            total++;

            List<Integer> next = graph.get(curr);
            if (next == null) continue;
            list.addAll(next);
        }
        return total;
    }

    public static List<Integer> componentsInGraph(List<List<Integer>> gb) {
        HashMap<Integer, List<Integer>> graph = new HashMap<>();
        Set<Integer> nodeSet = new HashSet<>();

        for (List<Integer> edge : gb) {
            int source = edge.get(0);
            int destination = edge.get(1);

            List<Integer> sourceN = graph.get(source);
            boolean ok = sourceN != null;
            if (!ok) {
                sourceN = new ArrayList<>();
            }
            sourceN.add(destination);
            if (!ok) {
                graph.put(source, sourceN);
            }

            List<Integer> destinationN = graph.get(destination);
            boolean ok2 = destinationN != null;
            if (!ok2) {
                destinationN = new ArrayList<>();
            }
            destinationN.add(source);
            if (!ok2) {
                graph.put(destination, destinationN);
            }

            nodeSet.add(source);
            nodeSet.add(destination);
        }

        int max = Integer.MIN_VALUE;
        int min = Integer.MAX_VALUE;
        HashMap<Integer, Boolean> visited = new HashMap<>();
        Iterator<Integer> nodes = nodeSet.iterator();
        while (nodes.hasNext()) {
            int node = nodes.next();
            if(visited.containsKey(node)){
                continue;
            }
            if (graph.get(node).size() == 0) {
                continue;
            }
            int result = getConnectedComponent(graph, node, visited);
            max = Math.max(max, result);
            min = Math.min(min, result);
        }

        List<Integer> ret = new ArrayList<>();
        ret.add(min);
        ret.add(max);
        return ret;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        List<List<Integer>> gb = new ArrayList<>();

        IntStream.range(0, n).forEach(i -> {
            try {
                gb.add(
                    Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                        .map(Integer::parseInt)
                        .collect(toList())
                );
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        List<Integer> result = Result.componentsInGraph(gb);

        bufferedWriter.write(
            result.stream()
                .map(Object::toString)
                .collect(joining(" "))
            + "\n"
        );

        bufferedReader.close();
        bufferedWriter.close();
    }
}
