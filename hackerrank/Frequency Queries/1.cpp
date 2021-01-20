#include <bits/stdc++.h>
using namespace std;

int main() {
  int q;
  scanf("%d", &q);

  map<int, int> valueFreq, freqFreq;

  for (int i = 0; i < q; i++) {
    int type, num;
    scanf("%d %d", &type, &num);

    if (type == 1) {
      int old = valueFreq[num]++;
      int curr = old + 1;
      freqFreq[curr]++;
      //if (freqFreq[old] > 0) {
        freqFreq[old]--;
      //}
    } else if (type == 2) {
      int old = valueFreq[num];
      bool found = old > 0;
      if (found) {
        int curr = old - 1;
        valueFreq[num]--;
        freqFreq[curr]++;
        //if (freqFreq[old] > 0) {
          freqFreq[old]--;
        //}
      }
    } else if (type == 3) {
      int result = freqFreq[num] > 0 ? 1 : 0;
      printf("%d\n", result);
    }
  }

  return 0;
}
