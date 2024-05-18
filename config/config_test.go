// Package config read and write configuration files.
package config

import (
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/adrg/xdg"
)

func TestPrivateKeyFilePath(t *testing.T) { //nolint:paralleltest
	t.Run("return file path that store private key", func(t *testing.T) { //nolint:paralleltest
		if runtime.GOOS == "windows" {
			t.Skip("Skip on Windows")
		}

		t.Setenv("XDG_CONFIG_HOME", filepath.Join("/", "home", "user", ".config"))
		xdg.Reload()

		want := filepath.Join("/", "home", "user", ".config", "honeycomb", "private_key")
		if got := PrivateKeyFilePath(); got != want {
			t.Errorf("PrivateKeyFilePath() = %v, want %v", got, want)
		}
	})
}

func TestReadPrivateKey(t *testing.T) { //nolint:paralleltest
	t.Run("return private key", func(t *testing.T) { //nolint:paralleltest
		t.Setenv("XDG_CONFIG_HOME", filepath.Join("testdata", appName, "good_private_key"))
		xdg.Reload()

		want := PrivateKey("nsec10ggv7zysc7szq2q75nrdepu5ul8x0ashyrl28rl9x6g35p4whvqsumqqpv")
		if err := WritePrivateKey(want); err != nil {
			t.Fatalf("WritePrivateKey() error = %v, want nil", err)
		}
		defer func() {
			if err := WritePrivateKey(""); err != nil {
				t.Fatalf("WritePrivateKey() error = %v, want nil", err)
			}
		}()

		got, err := ReadPrivateKey()
		if err != nil {
			t.Fatalf("ReadPrivateKey() error = %v, want nil", err)
		}
		if got != want {
			t.Errorf("ReadPrivateKey() = %v, want %v", got, want)
		}
	})

	t.Run("return error if the file does not exist", func(t *testing.T) { //nolint:paralleltest
		t.Setenv("XDG_CONFIG_HOME", filepath.Join("testdata", appName, "not_exist"))
		xdg.Reload()

		if _, err := ReadPrivateKey(); err == nil {
			t.Error("ReadPrivateKey() error = nil, want not nil")
		}
	})

	t.Run("return error if the private key is empty", func(t *testing.T) { //nolint:paralleltest
		t.Setenv("XDG_CONFIG_HOME", filepath.Join("testdata", appName, "private_key"))
		xdg.Reload()

		want := PrivateKey("")
		if err := WritePrivateKey(want); err != nil {
			t.Fatalf("WritePrivateKey() error = %v, want nil", err)
		}

		_, err := ReadPrivateKey()
		if !errors.Is(err, ErrEmptyPrivateKey) {
			t.Errorf("ReadPrivateKey() error = %v, want ErrEmptyPrivateKey", err)
		}
	})
}
