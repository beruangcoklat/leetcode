

// Complete the findMergeNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */
int findMergeNode(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {
    map<SinglyLinkedListNode*, int> dict;
    
    SinglyLinkedListNode* curr = head1;
    while(curr){
        dict[curr] = 1;
        curr = curr->next;
    }

    curr = head2;
    while(curr){
        if(dict[curr]){
            return curr->data;
        }
        curr = curr->next;
    }

    cout << "not found";
    return -1;
}

