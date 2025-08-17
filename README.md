# bshchk 0.2.3
**Dependency checker for bash scripts, to assure all external commands will be present when run.**

## Usage
```sh
bshchk BASH_SOURCE OUTFILE

# Will write to stdout
bshchk BASH_SOURCE

# For piping in, will read from stdin
bshchk - OUTFILE

# Will read from stdin, and write to stdout
bshchk

# Will only show the deps as a bash array
bshchk --deps-only

# Will only show the additional code
bshchk --deps-code
```

### Tags in source
* To explicitly add curl as a dependency, add this line to the source:
  `#bshchk:add-cmd curl`
* To prevent curl from being seen as a dependency, add this line to the source:
  `#bshchk:ignore-cmd curl`
* Multiple commands can be added and ignored as dependencies, like:
  `#bshchk:add-cmd curl wget`

## Help
```
bshchk v0.2.3 - Dependency checker for bash scripts
Usage: bshchk [--version] [--deps-only] [--deps-code] [--deps-name DEPS-NAME] [--ignore-shebang] [SOURCE [OUTFILE]]

Positional arguments:
  SOURCE                 If given as '-' or '': read from stdin
  OUTFILE                If not given: print to stdout

Options:
  --version, -V          Print version
  --deps-only, -d        Print dependencies as a bash array
  --deps-code, -c        Print additional code
  --deps-name DEPS-NAME, -n DEPS-NAME
                         Override deps variable name [default: deps]
  --ignore-shebang, -i   Ignore shebang requirement
  --help, -h             display this help and exit
```

## License
* **GPLv3+**
* Copyright (C) 2024 blek! <me@blek.codes> 2025 github.com/pepa65
* See [LICENSE]
