#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
    int n;
    
    int idxMax = 0;
    vector<int> max, stack;
    
    int query, value;
    
    scanf("%d", &n);
    for(int i=0 ; i<n ; i++){
        scanf("%d", &query);
        if(query == 1){
            scanf("%d", &value);
            stack.push_back(value);
            
            if(idxMax == 0){
                max.push_back(value);
            }else{   
                int temp = value > max[idxMax-1] ? value : max[idxMax-1];
                max.push_back(temp);
            }
            idxMax++;
        }else if(query == 2 && idxMax > 0){
            idxMax--;
            stack.pop_back();
            max.pop_back();
        }else{
            if(idxMax == 0) printf("0\n");
            else printf("%d\n", max[idxMax-1]);
        }
    }
    return 0;
}
