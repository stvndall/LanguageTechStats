package factors

// PrimeFactors will take a number, and return a set of prime factors
func PrimeFactors(number int) (pfs []int) {
	for number%2 == 0 {
		pfs = append(pfs, 2)
		number = number / 2
	}

	for i := 3; i*i <= number; i = i + 2 {
		for number%i == 0 {
			pfs = append(pfs, i)
			number = number / i
		}
	}
	if number > 2 {
		pfs = append(pfs, number)
	}
	return
}
