class Solution {
    public boolean isValid(String s) {
        List<Character> list = new ArrayList();
        for(char a : s.toCharArray()){
            if(a == '(' || a == '[' || a == '{'){
                list.add(a);
                continue;
            }
            
            if(list.size() == 0){
                return false;
            }
            char last = list.get(list.size()-1);
            if(a == ')') {
                if(last != '('){
                    return false;
                }
                list.remove(list.size()-1);
                continue;
            }
            if(a == ']') {
                if(last != '['){
                    return false;
                }
                list.remove(list.size()-1);
                continue;
            }
            if(a == '}') {
                if(last != '{'){
                    return false;
                }
                list.remove(list.size()-1);
                continue;
            }
        }
        return list.size() == 0;
    }
}
