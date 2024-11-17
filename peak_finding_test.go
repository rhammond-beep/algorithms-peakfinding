package peakfinder

import "testing"

func TestLogNPeakFinder(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 5}
	finder := &LogNPeakFinder{numbers: input}
	want := 6
	peak, err := finder.FindPeak(input)
	if err != nil {
		t.Fatal(err.Error())
	} else if peak != want {
		t.Fatalf("Failed to find peak correctly, wanted: %v, got: %v", want, peak)
	}
}

// func TestNLogNPeakFinder(t *testing.T) {
// 	input := []int{1, 2, 3, 4, 6, 5}
// 	finder := &NLogNPeakFinder{numbers: input}
// 	want := 6
// 	peak, err := finder.FindPeak(input)
// 	if err != nil {
// 		t.Fatal(err.Error())
// 	} else if peak != want {
// 		t.Fatalf("Failed to find peak correctly, wanted: %q, got: %q", want, peak)
// 	}
// }
