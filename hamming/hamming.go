package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("Strings are of different lengths!")
    } else {
    	diff := 0
    	for i:=0; i < len(a); i++ {
            if a[i] != b[i]{
                diff += 1
            }
        }
    	return diff, nil
    }
	return 0, nil
}
