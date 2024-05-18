package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExecute(t *testing.T) { //nolint:paralleltest
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "success",
			args:    []string{"honeycomb", "version"},
			wantErr: false,
		},
		{
			name:    "fail",
			args:    []string{"no-exist-subcommand", "--no-exist-option"},
			wantErr: true,
		},
	}
	for _, tt := range tests { //nolint:paralleltest
		os.Args = tt.args
		t.Run(tt.name, func(t *testing.T) { //nolint:paralleltest
			err := Execute()
			gotErr := err != nil
			if tt.wantErr != gotErr {
				t.Errorf("expected error return %v, got %v: %v", tt.wantErr, gotErr, err)
			}
		})
	}
}

func TestExecute_Version(t *testing.T) { //nolint:paralleltest
	tests := []struct {
		name   string
		args   []string
		stdout []string
	}{
		{
			name:   "success",
			args:   []string{"honeycomb", "version"},
			stdout: []string{"honeycomb version  (MIT LICENSE)", ""},
		},
	}
	for _, tt := range tests { //nolint:paralleltest
		tt := tt

		t.Run(tt.name, func(_ *testing.T) { //nolint:paralleltest
			orgStdout := os.Stdout
			orgStderr := os.Stderr
			pr, pw, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			os.Stdout = pw
			os.Stderr = pw
			os.Args = tt.args

			if err = Execute(); err != nil {
				t.Fatal(err)
			}
			if err := pw.Close(); err != nil {
				t.Fatal(err)
			}

			os.Stdout = orgStdout
			os.Stderr = orgStderr
			buf := bytes.Buffer{}
			if _, err = io.Copy(&buf, pr); err != nil {
				t.Fatal(err)
			}
			defer pr.Close() //nolint

			got := strings.Split(buf.String(), "\n")
			if diff := cmp.Diff(tt.stdout, got); diff != "" {
				t.Errorf("value is mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
