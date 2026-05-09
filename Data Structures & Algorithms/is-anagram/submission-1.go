func isAnagram(s string, t string) bool {
    firstword := make(map[string]int)
    secondword := make(map[string]int)

    for _,letter := range s {
        firstword[string(letter)]++
    }

    for _, letter := range t {
        secondword[string(letter)]++
    }

    if len(firstword) != len(secondword) {
        return false
    }

    for char, count := range firstword {
        if secondword[char] != count {
            return false
        }
    }
    return true
}
