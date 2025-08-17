# From bshchk (https://github.com/pepa65/bshchk)
{{.DepsVar}}=({{.Deps}})
non_ok=()
for d in ${{.DepsVar}}
do command -v $d >/dev/null 2>&1 || non_ok+=$d
done
((${#non_ok[@]})) &&
	cat <<-EOM >&2 &&
		This program requires these commands to be present in PATH:
		${{.DepsVar}}

		These commands are still missing:
		$non_ok

		Please make sure these commands or their corresponding packages are installed."
	EOM
	exit 1

unset non_ok {{.DepsVar}}

# Dependencies are OK beyond this point - bshchk

