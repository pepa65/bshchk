# bshchk (https://github.com/pepa65/bshchk)
{{.DepsVar}}=({{.Deps}})
non_ok=()

for d in ${{.DepsVar}}
do command -v $d >/dev/null 2>&1 || non_ok+=$d
done

((${#non_ok[@]})) &&
	cat <<-EOM >&2 &&
		This program requires these commands to be installed:
		\${{.DepsVar}}

		These are still missing:
		\$non_ok

		Please install the corresponding packages first.
	EOM
	exit 1

unset non_ok{{if .UnsetDeps}}
unset {{.DepsVar}}
{{end}}# Dependencies are OK at this point
