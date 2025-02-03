use std::collections::HashSet;

#[allow(dead_code)]
fn contains_duplicate(arr: &[isize]) -> bool {
    // put all the element in the hashset
    let hash_set: HashSet<isize> = arr.iter().copied().collect();
    // return the comparison of the length of the hashset
    arr.len() != hash_set.len() // if the length is not equal, then there is a duplicate
}

// Optimized soln
#[allow(dead_code)]
fn contains_duplicate_2(arr: &[isize]) -> bool {
    let mut hash_set = HashSet::new();
    for i in arr.iter() {
        if !hash_set.insert(i) {
            return true; // matched hence return true
        }
    }
    false // no match hence return false
}

#[cfg(test)]
mod tests {
    use super::*; // Import everything from the parent module

    #[test]
    fn test_contains_duplicate() {
        assert_eq!(contains_duplicate(&[1, 2, 3, 4, 5]), false); // No duplicates
        assert_eq!(contains_duplicate(&[1, 2, 3, 3, 5]), true); // Has a duplicate (3)
        assert_eq!(contains_duplicate(&[7, 7, 7, 7, 7]), true); // All elements are the same
        assert_eq!(contains_duplicate(&[]), false); // Empty array (no duplicates)
        assert_eq!(contains_duplicate(&[100, 200, 300, 400, 500]), false); // Unique numbers
    }

    #[test]
    fn test_contains_duplicate_2() {
        assert_eq!(contains_duplicate_2(&[1, 2, 3, 4, 5]), false); // No duplicates
        assert_eq!(contains_duplicate_2(&[1, 2, 3, 3, 5]), true); // Has a duplicate (3)
        assert_eq!(contains_duplicate_2(&[7, 7, 7, 7, 7]), true); // All elements are the same
        assert_eq!(contains_duplicate_2(&[]), false); // Empty array (no duplicates)
        assert_eq!(contains_duplicate_2(&[100, 200, 300, 400, 500]), false); // Unique numbers
    }
}
