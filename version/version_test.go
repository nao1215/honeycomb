package version

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetVersion(t *testing.T) { //nolint:paralleltest
	t.Run("If global variable 'Version' is empty, return '' from build information", func(t *testing.T) { //nolint:paralleltest
		got := GetVersion()
		want := ""
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("GetVersion() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("If global variable 'Version' is not empty, return 'Version' from build information", func(t *testing.T) { //nolint:paralleltest
		Version = "test"
		got := GetVersion()
		want := "test"
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("GetVersion() mismatch (-want +got):\n%s", diff)
		}
	})
}
