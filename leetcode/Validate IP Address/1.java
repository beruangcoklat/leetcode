class Solution {
    
    boolean validIpV4(String[] splits) {
        for (String s : splits) {
            int len = s.length();
            if (len < 1 || len > 3) {
                return false;
            }
            if (len == 1 &&  s.charAt(0)== '0') {
                continue;
            }
            int total = 0;
            for (char c : s.toCharArray()) {
                if (!Character.isDigit(c)) {
                    return false;
                }
                if (total == 0 && c == '0') {
                    return false;
                }
                total *= 10;
                total += (c - '0');
            }
            if (total > 255) {
                return false;
            }
        }
        return true;
    }

    boolean validIpV6(String[] splits) {
        for (String s : splits) {
            int len = s.length();
            if (len > 4 || len < 1) {
                return false;
            }
            for (char c : s.toCharArray()) {
                if (Character.isDigit(c)) continue;
                if (c >= 'a' && c <= 'f') continue;
                if (c >= 'A' && c <= 'F') continue;
                return false;
            }
        }
        return true;
    }
    
    public String validIPAddress(String IP) {
        if (IP.endsWith(".")) {
            return "Neither";
        }
        String splits[] = IP.split("[.]");
        if (splits.length == 4 && validIpV4(splits)) {
            return "IPv4";
        }

        if (IP.endsWith(":")) {
            return "Neither";
        }
        splits = IP.split("[:]");
        if (splits.length == 8 && validIpV6(splits)) {
            return "IPv6";
        }

        return "Neither";
    }
}