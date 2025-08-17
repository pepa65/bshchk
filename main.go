package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

const name = "bshchk"
const version = "0.2.0"

type args struct {
	Source        string `arg:"positional" help:"If given as '-' or '': read from stdin" default:""`
	Outfile       string `arg:"positional" help:"If not given: print to stdout" default:""`
	Version       bool   `arg:"-V,--version" help:"Print version"`
	DepsOnly      bool   `arg:"-d,--deps-only" help:"Print dependencies as a bash array" default:false`
	DepsCode      bool   `arg:"-c,--deps-code" help:"Print additional code" default:false`
	DepsName      string `arg:"-n,--deps-name" help:"Override deps variable name" default:"deps"`
	IgnoreShebang bool   `arg:"-i,--ignore-shebang" help:"Ignore shebang requirement" default:false`
}

func (args) Description() string {
	return fmt.Sprintf("%s v%s - Dependency checker for bash scripts", name, version)
}

func main() {
	var args args
	// Help implicit
	arg.MustParse(&args)
	if args.Version {
		fmt.Printf("%s v%s\n", name, version)
		os.Exit(0)
	}

	var file []byte
	if (args.Source != "") && (args.Source != "-") {
		f, err := os.ReadFile(args.Source)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "File %s does not exist!", args.Source)
				os.Exit(1)
			}
			panic(err)
		}
		file = f
	} else {
		data, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Printf("Coudln't read from stdin: %s", err.Error())
			os.Exit(1)
		}
		file = data
	}

	code := string(file)
	found, err := find(code)
	if err != nil {
		panic(err)
	}

	codelines := strings.Split(code, "\n")

	if len(codelines) < 2 {
		fmt.Fprintf(os.Stderr, "The code must have at least two lines!\n")
		os.Exit(3)
	}

	shebang := codelines[0]

	if !((shebang == "#!/bin/bash") || (shebang == "#!/usr/bin/bash") || (shebang == "#!/bin/bash") || (shebang == "#!/usr/bin/env bash")) && (!args.IgnoreShebang) {
		fmt.Fprintf(os.Stderr, "The code must start with a bash shebangs: #!/bin/bash OR #!/ust/bin/bash OR  #!/usr/bin/env bash\n")
		os.Exit(2)
	}

	var gen string
	if args.DepsOnly {
		gen = strings.Split(gencode(found), "\n")[1]
	} else if args.DepsCode {
		gen = gencode(found)
	} else {
		gen = shebang + "\n\n" + gencode(found) + "\n\n" + strings.Join(codelines[1:], "\n")
	}

	if args.Outfile == "" {
		fmt.Printf("%s\n", gen)
	} else {
		os.WriteFile(args.Outfile, []byte(gen), os.FileMode(0o755))
	}
}
