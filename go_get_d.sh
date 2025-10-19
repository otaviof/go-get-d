# go_get_d SHell function to invoke "go-get-d", collect the output of the command
# and switch the current working directory into it.
function go_get_d() {
    local project_path
    project_path=$(go-get-d --path "${1}")
    if [ $? -ne 0 ]; then
        echo "# Error: Failed to retrieve path using go-get-d." >&2
        return 1
    fi
    echo "# Changing directory to ${project_path}"
    cd "${project_path}"
}
