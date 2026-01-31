package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	for i := 0; ; {
		if n == 1 {
			return i, nil
		} else if n <= 0 {
			fmt.Println("found 0")
			return 0, errors.New("0 is not valid")
		} else if n%2 == 0 {
			n = n / 2
			i++
		} else if n%2 == 1 {
			n = (n * 3) + 1
			i++
		}
	}
}
