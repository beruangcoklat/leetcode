class Solution {
    int getValueFromIdx(ArrayList<Integer> values, ArrayList<Integer> counts, int idx) {
        for (int i = 0; i < values.size(); i++) {
            int value = values.get(i);
            int count = counts.get(i);

            idx -= count;
            if (idx <= 0) {
                return value;
            }
        }

        return -1;
    }

    public double[] sampleStats(int[] count) {
        double min = 0, max = 0, mode = 0, mean = 0, median = 0;
        boolean first = true;
        long totalValue = 0;
        int totalData = 0;
        int modeCount = 0;
        ArrayList<Integer> values = new ArrayList<>();
        ArrayList<Integer> counts = new ArrayList<>();

        for (int i = 0; i < count.length; i++) {
            if (count[i] == 0) continue;

            values.add(i);
            counts.add(count[i]);

            totalData += count[i];
            totalValue += (i * count[i]);

            max = i;

            if (first) {
                first = false;
                min = i;
            }

            if (modeCount < count[i]) {
                modeCount = count[i];
                mode = i;
            }
        }

        mean = (double) totalValue / totalData;

        if (totalData % 2 == 1) {
            median = getValueFromIdx(values, counts, (totalData / 2) + 1);
        } else {
            median += getValueFromIdx(values, counts, (totalData / 2) + 1);
            median += getValueFromIdx(values, counts, totalData / 2);
            median /= 2;
        }

        return new double[]{min, max, mean, median, mode};
    }
}