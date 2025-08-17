# bshchk (https://github.com/pepa65/bshchk)
{{.DepsVar}}=({{.Deps}})
non_ok=()

for d in ${{.DepsVar}}
do command -v $d >/dev/null 2>&1 || non_ok+=$d
done

if ((${#non_ok[@]}))
then
	>&2 echo "RDC Failed!"
	>&2 echo "  This program requires these commands:"
	>&2 echo "  > ${{.DepsVar}}"
	>&2 echo "    --- "
	>&2 echo "  From which, these are missing:"
	>&2 echo "  > $non_ok"
	>&2 echo "Make sure that those are installed and are present in \$PATH."
	exit 1
fi

unset non_ok{{if .UnsetDeps}}
unset {{.DepsVar}}
{{end}}# Dependencies are OK at this point
