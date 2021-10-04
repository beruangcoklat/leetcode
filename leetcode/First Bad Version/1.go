/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var cache map[int]bool

func isBadVersionCache(version int) bool {
	if data, ok := cache[version]; ok {
		return data
	}
	result := isBadVersion(version)
	cache[version] = result
	return result
}

func firstBadVersion(n int) int {
	cache = make(map[int]bool)
	left, right := 1, n

	for left <= right {
		mid := (left + right) / 2

		if !isBadVersionCache(mid) {
			left = mid + 1
			continue
		}

		if !isBadVersion(mid - 1) {
			return mid
		}

		right = mid - 1
	}

	return -1
}

