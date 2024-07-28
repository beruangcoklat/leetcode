class Solution {
    class Pair {
        String name;
        int height;
        Pair(String name, int height) {
            this.name = name;
            this.height = height;
        }
    }
    public String[] sortPeople(String[] names, int[] heights) {
        List<Pair> pairs = new ArrayList<>();
        for (int i = 0; i < names.length; i++) {
            pairs.add(new Pair(names[i], heights[i]));
        }

        Collections.sort(pairs, (a, b) -> {
            return b.height - a.height;
        });

        return pairs.stream().map(p -> p.name).toArray(String[]::new);
    }
}