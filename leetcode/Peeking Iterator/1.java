// Java Iterator interface reference:
// https://docs.oracle.com/javase/8/docs/api/java/util/Iterator.html

class PeekingIterator implements Iterator<Integer> {
    
    List<Integer> arr;
        int idx;

        public PeekingIterator(Iterator<Integer> iterator) {
            // initialize any member here.
            arr = new ArrayList();
            while(iterator.hasNext()){
                arr.add(iterator.next());
            }
        }

        // Returns the next element in the iteration without advancing the iterator.
        public Integer peek() {
            return arr.get(idx);
        }

        // hasNext() and next() should behave the same as in the Iterator interface.
        // Override them if needed.
        @Override
        public Integer next() {
            return arr.get(idx++);
        }

        @Override
        public boolean hasNext() {
            return idx < arr.size();
        }
}
