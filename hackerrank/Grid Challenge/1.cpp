#include <bits/stdc++.h>

using namespace std;

string gridChallenge(vector<string> grid) {
  for (int i = 0; i < grid.size(); i++) {
    sort(grid[i].begin(), grid[i].end());
  }

  for (int i = 0; i < grid.size() - 1; i++) {
    for (int j = 0; j < grid[i].length(); j++) {
      char curr = grid[i][j];
      char bawah = grid[i + 1][j];
      if (curr > bawah) {
        return "NO";
      }
    }
  }
  return "YES";
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int t;
    cin >> t;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int t_itr = 0; t_itr < t; t_itr++) {
        int n;
        cin >> n;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        vector<string> grid(n);

        for (int i = 0; i < n; i++) {
            string grid_item;
            getline(cin, grid_item);

            grid[i] = grid_item;
        }

        string result = gridChallenge(grid);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}
