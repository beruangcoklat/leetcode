#include <bits/stdc++.h>

using namespace std;

int countA(string x) {
  int ret = 0;
  for (int i = 0; i < x.length(); i++) {
    if (x.at(i) == 'a') {
      ret++;
    }
  }
  return ret;
}

long long repeatedString(string s, long long n) {
  int len = s.length();
  long long words = n / len;

  int sisa = n % len;
  string substring = s.substr(0, sisa);

  long long aaa = words * countA(s);
  long long bbb = countA(substring);

  return aaa + bbb;
}
int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    long n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    long long result = repeatedString(s, n);

    fout << result << "\n";

    fout.close();

    return 0;
}
