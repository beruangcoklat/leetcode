class Solution {
    fun largestLocal(grid: Array<IntArray>): Array<IntArray> {
        val ans: Array<IntArray> = Array<IntArray>(grid.size - 2) { intArrayOf() }
        for (y in 1..grid.size - 2) {
            var row = IntArray(grid.size - 2)
            for (x in 1..grid[0].size - 2) {
                row[x - 1] = getMax(grid, x, y)
            }
            ans[y - 1] = row
        }
        return ans
    }

    private fun getMax(grid: Array<IntArray>, x: Int, y: Int): Int {
        var max = grid[y][x]
        for (y in y - 1..y + 1) {
            for (x in x - 1..x + 1) {
                if (grid[y][x] > max) {
                    max = grid[y][x]
                }
            }
        }
        return max
    }
}