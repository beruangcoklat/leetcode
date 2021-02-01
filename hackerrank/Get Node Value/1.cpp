

// Complete the getNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */
int getNode(SinglyLinkedListNode* head, int positionFromTail) {
    while(head){
        SinglyLinkedListNode* top = head;
        for(int i=0 ; i<positionFromTail ; i++) top = top->next;
        if(top->next == 0) return head->data;
        head = head->next;
    }
    return -1;
}

