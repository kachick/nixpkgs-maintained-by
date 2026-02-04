package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	//go:embed filter.nix
	nixExpression string

	version = "0.1.0"
)

func main() {
	maintainerID := flag.String("id", "", "Maintainer ID in nixpkgs")
	jsonOutput := flag.Bool("json", false, "Output results in JSON format")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("nixpkgs-maintained-by %s\n", version)
		return
	}

	if *maintainerID == "" {
		fmt.Fprintf(os.Stderr, "Usage: nixpkgs-maintained-by -id <maintainer-id> [--json] [--version]\n")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Searching for packages maintained by %s...\n", *maintainerID)

	cmd := exec.Command("nix-instantiate", "--eval", "--json", "--argstr", "maintainerId", *maintainerID, "-E", nixExpression)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Stderr: %s\n", stderr.String())
		os.Exit(1)
	}

	var packages []string
	if err := json.Unmarshal(stdout.Bytes(), &packages); err != nil {
		fmt.Fprintf(os.Stderr, "JSON Error: %v\n", err)
		os.Exit(1)
	}

	if *jsonOutput {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(packages); err != nil {
			fmt.Fprintf(os.Stderr, "JSON Encoding Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if len(packages) == 0 {
		fmt.Fprintf(os.Stderr, "No packages found for: %s\n", *maintainerID)
		return
	}

	for _, pkg := range packages {
		fmt.Println(pkg)
	}
}
