class Solution {
    public int[] kWeakestRows(int[][] mat, int k) {
        int sorted[] = new int[mat.length];
        int ret[] = new int[k];
        HashMap<Integer, List<Integer>> map = new HashMap<>();

        int idx = -1;
        for (int[] row : mat) {
            idx++;
            for (int c : row) {
                sorted[idx] += c;
            }
            List<Integer> list = map.get(sorted[idx]);
            if (list == null) {
                list = new ArrayList<>();
                list.add(idx);
                map.put(sorted[idx], list);
            } else {
                list.add(idx);
            }
        }
        Arrays.sort(sorted);

        List<Integer> tes = new ArrayList<>();
        for(Map.Entry<Integer,List<Integer>> e: map.entrySet()){
            tes.add(e.getKey());
        }

        Collections.sort(tes);

        idx = 0;
        for (int i = 0; i < tes.size(); i++) {
            List<Integer> list = map.get(tes.get(i));
            for (int c : list) {
                ret[idx++] = c;
                if (idx == k) return ret;
            }
        }

        return ret;
    }
}
