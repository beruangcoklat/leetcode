class Solution {
    public int maximumScore(int a, int b, int c) {
        int arr[] = {a, b, c};

        int ret = 0;
        while (true) {
            ret++;

            Arrays.sort(arr);
            arr[arr.length - 1]--;
            arr[arr.length - 2]--;

            if (arr[arr.length - 1] == 0) {
                return ret;
            }

            if (arr[arr.length - 2] == 0) {
                if (arr.length == 2) {
                    return ret;
                }
                int newarr[] = new int[2];
                newarr[0] = arr[0];
                newarr[1] = arr[2];
                arr = newarr;
            }
        }
    }
}
