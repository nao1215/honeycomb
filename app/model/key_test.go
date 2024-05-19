package model

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

		oldXdgConfigHome := xdg.ConfigHome
		xdg.ConfigHome = filepath.Join("/", "home", "user", ".config")
		defer func() {
			xdg.ConfigHome = oldXdgConfigHome
		}()

		want := filepath.Join("/", "home", "user", ".config", "honeycomb", "private_key")
		if got := NSecretKeyFilePath(); got != want {
			t.Errorf("PrivateKeyFilePath() = %v, want %v", got, want)
		}
	})
}

func TestReadPrivateKey(t *testing.T) { //nolint:paralleltest
	t.Run("return private key", func(t *testing.T) { //nolint:paralleltest
		oldXdgConfigHome := xdg.ConfigHome
		xdg.ConfigHome = "testdata"
		defer func() {
			xdg.ConfigHome = oldXdgConfigHome
		}()

		want := NSecretKey("nsec10ggv7zysc7szq2q75nrdepu5ul8x0ashyrl28rl9x6g35p4whvqsumqqpv")
		if err := WriteNSecretKey(want); err != nil {
			t.Fatalf("WritePrivateKey() error = %v, want nil", err)
		}
		defer func() {
			if err := WriteNSecretKey(""); err != nil {
				t.Fatalf("WritePrivateKey() error = %v, want nil", err)
			}
		}()

		got, err := ReadNSecretKey()
		if err != nil {
			t.Fatalf("ReadPrivateKey() error = %v, want nil", err)
		}
		if got != want {
			t.Errorf("ReadPrivateKey() = %v, want %v", got, want)
		}
	})

	t.Run("return error if the file does not exist", func(t *testing.T) { //nolint:paralleltest
		oldXdgConfigHome := xdg.ConfigHome
		xdg.ConfigHome = filepath.Join("testdata", "not_exist")
		defer func() {
			xdg.ConfigHome = oldXdgConfigHome
		}()

		if _, err := ReadNSecretKey(); err == nil {
			t.Error("ReadPrivateKey() error = nil, want not nil")
		}
	})

	t.Run("return error if the private key is empty", func(t *testing.T) { //nolint:paralleltest
		oldXdgConfigHome := xdg.ConfigHome
		xdg.ConfigHome = "testdata"
		defer func() {
			xdg.ConfigHome = oldXdgConfigHome
		}()

		want := NSecretKey("")
		if err := WriteNSecretKey(want); err != nil {
			t.Fatalf("WritePrivateKey() error = %v, want nil", err)
		}

		_, err := ReadNSecretKey()
		if !errors.Is(err, ErrEmptyPrivateKey) {
			t.Errorf("ReadPrivateKey() error = %v, want ErrEmptyPrivateKey", err)
		}
	})
}
