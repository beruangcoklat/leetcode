#include <bits/stdc++.h>

using namespace std;

// Complete the hourglassSum function below.
int hourglassSum(vector<vector<int>> arr) {
    int dirX[] = {0,-1,0,1,-1,0,1};
    int dirY[] = {0,-1,-1,-1,1,1,1};
    bool first = true;
    int max = 0;

    for(int y=1 ; y<5 ; y++){
        for(int x=1 ; x<5 ; x++){
            int curr = 0;
            for(int i=0 ; i<7 ; i++){
                curr += arr[y + dirY[i]][x + dirX[i]];
            }

            if(first){
                max = curr;
                first = false;
            }
            else{
                max = max > curr ? max : curr;
            }
        }
    }

    return max;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    vector<vector<int>> arr(6);
    for (int i = 0; i < 6; i++) {
        arr[i].resize(6);

        for (int j = 0; j < 6; j++) {
            cin >> arr[i][j];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    int result = hourglassSum(arr);

    fout << result << "\n";

    fout.close();

    return 0;
}

