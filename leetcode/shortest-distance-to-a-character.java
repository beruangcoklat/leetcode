class Solution {
    
    int dist(String s, char c, int idx){        
        int inc = 0;
        
        while(true){
            if(idx + inc < s.length()){
                if(s.charAt(idx + inc) == c){
                    return inc;
                }
            }
            if(idx - inc >= 0){
                if(s.charAt(idx - inc) == c){
                    return inc;
                }
            }
            inc++;
        }
    }
    
    public int[] shortestToChar(String S, char C) {
        int []arr = new int[S.length()];
        for(int i=0 ; i<S.length() ; i++){
            arr[i] = dist(S, C, i);
        }
        return arr;
    }
}
