use std::collections::HashMap;

#[allow(dead_code)]
fn is_anagram(string1: &str, string2: &str) -> bool {
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

// TODO: DO IT LATER IN CLG
// use hashmap;
#[allow(dead_code)]
fn group_anagrams(strings: &[String]) -> Vec<Vec<String>> {
    let mut result: Vec<Vec<String>> = vec![];
    for i in 0..strings.len() {
        // string is empty since in condition it is given that no string will be empty
        if strings[i].len() == 0 {
            continue;
        }

        for j in i + 1..strings.len() {
            if strings[j].len() == 0 {
                continue;
            }

            if is_anagram(&strings[i], &strings[j]) {
                result.push(vec![strings[i].clone(), strings[j].clone()]);
                break;
            }
        }
    }
    return vec![];
}

#[allow(dead_code)]
fn group_anagrams_2(strings: &[String]) -> Vec<Vec<String>> {
    let mut map: HashMap<[u8; 26], Vec<String>> = HashMap::new();

    for string in strings {
        // get the hash for the current string
        let mut key: [u8; 26] = [0; 26];
        for c in string.chars() {
            key[(c as u8 - b'a') as usize] += 1;
        }

        // check if the hash already exists in the map
        if let Some(v) = map.get_mut(&key) {
            v.push(string.clone()); // push the key if the value is already present
        } else {
            // if not, create a new vector with the current string
            map.insert(key, vec![string.clone()]);
        }
    }

    map.into_values().collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_group_anagrams_empty() {
        let input: Vec<String> = Vec::new();
        let result = group_anagrams_2(&input);
        assert_eq!(result, Vec::<Vec<String>>::new());
    }

    #[test]
    fn test_group_anagrams_single_word() {
        let input = vec!["eat".to_string(), "tea".to_string(), "tan".to_string()];
        let result = group_anagrams_2(&input);
        let expected = vec![
            vec!["eat".to_string(), "tea".to_string()],
            vec!["tan".to_string()],
        ];

        // Sort both the result and expected values before comparison
        let mut sorted_result = result.clone();
        for group in &mut sorted_result {
            group.sort();
        }
        sorted_result.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        let mut sorted_expected = expected.clone();
        for group in &mut sorted_expected {
            group.sort();
        }
        sorted_expected.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        assert_eq!(sorted_result, sorted_expected);
    }

    #[test]
    fn test_group_anagrams_all_same() {
        let input = vec![
            "listen".to_string(),
            "silent".to_string(),
            "enlist".to_string(),
        ];
        let result = group_anagrams_2(&input);
        let expected = vec![vec![
            "listen".to_string(),
            "silent".to_string(),
            "enlist".to_string(),
        ]];

        // Sort both the result and expected values before comparison
        let mut sorted_result = result.clone();
        for group in &mut sorted_result {
            group.sort();
        }
        sorted_result.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        let mut sorted_expected = expected.clone();
        for group in &mut sorted_expected {
            group.sort();
        }
        sorted_expected.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        assert_eq!(sorted_result, sorted_expected);
    }

    #[test]
    fn test_group_anagrams_no_anagrams() {
        let input = vec![
            "apple".to_string(),
            "banana".to_string(),
            "cherry".to_string(),
        ];
        let result = group_anagrams_2(&input);
        println!("{:?}", result);

        let expected = vec![
            vec!["apple".to_string()],
            vec!["banana".to_string()],
            vec!["cherry".to_string()],
        ];

        // Sort both the result and expected values before comparison
        let mut sorted_result = result.clone();
        for group in &mut sorted_result {
            group.sort();
        }
        sorted_result.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        let mut sorted_expected = expected.clone();
        for group in &mut sorted_expected {
            group.sort();
        }
        sorted_expected.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        assert_eq!(sorted_result, sorted_expected);
    }

    #[test]
    fn test_group_anagrams_mixed() {
        let input = vec![
            "bat".to_string(),
            "tab".to_string(),
            "cat".to_string(),
            "act".to_string(),
            "tac".to_string(),
        ];
        let result = group_anagrams_2(&input);
        let expected = vec![
            vec!["bat".to_string(), "tab".to_string()],
            vec!["cat".to_string(), "act".to_string(), "tac".to_string()],
        ];

        // Sort both the result and expected values before comparison
        let mut sorted_result = result.clone();
        for group in &mut sorted_result {
            group.sort();
        }
        sorted_result.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        let mut sorted_expected = expected.clone();
        for group in &mut sorted_expected {
            group.sort();
        }
        sorted_expected.sort_by(|a, b| a[0].cmp(&b[0])); // Sort by first element for consistent order

        assert_eq!(sorted_result, sorted_expected);
    }
}
