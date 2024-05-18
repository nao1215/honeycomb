// Package main is the entry point of the honeycomb application.
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_main(t *testing.T) { //nolint:paralleltest
	t.Run("show version", func(t *testing.T) { //nolint:paralleltest
		orgStdout := os.Stdout
		orgStderr := os.Stderr
		pr, pw, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		os.Stdout = pw
		os.Stderr = pw

		defer func() {
			os.Stdout = orgStdout
			os.Stderr = orgStderr
			pw.Close() //nolint
			pr.Close() //nolint
		}()

		done := make(chan struct{})
		go func() {
			defer close(done)
			os.Args = []string{"honeycomb", "version"}
			main()
			if err := pw.Close(); err != nil {
				t.Fatal(err)
			}
		}()

		var buf bytes.Buffer
		if _, err = io.Copy(&buf, pr); err != nil {
			t.Fatal(err)
		}
		<-done

		want := "honeycomb version (devel) (MIT LICENSE)"
		got := strings.TrimSpace(buf.String())
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
