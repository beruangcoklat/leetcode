class Solution {
    public int brokenCalc(int X, int Y) {
        int ret = 0;
        while(X != Y){
            ret++;
            if(X >= Y){
                ret += (X - Y) - 1;
                break;
            }
            if(Y % 2 == 1){
                Y++;
            }else{
                Y /= 2;
            }
        }
        
        return ret;
    }
}
