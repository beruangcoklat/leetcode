class Solution {
    class Node {
        String str;
        boolean visited;
        List<Node> parent;

        Node(String str) {
            this.str = str;
            this.visited = false;
            this.parent = new ArrayList<>();
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

    List<String> cloneList(List<String> nodes) {
        List<String> ret = new ArrayList<>();
        for (String n : nodes) {
            ret.add(n);
        }
        return ret;
    }

    List<Node> removeDuplicateList(List<Node> nodes) {
        HashMap<Node, Boolean> map = new HashMap<>();
        List<Node> ret = new ArrayList<>();
        for (Node n : nodes) {
            if (map.containsKey(n)) continue;
            ret.add(n);
            map.put(n, true);
        }
        return ret;
    }

    void recursive(List<List<String>> allResult, List<String> result, Node node, HashMap<Node, Boolean> visited) {
        visited.put(node, true);
        result.add(node.str);
        if (node.parent.isEmpty()) {
            allResult.add(result);
            return;
        }

        for (Node n : node.parent) {
            if (visited.containsKey(n)) continue;
            List<String> cloneResult = cloneList(result);
            recursive(allResult, cloneResult, n, (HashMap<Node, Boolean>) visited.clone());
        }
    }

    List<List<String>> getPath(Node node) {
        List<List<String>> ret = new ArrayList<>();
        recursive(ret, new ArrayList<>(), node, new HashMap<Node, Boolean>());
        return ret;
    }

    List<String> reverseList(List<String> list) {
        List<String> ret = new ArrayList<>();
        for (int i = list.size() - 1; i >= 0; i--) ret.add(list.get(i));
        return ret;
    }

    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        HashMap<String, List<Node>> graph = new HashMap<>();
        HashMap<String, Node> nodes = new HashMap<>();

        boolean found = false;
        for (int i = 0; i < wordList.size(); i++) {
            if (wordList.get(i).equals(beginWord)) {
                found = true;
            }
            nodes.put(wordList.get(i), new Node(wordList.get(i)));
        }
        if (!found) {
            nodes.put(beginWord, new Node(beginWord));
            wordList.add(beginWord);
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
        start.visited = true;

        HashMap<Node, Boolean> openMap = new HashMap<>();
        List<Node> openList = new ArrayList<>();
        List<Node> nextList = new ArrayList<>();

        openList.add(start);
        openMap.put(start, true);

        List<List<String>> ret = new ArrayList<>();

        while (true) {
            if (openList.isEmpty()) {
                if (!ret.isEmpty()) break;
                if (nextList.isEmpty()) break;

                openList = nextList;
                nextList = new ArrayList<>();

                openMap.clear();
                for (Node a : openList) openMap.put(a, true);
            }

            openList = removeDuplicateList(openList);
            Node curr = openList.get(0);
            openList.remove(0);
            curr.visited = true;

            if (curr.str.equals(endWord)) {
                ret.addAll(getPath(curr));
            }

            List<Node> neighbors = graph.get(curr.str);
            for (Node neighbor : neighbors) {
                if (neighbor.visited) continue;
                if (openMap.containsKey(neighbor)) continue;

                nextList.add(neighbor);
                neighbor.parent.add(curr);
            }
        }

        List<List<String>> reverseRet = new ArrayList<>();
        for (List<String> a : ret) reverseRet.add(reverseList(a));
        return reverseRet;
    }
}
