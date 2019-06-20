package FdfExtract

import "testing"

// TestCheckpath tests checkpath function in checkpath.go
func TestCheckpath(t *testing.T) {
	testcases := []struct {
		testname string
		testpath string
		out      bool
	}{
		{"Simple Test .fdf", "test.fdf", true},
		{"Simple Test no .fdf", "test.pdf", false},
		{"Test fdf within file path", "testfdf.txt", false},
	}
	for i := range testcases {
		if (checkpath(testcases[i].testpath) == nil) != testcases[i].out {
			t.Errorf("%s FAILED", testcases[i].testname)
		}
	}
}
