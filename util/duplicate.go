package util

func RemoveDuplicate(in []string) []string {
	m := map[string]bool{}
	res := []string{}

	for _, i := range in {
		if _, ok := m[i]; !ok {
			m[i] = true
			res = append(res, i)
		}
	}

	return res
}
