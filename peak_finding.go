package peakfinder

// Accept a 1D array of integers, retun the first peak otherwise return an error
type PeakFinder interface {
	FindPeak(numbers []int) (int, error)
}

// Less Efficient Peak Finder
type LogNPeakFinder struct {
	numbers []int
}

// More Efficient Peak Finder
type NLogNPeakFinder struct {
	numbers []int
}

func (p *LogNPeakFinder) FindPeak(numbers []int) (int, error) {
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

func (p *NLogNPeakFinder) FindPeak(numbers []int) (int, error) {
	var peak int

	if len(numbers) == 1 {
		return numbers[0], nil
	}

	return peak, nil
}
