impl Solution {
    pub fn most_words_found(sentences: Vec<String>) -> i32 {
        let mut max = 0;
        for sentence in sentences {
            let mut words = sentence.split_whitespace();
            let mut count = 0;
            loop {
                match words.next() {
                    None => break,
                    Some(_) => count += 1,
                }
            }
            if count > max {
                max = count;
            }
        }
        max
    }
}
