

// Complete the sortedInsert function below.

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     int data;
 *     DoublyLinkedListNode* next;
 *     DoublyLinkedListNode* prev;
 * };
 *
 */
DoublyLinkedListNode* sortedInsert(DoublyLinkedListNode* head, int data) {
    DoublyLinkedListNode *newNode = new DoublyLinkedListNode(data);
    DoublyLinkedListNode *curr = head, *prev;

    if(data <= head->data){
        newNode->next = head;
        head->prev = newNode;
        head = newNode;
    }
    else{
        bool inserted = false;
        while(curr)
        {
            if(curr->data >= data){
                newNode->prev = curr->prev;
                newNode->next = curr;
                curr->prev->next = newNode;
                curr->prev = newNode;
                inserted = true;
                break;
            }
            prev = curr;
            curr = curr->next;
        }
        if(!inserted){
            prev->next = newNode;
            newNode->prev = prev;
            
        }
    }

    return head;
}

