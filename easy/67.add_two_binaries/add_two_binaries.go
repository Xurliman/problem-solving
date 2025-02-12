package add_two_binaries

import "log"

func addBinary(a string, b string) string {
	al, bl := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for al >= 0 || bl >= 0 || carry > 0 {
		sum := carry
		if al >= 0 {
			sum += int(a[al] - '0')
			al--
		}
		if bl >= 0 {
			sum += int(b[bl] - '0')
			bl--
		}
		carry = sum / 2
		log.Println("sum", sum)
		result = string(sum%2+'0') + result
	}

	return result
}
