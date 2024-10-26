class Solution {
    fun minBitFlips(start: Int, goal: Int): Int {
        val startBinary = decimalToBinary(start)
        var goalBinary = decimalToBinary(goal)
        var ans = 0
        var i = 1
        while (startBinary.length - i >= 0 && goalBinary.length - i >= 0) {
            if (startBinary[startBinary.length - i] != goalBinary[goalBinary.length - i]) {
                println(i)
                ans++
            }
            i++
        }
        while (startBinary.length - i >= 0) {
            if (startBinary[startBinary.length - i] == '1') {
                ans++
            }
            i++
        }
        while (goalBinary.length - i >= 0) {
            if (goalBinary[goalBinary.length - i] == '1') {
                ans++
            }
            i++
        }
        return ans
    }

    private fun decimalToBinary(_decimal: Int): String {
        var ans = ""
        var decimal = _decimal
        while (decimal > 0) {
            ans = (if (decimal % 2 == 0) "0" else "1") + ans
            decimal = decimal shr 1
        }
        return ans
    }
}