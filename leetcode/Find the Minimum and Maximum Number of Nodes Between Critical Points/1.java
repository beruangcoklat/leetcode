/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
	public int[] nodesBetweenCriticalPoints(ListNode head) {
		ListNode curr = head;
		ListNode prev = null;
		List<Integer> criticalPoints = new ArrayList<>();

		int idx = -1;
		while (curr != null) {
			idx++;
			if (prev != null && curr.next != null) {
				boolean localMaxima = curr.val > prev.val && curr.val > curr.next.val;
				boolean localMinima = curr.val < prev.val && curr.val < curr.next.val;
				if (localMaxima || localMinima) {
					criticalPoints.add(idx);
				}
			}
			prev = curr;
			curr = curr.next;
		}

		int minDiff = 0;
		for (int i = 0; i < criticalPoints.size() - 1; i++) {
			int diff = criticalPoints.get(i + 1) - criticalPoints.get(i);
			if (minDiff == 0 || diff < minDiff) {
				minDiff = diff;
			}
		}
		if (criticalPoints.size() <= 1) {
			return new int[]{-1, -1};
		}
		return new int[]{minDiff, criticalPoints.get(criticalPoints.size() - 1) - criticalPoints.get(0)};
	}
}
