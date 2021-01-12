class Solution {
    class Node {
        String str;
        boolean visited;

        Node(String str) {
            this.str = str;
            this.visited = false;
        }
    }

    boolean connectable(String a, String b) {
        int diff = 0;

        for (int i = 0; i < a.length(); i++) {
            if (a.charAt(i) != b.charAt(i)) diff++;
            if (diff > 1) break;
        }

        return diff == 1;
    }

    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        HashMap<String, List<Node>> graph = new HashMap<>();
        HashMap<String, Node> nodes = new HashMap<>();

        wordList.add(beginWord);
        for (int i = 0; i < wordList.size(); i++) {
            nodes.put(wordList.get(i), new Node(wordList.get(i)));
        }

        for (int i = 0; i < wordList.size(); i++) {
            List<Node> neighbor = new ArrayList<>();
            for (int j = 0; j < wordList.size(); j++) {
                if (i == j) continue;
                if (!connectable(wordList.get(i), wordList.get(j))) continue;
                neighbor.add(nodes.get(wordList.get(j)));
            }
            graph.put(wordList.get(i), neighbor);
        }

        Node start = nodes.get(beginWord);
        List<Node> openList = new ArrayList<>();
        List<Node> nextList = new ArrayList<>();
        openList.add(start);
        start.visited = true;

        int ret = 1;
        while (true) {
            if (openList.isEmpty()) {
                if (nextList.isEmpty()) {
                    return 0;
                }
                openList = nextList;
                nextList = new ArrayList<>();
                ret++;
            }

            Node curr = openList.get(0);
            if (curr.str.equals(endWord)) {
                return ret;
            }
            openList.remove(0);

            List<Node> neighbors = graph.get(curr.str);
            for (Node neighbor : neighbors) {
                if (neighbor.visited) continue;
                nextList.add(neighbor);
                neighbor.visited = true;
            }
        }
    }
}
