class Solution {
    fun halvesAreAlike(s: String): Boolean {
        val a = s.subSequence(0, s.length / 2).filter { isVowel(it) }.toList().size
        val b = s.subSequence(s.length / 2, s.length).filter { isVowel(it) }.toList().size
        return a == b
    }

    private fun isVowel(c: Char): Boolean =
        when (c) {
            'a' -> true
            'e' -> true
            'i' -> true
            'o' -> true
            'u' -> true
            'A' -> true
            'E' -> true
            'I' -> true
            'O' -> true
            'U' -> true
            else -> false
        }
}