package main
import "fmt"
import "unsafe"

// #include <stdlib.h>
import "C"


//export BashCommandWithTokens
func BashCommandWithTokens(nTokens int) unsafe.Pointer{
    return unsafe.Pointer(C.CString(bashCommandWithTokens(nTokens)))
}

// bashCommandWithTokens returns a bash script that should be executed with nTokens number of bash arguments
//
//  Each argument to bash is evaluated by the shell before becoming a token in the resulting script
//  Example:
//    Given nTokens=2 the returned script will contain `"$(eval echo \"$0\")" "$(eval echo \"$1\")"`
//      and should be evaluated with  `bash -c '"$(eval echo \"$0\")" "$(eval echo \"$1\")"' <command> <arg>'
//    Token evaluation example:
//      "$(eval echo \"$0\"`)" //  given $0='$FOO' and $FOO='arg with spaces" && quotes'
//      -> "$(eval echo \"'$FOO'\")"
//      -> "$(echo \"'arg with spaces" && quotes'\")"
//      -> "arg with spaces\" && quotes" // this is an evaluated and properly quoted token
func bashCommandWithTokens(nTokens int) string {
    commandScript := `"$(eval echo \"$0\")"`
    for i := 1; i < nTokens; i++ {
        commandScript += fmt.Sprintf(` "$(eval echo \"${%d}\")"`, i)
    }
    return fmt.Sprintf(`exec bash -c '%s' "${@:1}"`, commandScript)
}

func main() {}
