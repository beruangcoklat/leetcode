#include <bits/stdc++.h>

using namespace std;

vector<string> split_string(string);

int memo[100004];
bool memoBool[100004];

int solve(vector<int> const &arr, int currIndex){
    if(currIndex >= arr.size()){
        return 0;
    }
    
    if(memoBool[currIndex]){
        return memo[currIndex];
    }
    
    int take = solve(arr, currIndex + 2) + arr[currIndex];
    int dontTake = solve(arr, currIndex + 1);
    
    int max = take > dontTake ? take : dontTake;
    memo[currIndex] = max;
    memoBool[currIndex] = true;
    return max;
}

int maxSubsetSum(vector<int> arr) {
    return solve(arr, 0);
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string arr_temp_temp;
    getline(cin, arr_temp_temp);

    vector<string> arr_temp = split_string(arr_temp_temp);

    vector<int> arr(n);

    for (int i = 0; i < n; i++) {
        int arr_item = stoi(arr_temp[i]);

        arr[i] = arr_item;
    }

    int res = maxSubsetSum(arr);

    fout << res << "\n";

    fout.close();

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}
