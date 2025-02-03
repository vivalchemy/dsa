// Optimized soln for non sorted array
#[allow(dead_code)]
fn two_sum(nums: &[i32], target: i32) -> Vec<i32> {
    let mut map = std::collections::HashMap::new();

    for i in 0..nums.len() {
        let complement = target - nums[i];

        if map.contains_key(&complement) {
            return vec![map[&complement], i as i32];
        }

        map.insert(nums[i], i as i32);
    }

    vec![]
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn two_sum_test() {
        assert_eq!(two_sum(&[2, 7, 11, 15], 9), vec![0, 1]);
        assert_eq!(two_sum(&[3, 2, 4], 6), vec![1, 2]);
        assert_eq!(two_sum(&[3, 3], 6), vec![0, 1]);
        assert_eq!(two_sum(&[3, 3, 4, 4], 8), vec![2, 3]);
        assert_eq!(two_sum(&[4, 6, 6, 8, 8, 10], 24), vec![]);
        assert_eq!(two_sum(&[1, 2, 3, 4, 5, 6, 7, 8, 9], 17), vec![7, 8]);
    }
}
