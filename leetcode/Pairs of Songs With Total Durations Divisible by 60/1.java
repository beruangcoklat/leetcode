class Solution {
    
    void incr(HashMap<Integer, Integer> map, int key){
        int v = 1;
        if(map.containsKey(key)) v += map.get(key);
        map.put(key, v);
    }
    
    int normalize(int x){
        if(x == 60) return 0;
        return x;
    }
    
    public int numPairsDivisibleBy60(int[] time) {
        int ret = 0;
        HashMap<Integer, Integer> map = new HashMap();
        
        for(int t : time){
            int key = normalize(60 - t % 60);
            if(map.containsKey(key)){
                ret += map.get(key);
            }
            incr(map, t % 60);
        }
        
        return ret;
    }
}
