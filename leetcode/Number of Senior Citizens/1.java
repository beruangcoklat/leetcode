class Solution {
    public int countSeniors(String[] details) {
        long count = Arrays.stream(details).filter(s -> {
            int age = ((s.charAt(11) - '0') * 10) + (s.charAt(12) - '0');
            return age > 60;
        }).count();
        return (int) count;
    }
}