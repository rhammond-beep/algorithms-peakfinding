package peakfinder

// Accept a 1D array of integers, retun the first peak otherwise return an error
type PeakFinder interface {
	FindPeak(numbers []int) (int, error)
}

type LogNPeakFinder struct {
	numbers []int
}

type NlogNPeakFinder struct {
	numbers []int
}

// TODO - implement
func (p *LogNPeakFinder) FindPeak(numbers []int) (int, error) {

}

// TODO - implement
func (p *NlogNPeakFinder) FindPeak(numbers []int) (int, error) {
}

func FindPeak() {
}
