class Solution {
    public String countOfAtoms(String formula) {
        List<String> tokens = new ArrayList<>();
        for (int i = 0; i < formula.length(); i++) {
            char curr = formula.charAt(i);
            String currStr = curr + "";

            if (curr == ')') {
                int idx = i + 1;
                int number = 0;
                while (idx < formula.length()) {
                    char next = formula.charAt(idx);
                    if (Character.isDigit(next)) {
                        number = (number * 10) + (next - '0');
                    } else {
                        break;
                    }
                    idx++;
                    i = idx - 1;
                }

                number = Math.max(number, 1);

                idx = tokens.size() - 1;
                while (idx >= 0) {
                    String token = tokens.get(idx);
                    if (token.equals("(")) {
                        tokens.remove(idx);
                        break;
                    }
                    tokens.set(idx, multiplyAtom(token, number));
                    idx--;
                }
                continue;
            }

            if (Character.isUpperCase(curr)) {
                int idx = i + 1;
                while (idx < formula.length()) {
                    char next = formula.charAt(idx);
                    if (Character.isUpperCase(next) || next == '(' || next == ')') {
                        break;
                    }
                    if (Character.isLowerCase(next)) {
                        currStr += next;
                    } else if (Character.isDigit(next)) {
                        currStr += next;
                    }
                    idx++;
                    i = idx - 1;
                }
            }
            tokens.add(currStr);
        }

        Map<String, Integer> atomMap = new HashMap<>();
        for (String token : tokens) {
            Atom atom = getAtom(token);
            int count = atomMap.getOrDefault(atom.atom, 0);
            atomMap.put(atom.atom, count + atom.count);
        }

        List<Atom> atoms = new ArrayList<>();
        atomMap.forEach((atom, count) -> atoms.add(new Atom(atom, count)));

        Collections.sort(atoms, (a, b) -> {
            return a.atom.compareTo(b.atom);
        });

        String result = "";
        for (int i = 0; i < atoms.size(); i++) {
            result += atoms.get(i).String();
        }
        return result;
    }

    String multiplyAtom(String str, int multiplier) {
        Atom atom = getAtom(str);
        return atom.atom + (atom.count * multiplier);
    }

    Atom getAtom(String str) {
        String atom = "";
        int count = 0;
        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            if (Character.isLetter(c)) {
                atom += c;
            } else {
                count = (count * 10) + (c - '0');
            }
        }
        return new Atom(atom, Math.max(count , 1));
    }

    class Atom {
        int count;
        String atom;

        Atom(String atom, int count) {
            this.atom = atom;
            this.count = count;
        }

        String String() {
            if (count == 1) {
                return atom;
            }
            return atom + count;
        }
    }
}