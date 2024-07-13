class Solution {
	class Node {
		private int value;
		private Node next, prev;

		public Node(int value) {
			this.value = value;
			next = null;
			prev = null;
		}
	}

	private Node head, tail;

	private void insert(int value) {
		Node newNode = new Node(value);
		if (head == null) {
			head = newNode;
			tail = newNode;
			return;
		}
		tail.next = newNode;
		newNode.prev = tail;
		tail = newNode;
	}

	public int findTheWinner(int n, int k) {
		head = null;
		for (int i = 1; i <= n; i++) {
			insert(i);
		}
		tail.next = head;
		head.prev = tail;

		while (head != null) {
			for (int i = 0; i < k - 1; i++) {
				head = head.next;
			}
			if (head.next == head.prev) {
				return head.next.value;
			}
			Node nextHead = head.next;
			head.prev.next = head.next;
			head.next.prev = head.prev;
			head = nextHead;
		}
		return 0;
	}
}