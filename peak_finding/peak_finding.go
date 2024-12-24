package peakfinder

// Accept a 1D array of integers, retun the first peak otherwise return an error
type PeakFinder interface {
	FindPeak() (int, error)
}

// Less Efficient Peak Finder
type LinearPeakFinder struct {
	Numbers []int
}

// More Efficient Peak Finder
type NLogNPeakFinder struct {
}

func (p *LinearPeakFinder) FindPeak() (int, error) {
	var peak int
	input_size := len(p.Numbers)

	if input_size == 1 {
		return p.Numbers[0], nil
	}

	if p.Numbers[0] >= p.Numbers[1] {
		return p.Numbers[0], nil
	}

	if p.Numbers[len(p.Numbers)-1] >= p.Numbers[len(p.Numbers)-2] {
		return p.Numbers[0], nil
	}

	for i := 1; i < len(p.Numbers)-1; i++ {
		if p.Numbers[i] >= p.Numbers[i+1] && p.Numbers[i] >= p.Numbers[i-1] {
			peak = p.Numbers[i]
			break
		}
	}

	return peak, nil
}

/*
* The more efficient recursive implementation
 */
func (p *NLogNPeakFinder) FindPeak(numbers []int) (int, error) {
	n := len(numbers)
	if n == 1 { // Handle the successful base case
		return numbers[0], nil
	}
	if numbers[n/2] <= numbers[(n/2)-1] {
		return p.FindPeak(numbers[:(n / 2)])
	} else if numbers[n/2] <= numbers[(n/2)+1] {
		return p.FindPeak(numbers[(n / 2):])
	} else {
		return numbers[n/2], nil
	}
}
