impl Solution {
    pub fn minimum_sum(num: i32) -> i32 {
        let mut digits: Vec<i32> = Vec::new();
        let mut divisor = 1;
        while divisor <= num {
            let digit = num / divisor % 10;
            digits.push(digit);
            divisor *= 10;
        }
        digits.sort();

        let a = digits.get(0).unwrap() * 10 + digits.get(2).unwrap();
        let b = digits.get(1).unwrap() * 10 + digits.get(3).unwrap();
        a + b
    }
}
