/* Hidden stub code will pass a root argument to the function below. Complete the function to solve the challenge. Hint: you may want to write one or more helper functions.  

The Node class is defined as follows:
    class Node {
        int data;
        Node left;
        Node right;
     }
*/
    
    List<Integer> list;
    int last = 0;
    boolean valid;

    void inorder(Node root){
        if(root == null) return;
        inorder(root.left);
        if(last >= root.data){
            valid = false;
            return;
        }
        last = root.data;
        inorder(root.right);
    }

    boolean checkBST(Node root) {
        list = new ArrayList();
        valid = true;
        inorder(root);
        return valid;
    }