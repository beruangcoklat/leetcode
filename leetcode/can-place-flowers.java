class Solution {
    
    boolean check(int[] flowerbed, int idx){
        if(flowerbed[idx] == 1) return false;
        if(idx-1 >= 0 && flowerbed[idx-1] == 1) return false;
        if(idx+1 < flowerbed.length && flowerbed[idx+1] == 1) return false;
        return true;
    }
    
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        if(n == 0){
            return true;
        }
        for(int i=0 ; i<flowerbed.length ; i++){
            if(check(flowerbed, i)){
                flowerbed[i] = 1;
                n--;
            }
            if(n == 0){
                return true;
            }
        }
        return false;
    }
}
