class Solution {
    public int numberOfMatches(int n) {
        int ret = 0;
        while(n > 1){
            ret += (n / 2);
            if(n % 2 == 0){
                n /= 2;
            }else{
                n = (n / 2)+1;
            }
        }
        return ret;
    }
}
