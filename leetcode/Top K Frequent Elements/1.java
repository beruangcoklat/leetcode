class Solution {
    class Node {
        int value;
        int count;

        Node(int value, int count) {
            this.value = value;
            this.count = count;
        }
    }

    public int[] topKFrequent(int[] nums, int k) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int a : nums) {
            Integer count = map.get(a);
            if (count == null) count = 0;
            map.put(a, count + 1);
        }

        List<Node> nodes = new ArrayList<>();
        for (Map.Entry<Integer, Integer> e : map.entrySet()) {
            nodes.add(new Node(e.getKey(), e.getValue()));
        }

        Collections.sort(nodes, (Node a, Node b) -> {
            return b.count - a.count;
        });

        int ret[] = new int[k];
        for (int i = 0; i < k; i++) {
            ret[i] = nodes.get(i).value;
        }

        return ret;
    }
}