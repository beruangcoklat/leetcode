class Solution {
    public boolean carPooling(int[][] trips, int capacity) {
        HashMap<Integer, Integer> memo = new HashMap();
        for(int i=0 ; i<trips.length ; i++){
            int numsPassengers = trips[i][0];
            int startLocation = trips[i][1];
            int endLocation = trips[i][2];

            for(int j=startLocation ; j< endLocation ; j++){
                int total = numsPassengers;
                if(memo.containsKey(j)){
                    total += memo.get(j);
                }
                if(total > capacity){
                    return false;
                }
                memo.put(j, total);
            }
        }
        return true;
    }
}