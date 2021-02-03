#include <bits/stdc++.h>

using namespace std;

string multiply(string a, string b) {
  string rows[a.length()];
  int carier = 0;

  for (int i = a.length() - 1; i >= 0; i--) {
    carier = 0;
    for (int j = 0; j < a.length() - i - 1; j++)
      rows[i] += "0";

    for (int j = b.length() - 1; j >= 0; j--) {
      int num1 = a[i] - '0';
      int num2 = b[j] - '0';
      int total = (num1 * num2) + carier;
      carier = 0;
      if (total > 9 && j != 0) {
        carier = total / 10;
        total %= 10;
      }
      string result = to_string(total);
      rows[i] = result + rows[i];
    }
  }

  string ret = "";
  carier = 0;
  for (int i = 0; i < rows[0].length(); i++) {
    int total = carier;
    carier = 0;

    for (int j = 0; j < a.length(); j++) {
      int idx = rows[j].length() - 1 - i;
      if (idx >= 0) {
        total += rows[j][idx] - '0';
      }
    }

    if (total > 9 && i != rows[0].length() - 1) {
      carier = total / 10;
      total %= 10;
    }
    ret = to_string(total) + ret;
  }

  return ret;
}

void extraLongFactorials(int n) {
  string ret = "1";
  for (int i = 2; i < n + 1; i++) {
    ret = multiply(ret, to_string(i));
  }
  cout << ret;
}

int main()
{
    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    extraLongFactorials(n);

    return 0;
}
