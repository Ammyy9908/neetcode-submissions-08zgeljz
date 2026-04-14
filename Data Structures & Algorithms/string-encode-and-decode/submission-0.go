type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var encodedString string
	for _,v := range strs{
		encodedString += fmt.Sprintf("%d#%s", len(v), v)
	}

return encodedString


}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0

	for i < len(encoded) {
		// find delimiter '#'
		j := i
		for encoded[j] != '#' {
			j++
		}

		// length of string
		length, _ := strconv.Atoi(encoded[i:j])

		// extract string
		start := j + 1
		end := start + length
		res = append(res, encoded[start:end])

		// move pointer
		i = end
	}

	return res
}
