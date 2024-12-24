package peakfinder

import "testing"

func TestNLogNPeakFinder(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 5}
	finder := &NLogNPeakFinder{}
	want := 6
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestNLogNPeakFinderSingleValue(t *testing.T) {
	input := []int{1}
	finder := &NLogNPeakFinder{}
	want := 1
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestNLogNPeakFinderAllValuesAreTheSame(t *testing.T) {
	input := []int{0, 0, 0, 0, 0}
	finder := &NLogNPeakFinder{}
	want := 0
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

func TestLogNEdgePeaksAreCorrectlyHandled(t *testing.T) {
	input := []int{1, 0}
	finder := &NLogNPeakFinder{}
	want := 1
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}
