impl Solution {
    pub fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
        let mut result: Vec<i32> = Vec::new();
        for i in 0..nums.len() {
            result.push(*nums.get(i).unwrap());
        }
        for i in 0..nums.len() {
            result.push(*nums.get(i).unwrap());
        }
        result
    }
}
