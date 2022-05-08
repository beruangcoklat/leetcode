impl Solution {
    pub fn final_value_after_operations(operations: Vec<String>) -> i32 {
        let mut x = 0;
        for ops in operations {
            if ops.contains("++") {
                x += 1;
            } else {
                x -= 1;
            }
        }
        x
    }
}
