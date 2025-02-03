#[allow(dead_code)]
fn valid_anagram(string1: &str, string2: &str) -> bool {
    let mut arr: [u8; 26] = [0; 26]; // initialize an empty arr

    // increment the value for first array
    for i in string1.to_ascii_lowercase().chars() {
        //arr[(i as u8 % 26) as usize] += 1; // you won't know the location of the word in the array
        // as 'a' is at 95
        arr[(i as u8 - b'a') as usize] += 1;
    }

    // decrement the value for the second array while check if it doesn't fall below 0
    for i in string2.to_ascii_lowercase().chars() {
        if arr[(i as u8 - b'a') as usize] <= 0 {
            return false;
        }
        arr[(i as u8 - b'a') as usize] -= 1;
    }

    // check if there are any values remaining
    for &i in arr.iter() {
        if i != 0 {
            return false;
        }
    }

    return true;
}

// Optimized soln
#[allow(dead_code)]
fn valid_anagram_2(string1: &str, string2: &str) -> bool {
    if string1.len() != string2.len() {
        return false;
    }
    let mut arr: [i16; 26] = [0; 26];
    for i in 0..string1.len() {
        arr[(string1.as_bytes()[i] - b'a') as usize] += 1;
        arr[(string2.as_bytes()[i] - b'a') as usize] -= 1;
    }
    for &i in arr.iter() {
        if i != 0 {
            return false;
        }
    }
    return true;
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn valid_anagram_test() {
        assert_eq!(valid_anagram("anagram", "nagaram"), true);
        assert_eq!(valid_anagram("rat", "car"), false);
        assert_eq!(valid_anagram("hello", "llloh"), false);
        assert_eq!(valid_anagram("hello", "llohe"), true);
        assert_eq!(valid_anagram("hello", "lolhe"), true);
        assert_eq!(
            valid_anagram(
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
            ),
            false
        );
    }

    #[test]
    fn valid_anagram_2_test() {
        assert_eq!(valid_anagram_2("anagram", "nagaram"), true);
        assert_eq!(valid_anagram_2("rat", "car"), false);
        assert_eq!(valid_anagram_2("hello", "llloh"), false);
        assert_eq!(valid_anagram_2("hello", "llohe"), true);
        assert_eq!(valid_anagram_2("hello", "lolhe"), true);
        assert_eq!(
            valid_anagram_2(
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
            ),
            false
        );
    }
}
