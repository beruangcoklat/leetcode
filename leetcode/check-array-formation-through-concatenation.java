class Solution {
    
    String convert(int[] arr) {
        String ret = "";
        for (int i = 0; i < arr.length; i++) {
            ret += ("_" + arr[i] + "_");
        }
        return ret;
    }

    public boolean canFormArray(int[] arr, int[][] pieces) {
        String arrstr = convert(arr);

        for (int i = 0; i < pieces.length; i++) {
            String curr = convert(pieces[i]);

            arrstr = arrstr.replace(curr, "#");
            if (arrstr.length() == 0) {
                return true;
            }
        }

        return arrstr.replace("#", "").length() == 0;
    }
}
