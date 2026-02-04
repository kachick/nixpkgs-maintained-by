package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestNixFilter(t *testing.T) {
	absMockPath, err := filepath.Abs("mock_pkgs.nix")
	if err != nil {
		t.Fatalf("Failed to get absolute path of mock_pkgs.nix: %v", err)
	}

	tests := []struct {
		name         string
		maintainerID string
		want         []string
	}{
		{
			name:         "Find packages for user1",
			maintainerID: "user1",
			want:         []string{"packageSingleMaintainer", "packageSharedMaintainers"},
		},
		{
			name:         "Find packages for user2",
			maintainerID: "user2",
			want:         []string{"packageSharedMaintainers"},
		},
		{
			name:         "Find no packages for unknown user",
			maintainerID: "unknown",
			want:         []string{}, // Empty list (JSON [])
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We construct a command similar to main.go but injecting our mock pkgs
			// --arg pkgs 'import /path/to/mock_pkgs.nix'
			cmd := exec.Command("nix-instantiate",
				"--eval", "--json",
				"--argstr", "maintainerId", tt.maintainerID,
				"--arg", "pkgs", "import "+absMockPath,
				"-E", nixExpression,
			)

			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				t.Fatalf("nix-instantiate failed: %v\nStderr: %s", err, stderr.String())
			}

			var got []string
			if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
				t.Fatalf("Failed to parse JSON output: %v\nOutput: %s", err, stdout.String())
			}

			// Normalize for comparison
			if got == nil {
				got = []string{}
			}
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
