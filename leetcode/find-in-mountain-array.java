/**
 * // This is MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * interface MountainArray {
 *     public int get(int index) {}
 *     public int length() {}
 * }
 */
 
class Solution {
    HashMap<Integer, Integer> memo = new HashMap<>();

    int get(MountainArray mountainArr, int idx) {
        if (memo.containsKey(idx)) {
            return memo.get(idx);
        }

        int x = mountainArr.get(idx);
        memo.put(idx, x);
        return x;
    }

    int binarySearch(int left, int right, MountainArray mountainArr) {
        if (left == right) {
            return -1;
        }

        int mid = (left + right) / 2;

        int midVal = get(mountainArr, mid);
        int midNextVal = get(mountainArr, mid + 1);
        int midPrevVal = get(mountainArr, mid - 1);

        if (midPrevVal < midVal && midVal > midNextVal) {
            return mid;
        }

        if (midPrevVal < midVal && midVal < midNextVal) {
            return binarySearch(mid + 1, right, mountainArr);
        }

        if (midPrevVal > midVal && midVal > midNextVal) {
            return binarySearch(left, mid, mountainArr);
        }

        return -2;
    }

    int binarySearch2(int left, int right, int target, MountainArray mountainArr) {
        if (left == right) {
            if (get(mountainArr, left) == target) {
                return left;
            }
            return -1;
        }

        int mid = (left + right) / 2;
        int midVal = get(mountainArr, mid);

        if (midVal == target) {
            return mid;
        }

        if (target < midVal) {
            return binarySearch2(left, mid, target, mountainArr);
        }

        if (target > midVal) {
            return binarySearch2(mid + 1, right, target, mountainArr);
        }

        return -2;
    }

    int binarySearch3(int left, int right, int target, MountainArray mountainArr) {
        if (left == right) {
            if (get(mountainArr, left) == target) {
                return left;
            }
            return -1;
        }

        int mid = (left + right) / 2;
        int midVal = get(mountainArr, mid);

        if (midVal == target) {
            return mid;
        }

        if (target > midVal) {
            return binarySearch3(mid + 1, right, target, mountainArr);
        }

        if (target < midVal) {
            return binarySearch3(mid + 1, right, target, mountainArr);
        }

        return -2;
    }

    public int findInMountainArray(int target, MountainArray mountainArr) {
        memo = new HashMap<>();
        int len = mountainArr.length();

        int pivot = binarySearch(0, len - 1, mountainArr);
        if (pivot < 0) {
            return -1;
        }

        int res = binarySearch2(0, pivot, target, mountainArr);
        if (res >= 0) {
            return res;
        }

        res = binarySearch3(pivot, len - 1, target, mountainArr);
        if (res >= 0) {
            return res;
        }

        return -1;
    }
}
