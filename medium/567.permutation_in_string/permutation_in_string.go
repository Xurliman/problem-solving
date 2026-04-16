package main

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	var s1Arr, s2Arr [26]int
	for i := 0; i < n1; i++ {
		s1Arr[s1[i]-'a']++
		s2Arr[s2[i]-'a']++
	}

	for i := 0; i < n2-n1; i++ {
		if s1Arr == s2Arr {
			return true
		}

		s2Arr[s2[i]-'a']--
		s2Arr[s2[i+n1]-'a']++
	}
	return s1Arr == s2Arr
}
