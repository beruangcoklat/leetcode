class Solution {
    public boolean carPooling(int[][] trips, int capacity) {
        int arr[] = new int[1001];
        for(int i=0 ; i<trips.length ; i++){
            int numsPassengers = trips[i][0];
            int startLocation = trips[i][1];
            int endLocation = trips[i][2];
            
            arr[startLocation] += numsPassengers;
            arr[endLocation] -= numsPassengers;
        }
        int curr = 0;
        for(int i=0 ; i<arr.length ; i++){
            curr += arr[i];
            if(curr > capacity){
                return false;
            }
        }
        return true;
    }
}
