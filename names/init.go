package names

import (
	"strconv"
	"strings"
)

var male map[string]float32
var female map[string]float32
var surname map[string]float32

func init() {
	male = toMap(Male)
	female = toMap(Female)
	surname = toMap(Surname)
}

func toMap(data string) map[string]float32 {
	results := map[string]float32{}

	for _, line := range strings.Split(data, "\n") {
		v := strings.TrimSpace(line)
		if v != "" {
			parts := strings.Split(v, ",")
			name := parts[0]
			occurance, _ := strconv.ParseFloat(parts[1], 32)

			results[name] = float32(occurance)
		}
	}

	return results
}

func IsMale(s string) (float32, bool) {
	s = normalize(s)
	occurance, found := male[s]
	return occurance, found
}

func IsFemale(s string) (float32, bool) {
	s = normalize(s)
	occurance, found := female[s]
	return occurance, found
}

func IsSurname(s string) (float32, bool) {
	s = normalize(s)
	occurance, found := surname[s]
	return occurance, found
}

func IsFirstName(s string) bool {
	if _, found := IsFemale(s); found {
		return true
	}

	if _, found := IsMale(s); found {
		return true
	}

	return false
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = strings.Replace(s, "'", "", -1)
	return s
}
