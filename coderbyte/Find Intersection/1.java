import java.util.*; 
import java.io.*;

class Main {

  public static String FindIntersection(String[] strArr) {
    HashMap<String, Boolean> map = new HashMap();
    String ret = "";
    boolean first =  true;

    for(String str : strArr){
      String[] sp = str.replaceAll(" ", "").split(",");
      for(String s : sp){
        if(map.containsKey(s)){
          if(first){
            first = false;
          }else{
            ret += ",";
          }
          ret += s;
        }else{
          map.put(s, true); 
        }
      }
    }

    if(first) ret = "false";
    return ret;
  }

  public static void main (String[] args) {  
    // keep this function call here     
    Scanner s = new Scanner(System.in);
    System.out.print(FindIntersection(s.nextLine())); 
  }

}
