package uitls

func IsSet(d []string, key int) bool {
	for k, _ := range d {
		if k == key {
			return true
		}
	}
	return false
}
