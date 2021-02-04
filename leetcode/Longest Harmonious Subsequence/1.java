class Solution {
    public int findLHS(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>();
        List<Integer> distinct = new ArrayList<>();
        for (int a : nums) {
            Integer count = map.get(a);
            if (count == null) {
                count = 0;
                distinct.add(a);
            }
            map.put(a, count + 1);
        }

        Collections.sort(distinct);
        int max = 0;
        for(int i=0 ; i<distinct.size() - 1 ; i++) {
            int a = distinct.get(i);
            int b = distinct.get(i + 1);
            if (b - a != 1) continue;
            max = Math.max(max, map.get(a) + map.get(b));
        }
        return max;
    }
}
