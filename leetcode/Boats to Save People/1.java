class Solution {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int left = 0;
        int right = people.length - 1;
        int ret = 0;
        while (left <= right) {
            int a = people[left];
            int b = people[right];

            if (a + b <= limit) {
                left++;
                right--;
            } else {
                right--;
            }
            ret++;
        }
        return ret;
    }
}
