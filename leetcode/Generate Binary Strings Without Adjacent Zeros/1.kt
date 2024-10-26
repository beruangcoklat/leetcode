class Solution {
    private var ans: MutableList<String> = mutableListOf()

    fun validStrings(n: Int): List<String> {
        ans = mutableListOf()
        generateAns(n)
        return ans
    }

    private fun generateAns(n: Int, str: String = "") {
        if (str.length == n) {
            ans.add(str)
            return
        }

        generateAns(n, str + "1")
        if (str.isEmpty() || str[str.length - 1] != '0') {
            generateAns(n, str + "0")
        }
    }
}