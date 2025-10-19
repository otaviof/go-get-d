# go_get_d Shell function to invoke "go-get-d", collect the output of the command
# and switch the current working directory into it. The function takes a single
# input module name.
go_get_d() {
    if [[ " ${@} " =~ " --help " ]]; then
        go-get-d --help
        return 0
    fi
    if [ -z "${1}" ]; then
        echo "# Error: No module provided. " \
            "Use 'go-get-d --help' for more information!" >&2
        return 1
    fi
    local output
    output=$(go-get-d --path ${@}) || go_get_d_exit="${?}"
    if [ "${go_get_d_exit:-0}" -ne 0 ]; then
        echo "# Error: Failed to retrieve path using go-get-d." \
            "Exit code '${go_get_d_exit}'!" >&2
        echo ${output} >&2
        return 1
    fi
    echo "# Changing directory into '${output}'..."
    cd "${output}" || {
        echo "# Error: Unable to change directory to \"${output}\"!" >&2
        return 1
    }
    return 0
}
