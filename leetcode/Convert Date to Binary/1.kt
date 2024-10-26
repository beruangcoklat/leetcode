class Solution {
    fun convertDateToBinary(date: String): String {
        val split = date.split("-")
        var ans = mutableListOf<String>()
        split.forEach { s -> ans.add(decimalToBinary(s)) }
        return ans.joinToString("-")
    }

    private fun decimalToBinary(decimal: String): String {
        var decimalInt = decimal.toInt()
        var ans = ""
        while (decimalInt > 0) {
            if (decimalInt % 2 == 1) {
                ans = "1" + ans
            } else {
                ans = "0" + ans
            }
            decimalInt = decimalInt shr 1
        }
        return ans.trimStart('0')
    }
}