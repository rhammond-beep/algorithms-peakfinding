package peakfinder

// Accept a 1D array of integers, retun the first peak otherwise return an error
type PeakFinder interface {
	FindPeak(numbers []int) (int, error)
}

// Less Efficient Peak Finder
type LinearPeakFinder struct {
	numbers []int
}

// More Efficient Peak Finder
type NLogNPeakFinder struct {
	numbers []int
}

func (p *LinearPeakFinder) FindPeak(numbers []int) (int, error) {
	var peak int
	input_size := len(numbers)

	if input_size == 1 {
		return numbers[0], nil
	}

	if numbers[0] >= numbers[1] {
		return numbers[0], nil
	}

	if numbers[len(numbers)-1] >= numbers[len(numbers)-2] {
		return numbers[0], nil
	}

	for i := 1; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] && numbers[i] >= numbers[i-1] {
			peak = numbers[i]
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
