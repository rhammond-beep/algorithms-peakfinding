package peakfinder

import "testing"

func TestLogNPeakFinderSmallCase(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 5}
	finder := &LinearPeakFinder{Numbers: input}
	want := 6
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestLogNPeakFinderSingleValue(t *testing.T) {
	input := []int{1}
	finder := &LinearPeakFinder{Numbers: input}
	want := 1
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestLogNPeakFinderAllValuesAreTheSame(t *testing.T) {
	input := []int{0, 0, 0, 0, 0}
	finder := &LinearPeakFinder{Numbers: input}
	want := 0
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestEdgePeaksAreCorrectlyHandled(t *testing.T) {
	input := []int{1, 0}
	finder := &LinearPeakFinder{Numbers: input}
	want := 1
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}
