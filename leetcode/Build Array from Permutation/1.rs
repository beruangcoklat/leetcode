impl Solution {
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        let mut result: Vec<i32> = Vec::new();
        for i in 0..nums.len() {
            let x = *nums.get(i).unwrap() as usize;
            let x = *nums.get(x).unwrap();
            result.push(x);
        }
        result
    }
}
