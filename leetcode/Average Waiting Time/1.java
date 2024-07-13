class Solution {
    public double averageWaitingTime(int[][] customers) {
        int[] firstCustomer = customers[0];
        double waitingTime = firstCustomer[1];
        int chefReadyTime = firstCustomer[0] + firstCustomer[1];
        for (int i = 1; i < customers.length; i++) {
            int[] customer = customers[i];
            int arrival = customer[0];
            int time = customer[1];

            int finishTime = Math.max(arrival, chefReadyTime) + time;
            chefReadyTime = finishTime;
            waitingTime += finishTime - arrival;
        }
        return waitingTime / customers.length;
    }
}