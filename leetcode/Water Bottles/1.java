class Solution {
    public int numWaterBottles(int numBottles, int numExchange) {
        int result = numBottles;
        while (numBottles >= numExchange) {
            int remainingBottles = numBottles % numExchange;
            numBottles /= numExchange;
            result += numBottles;
            numBottles += remainingBottles;
        }
        return result;
    }
}