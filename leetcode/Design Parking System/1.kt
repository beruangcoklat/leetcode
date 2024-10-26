class ParkingSystem(big: Int, medium: Int, small: Int) {
    private val big = big
    private val medium = medium
    private val small = small
    private var currBig = 0
    private var currMedium = 0
    private var currSmall = 0

    fun addCar(carType: Int): Boolean {
        when (carType) {
            1 -> {
                if (currBig == big) {
                    return false
                }
                currBig++
            }

            2 -> {
                if (currMedium == medium) {
                    return false
                }
                currMedium++
            }

            3 -> {
                if (currSmall == small) {
                    return false
                }
                currSmall++
            }
        }
        return true
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * var obj = ParkingSystem(big, medium, small)
 * var param_1 = obj.addCar(carType)
 */