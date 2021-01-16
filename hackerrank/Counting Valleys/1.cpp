#include <bits/stdc++.h>

using namespace std;

// Complete the countingValleys function below.
int countingValleys(int n, string s) {
    int path = 0;
    vector<int> cumulative, zeroIndex;

    for(int i=0 ; i<s.length() ; i++){
        char curr = s.at(i);
        if(curr == 'D'){
            path--;
        }
        else{
            path++;
        }
        cumulative.push_back(path);
        if(path == 0){
            zeroIndex.push_back(i);
        }
    }

    bool habisTurun, akanTurun;

    int ret = 0;
    for(int i=0 ; i<zeroIndex.size() ; i++){
        int idx = zeroIndex[i];
        int prev = cumulative[idx - 1];
        int next = cumulative[idx + 1];

        bool preCondition = false;
        if(akanTurun){
            preCondition = true;
        }

        habisTurun = prev == 1;
        akanTurun = next == -1;

        if(!habisTurun && preCondition){
            ret++;
        }
    }

    return ret;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string s;
    getline(cin, s);

    int result = countingValleys(n, s);

    fout << result << "\n";

    fout.close();

    return 0;
}
