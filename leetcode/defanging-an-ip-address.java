class Solution {
    public String defangIPaddr(String address) {
        String[]split = address.split("[.]");
        if(split.length == 0) return "";
        String ret = split[0];
        for(int i=1 ; i<split.length ; i++){
            ret += "[.]";
            ret += split[i];
        }
        return ret;
    }
}
