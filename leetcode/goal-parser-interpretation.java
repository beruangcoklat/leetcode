class Solution {
    public String interpret(String command) {
        String ret = "";
        while(command.length() > 0){
            if(command.charAt(0) == 'G'){
                command = command.substring(1, command.length());
                ret += "G";
                continue;
            }
            
            if(command.substring(0, 2).equals("()")){
                command = command.substring(2, command.length());
                ret += "o";
                continue;
            }
            
            if(command.substring(0, 4).equals("(al)")){
                command = command.substring(4, command.length());
                ret += "al";
                continue;
            }
        }
        return ret;
    }
}
