import java.util.*;
import java.io.*;

class Node {
    Node left;
    Node right;
    int data;
    
    Node(int data) {
        this.data = data;
        left = null;
        right = null;
    }
}

class Solution {

	/*
    class Node 
    	int data;
    	Node left;
    	Node right;
	*/
    
    static List<Integer> find(Node root, int v, List<Integer> list){
        list.add(root.data);
        if(v == root.data) return list;
        if(v > root.data) return find(root.right, v, list);
        return find(root.left, v, list);
    }
    
    static Node findNode(Node root, int v){
        if(v == root.data) return root;
        if(v > root.data) return findNode(root.right, v);
        return findNode(root.left, v);
    }
    
	public static Node lca(Node root, int v1, int v2) {
        List<Integer> a = find(root, v1, new ArrayList());
        List<Integer> b = find(root, v2, new ArrayList());
        HashMap<Integer, Boolean> map = new HashMap();
        
        int aIdx = a.size() - 1;
        int bIdx = b.size() - 1;
        
        while(aIdx > 0 || bIdx > 0){
            if(aIdx > 0){
                int av = a.get(aIdx);
                aIdx--;
                if(map.containsKey(av)){
                    return findNode(root, av);
                }
                map.put(av, true);    
            }
            if(bIdx > 0){
                int bv = b.get(bIdx);
                bIdx--;
                
                if(map.containsKey(bv)){
                    return findNode(root, bv);
                }
                map.put(bv, true);
            }
        }
        return root;
    }

	public static Node insert(Node root, int data) {