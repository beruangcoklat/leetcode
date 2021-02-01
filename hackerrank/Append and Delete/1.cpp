#include <bits/stdc++.h>

using namespace std;

bool appendAndDelete(string s, string t, int k) {
  int slen = s.length();
  int tlen = t.length();

  int max = slen > tlen ? slen : tlen;
  int min = slen < tlen ? slen : tlen;

  int common = 0;
  for (int i = 0; i < min; i++) {
    if (s[i] == t[i]) {
      common++;
      continue;
    }
    break;
  }

  int dif = max - common;
  int need = (tlen - common) + (slen - common);

  if (dif > k)
    return false;
  if (need > k)
    return false;

  if ((k - dif) % 2 == 0)
    return true;
  if ((k - need) % 2 == 0)
    return true;

  if (k >= slen + tlen)
    return true;

  return false;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    string t;
    getline(cin, t);

    int k;
    cin >> k;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string result = appendAndDelete(s, t, k) ? "Yes" : "No";

    fout << result << "\n";

    fout.close();

    return 0;
}
