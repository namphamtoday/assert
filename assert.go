package assert

import "testing"

func AssertEqual[T comparable](t *testing.T, expect T, output T) {
	t.Helper()
	if expect != output {
		t.Errorf("========> FAILED: Expect %v - Ouput %v", expect, output)
	} else {
		t.Logf("PASSED: [%v]", output)
	}
}
