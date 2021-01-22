package main

func main() {

}

func findRepeatedDnaSequences(s string) []string {
	repeatTimes := make(map[string]int)
	res := []string{}
	for i := 0; i+10 <= len(s); i++ {
		repeatTimes[s[i:i+10]]++
	}
	for s, t := range repeatTimes {
		if t > 1 {
			res = append(res, s)
		}
	}
	return res
}
