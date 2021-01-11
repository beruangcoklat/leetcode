class Solution {
    
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int idx1 = 0, idx2 = 0;
        ArrayList<Integer> arr = new ArrayList<>();
        while (idx1 < m && idx2 < n) {
            int min = 0;
            if (nums1[idx1] < nums2[idx2]) {
                min = nums1[idx1];
                idx1++;
            } else {
                min = nums2[idx2];
                idx2++;
            }
            arr.add(min);
        }

        while (idx1 < m) {
            arr.add(nums1[idx1]);
            idx1++;
        }

        while (idx2 < n) {
            arr.add(nums2[idx2]);
            idx2++;
        }

        for (int i = 0; i < nums1.length; i++) {
            nums1[i] = arr.get(i);
        }
    }
    
}
