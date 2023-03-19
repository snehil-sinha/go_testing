package funcs_test

import "testing"

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleanup")
	})
	t.Log("Running some test")
}
