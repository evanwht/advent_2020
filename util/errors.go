package util

// Check panics if the err is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
