class Solution {
    public boolean validMountainArray(int[] arr) {
        boolean sudahNaik = false;
        boolean sudahTurun = false;
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] == arr[i + 1]) return false;

            boolean curr = arr[i + 1] > arr[i];
            if (curr) {
                sudahNaik = true;
            } else {
                sudahTurun = true;
            }

//            udah naik dan turun, terus naik lagi
            if (sudahNaik && sudahTurun && curr) {
                return false;
            }

//            belum naik udah turun
            if (!sudahNaik && !curr) {
                return false;
            }
        }

        return sudahNaik && sudahTurun;
    }
}
