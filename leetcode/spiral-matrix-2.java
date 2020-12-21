class Solution {
    
    boolean valid(int[][] arr, int x, int y,  int n){
        if(x < 0 || x >= n) return false;
        if(y < 0 || y >= n) return false;
        if(arr[y][x] != 0) return false;
        return true;
    }
    
    public int[][] generateMatrix(int n) {
        int [][] arr = new int[n][n];
        int x = 0;
        int y = 0;
        
        int dir = 0;
        /*
        0 kanan
        1 bawah
        2 kiri
        3 atas
        */
        
        int []dirX = {1,0,-1,0};
        int []dirY = {0,1,0,-1};
        
        int counter = 1;
        for(int i=0 ; i<n*n ; i++){
            arr[y][x] = counter;
            counter++;
            
            int nx = x + dirX[dir];
            int ny = y + dirY[dir];
            if(valid(arr, nx, ny, n)){
                x = nx;
                y = ny;
                continue;
            }
            
            dir = (dir + 1) % 4;
            x += dirX[dir];
            y += dirY[dir];
        }
        
        return arr;
    }
}
