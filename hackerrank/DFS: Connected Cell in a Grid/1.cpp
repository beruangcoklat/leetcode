#include <bits/stdc++.h>

using namespace std;

int max = 0;
int realMax = 0;

void dfs(int x, int y, vector<vector<int>> &grid, vector<vector<bool>> &visited){
    int dirX[] = {0,0,1,-1,1,1,-1,-1};
    int dirY[] = {-1,1,0,0,1,-1,1,-1};

    visited[y][x] = true;
    ::max++;


    for(int i=0 ; i<8 ; i++){
        int nx = x + dirX[i];
        int ny = y + dirY[i];

        if(nx < 0 || nx >= grid[0].size()) continue;
        if(ny < 0 || ny >= grid.size()) continue;
        
        if(visited[ny][nx]) continue;
        if(grid[ny][nx] == 0) continue;
        visited[ny][nx] = false;
        
        dfs(nx, ny, grid, visited);
    }
}

int maxRegion(vector<vector<int>> grid) {
    vector<vector<bool>> visited;
    for(int i=0 ; i<grid.size() ; i++){
        visited.push_back(vector<bool>());
        for(int j=0 ; j<grid[i].size() ; j++){
            visited[i].push_back(false);
        }
    }     
    
    ::max = 0;
    for(int i=0 ; i<grid.size() ; i++){
        for(int j=0 ; j<grid[i].size() ; j++){
        
//            cout << ::max << endl;
            ::realMax = :: realMax > ::max ? ::realMax : ::max;
            
            if(grid[i][j] == 0) {    
                continue;
            }
            if(visited[i][j]) {
                   continue;
            }
         
            ::max = 0;
            
            dfs(j, i, grid, visited);
        }
    }

    return ::realMax;
}



int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    int m;
    cin >> m;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<vector<int>> grid(n);
    for (int i = 0; i < n; i++) {
        grid[i].resize(m);

        for (int j = 0; j < m; j++) {
            cin >> grid[i][j];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    int res = maxRegion(grid);

    fout << res << "\n";

    fout.close();

    return 0;
}
