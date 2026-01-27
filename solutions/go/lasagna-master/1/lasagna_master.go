package lasagna

func PreparationTime(l []string, t int)(estimate int) {
	if t == 0 {
		t = 2
	}
	estimate = len(l) * t
	return
}

func Quantities(l []string)(noodles int, sauce float64) {
	for _, v := range l {
		switch v {
			case "noodles":
				noodles += 50
			case "sauce":
				sauce += 0.2
		}	
	}
	return
}

func AddSecretIngredient(f []string, m []string) {
	m[len(m)-1] = f[len(f)-1]
}

func ScaleRecipe(q []float64, p int)([]float64) {
	var n []float64
	for _,v := range q {
		n = append(n, (float64(p) * v/2))
	}
	return n
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
