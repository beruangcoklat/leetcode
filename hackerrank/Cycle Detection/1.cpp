

// Complete the has_cycle function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */
bool has_cycle(SinglyLinkedListNode* head) {
    map<SinglyLinkedListNode*, int> dict;
    SinglyLinkedListNode *curr = head;
    while(curr)
    {
        if(dict[curr]){
            return true;
        }
        dict[curr] = 1;
        curr = curr->next;
    }
    return false;
}

