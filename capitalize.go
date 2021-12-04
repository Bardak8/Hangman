package main

func IsLetter2(r rune) bool {
	lettermaj := ('A' <= r) && (r <= 'Z')
	lettermin := ('a' <= r) && (r <= 'z')
	letternum := ('0' <= r) && (r <= '9')
	if lettermaj || lettermin || letternum {
		return true
	}
	return false
}

func IsMin(r rune) bool {
	lettermin := ('a' <= r) && (r <= 'z')
	return lettermin
}

func IsMaj(r rune) bool {
	lettermaj := ('A' <= r) && (r <= 'Z')
	return lettermaj
}

func Capitalize(s string) string {
	str2 := []rune(s)
	for index := range str2 {
		if index == 0 {
			if IsMin(str2[index]) {
				str2[index] -= 32
			}
		} else if !IsLetter2(str2[index-1]) && IsMin(str2[index]) {
			str2[index] -= 32
		} else if IsLetter2(str2[index-1]) && IsMaj(str2[index]) {
			str2[index] += 32
		}
	}
	return string(str2)
}
func Contains(s string, char rune) bool { // Si une string est contenue dans un tableau
	for _, a := range s {
		if a == char {
			return true
		}
	}
	return false
}
