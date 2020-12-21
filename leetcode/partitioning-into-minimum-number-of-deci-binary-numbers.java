class Solution {
    public int minPartitions(String n) {
        int max = -1;
        for(int i=0 ; i<n.length() ; i++){
            if (n.charAt(i) > max){
                max = n.charAt(i);
            }
        }
        return max - '0';
    }
}
