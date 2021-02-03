class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        Arrays.sort(arr);
        List<List<Integer>> ret = new ArrayList();
        HashMap<Integer, List<List<Integer>>> map = new HashMap();
        HashMap<String, Boolean> mark = new HashMap<>();
        int minDiff = Integer.MAX_VALUE;
        for (int i = 0; i < arr.length - 1; i++) {
            int max = arr[i + 1];
            int min = arr[i];

            String key = min + "_" + max;
            if (mark.containsKey(key)) continue;
            mark.put(key, true);

            int diff = max - min;
            minDiff = Math.min(minDiff, diff);
            List<Integer> pair = new ArrayList<>();
            pair.add(min);
            pair.add(max);

            List<List<Integer>> allList = map.get(diff);
            boolean ok = allList != null;
            if (!ok) allList = new ArrayList<>();
            allList.add(pair);
            if (!ok) map.put(diff, allList);
        }

        return map.get(minDiff);
    }
}