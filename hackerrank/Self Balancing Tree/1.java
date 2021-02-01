	/* Class node is defined as :
    class Node 
    	int val;	//Value
    	int ht;		//Height
    	Node left;	//Left child
    	Node right;	//Right child

	*/

    static Node newNode(int val, int ht){
        Node n = new Node();
        n.val = val;
        n.ht = ht;
        n.right = null;
        n.left = null;
        return n;
    }

    static int height(Node root){
        return root == null ? -1 : root.ht;
    }

    static Node rightRotate(Node root){
        Node left = root.left;
        Node leftRight = left.right;
        
        left.right = root;
        root.left = leftRight;
        
        root.ht = Math.max(height(root.right), height(root.left)) + 1;
        left.ht = Math.max(height(left.right), height(left.left)) + 1;
        
        return left;
    }

    static Node leftRotate(Node root){
        Node right = root.right;
        Node rightLeft = right.left;
        
        right.left = root;
        root.right = rightLeft;
        
        root.ht = Math.max(height(root.right), height(root.left)) + 1;
        right.ht = Math.max(height(right.right), height(right.left)) + 1;
        
        return right;
    }

	static Node insert(Node root,int val)
    {
    	if(root == null) return newNode(val, 0);
        if(val > root.val) root.right = insert(root.right, val);
        else root.left = insert(root.left, val);
        
        int left = height(root.left);
        int right = height(root.right);
        
        root.ht = Math.max(left, right) + 1;
        
        int balance = Math.abs(left - right);
        if(balance < 2) return root;
        
        if(right > left){
            if(val > root.right.val){
                // left
                return leftRotate(root);
            }else{
                // right left
                root.right = rightRotate(root.right);
                return leftRotate(root);
            }
        }else{
            if(val < root.left.val){
                // right
                return rightRotate(root);
            }else{
                // left right
                root.left = leftRotate(root.left);
                return rightRotate(root);
            }
        }
    }
