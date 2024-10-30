class Solution {
    fun moveZeroes(nums: IntArray): Unit {
        var zeroPointer = 0
        var nonZeroPointer = 0

        while (zeroPointer < nums.size && nonZeroPointer < nums.size) {
            while (zeroPointer < nums.size) {
                if (nums[zeroPointer] == 0) {
                    break
                }
                zeroPointer++
            }
            while (nonZeroPointer < nums.size) {
                if (nums[nonZeroPointer] != 0) {
                    break
                }
                nonZeroPointer++
            }

            if (zeroPointer > nonZeroPointer) {
                nonZeroPointer++
                continue
            }

            if (zeroPointer < nums.size && nonZeroPointer < nums.size &&
                nums[zeroPointer] == 0 && nums[nonZeroPointer] != 0
            ) {
                val temp = nums[zeroPointer]
                nums[zeroPointer] = nums[nonZeroPointer]
                nums[nonZeroPointer] = temp
            }
        }
    }
}