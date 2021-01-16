#include <bits/stdc++.h>

using namespace std;

vector<string> getSubstring(string s) {
  vector<string> ret;
  for (int i = 1; i < s.length(); i++) {
    for (int j = 0; j < s.length(); j++) {
      string curr = s.substr(j, i);
      if (curr.length() != i)
        break;
      ret.push_back(curr);
    }
  }
  return ret;
}

bool isAnagram(string a, string b) {
  if (a.length() != b.length())
    return false;

  int arr[26] = {0};

  for (int i = 0; i < a.length(); i++) {
    arr[a[i] - 'a']++;
    arr[b[i] - 'a']--;
  }
  for (int i = 0; i < 26; i++) {
    if (arr[i] != 0)
      return false;
  }
  return true;
}

// Complete the sherlockAndAnagrams function below.
int sherlockAndAnagrams(string s) {
  vector<string> substrings = getSubstring(s);

  int ret = 0;
  for (int i = 0; i < substrings.size(); i++) {
    for (int j = i + 1; j < substrings.size(); j++) {
      string a = substrings[i];
      string b = substrings[j];
      if (a.length() != b.length())
        break;
      if (isAnagram(a, b)) {
        cout << a << " - " << b << endl;
        ret++;
      }
    }
  }
  return ret;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int q;
    cin >> q;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int q_itr = 0; q_itr < q; q_itr++) {
        string s;
        getline(cin, s);

        int result = sherlockAndAnagrams(s);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}
