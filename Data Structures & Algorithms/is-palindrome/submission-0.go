

func isPalindrome(s string) bool {
    cleaned := []rune{}

    for _, c := range s {
        if unicode.IsLetter(c) || unicode.IsDigit(c) {
            cleaned = append(cleaned, unicode.ToLower(c))
        }
    }

    i, j := 0, len(cleaned)-1
    for i < j {
        if cleaned[i] != cleaned[j] {
            return false
        }
        i++
        j--
    }

    return true
}
