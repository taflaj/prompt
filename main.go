// main.go

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/taflaj/prompt/prompt"
)

const version = "1.3.0"

func init() {}

func doHelp() {
	fmt.Printf("Usage: %v <command>\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  help     Displays this message.")
	fmt.Println("  init     Displays text to be used inside a shell script.")
	fmt.Println("  show     Displays the prompt according to the parameters.")
	fmt.Printf("  version  Displays the current version (%v).\n", version)
}

func main() {
	if len(os.Args) == 1 {
		doHelp()
	} else {
		switch os.Args[1] {
		case "help":
			doHelp()
		case "init":
			fmt.Println("PS0='${t:0:$((t=$(date +%s%N),0))}'")
			fmt.Println("PROMPT_COMMAND=set_prompt")
			fmt.Println("set_prompt() {")
			fmt.Println("  PS1=\"$(code=$? jobs=$(jobs -p | wc -l) options=$PROMPT time=$t now=$(date +%s%N) prompt show)\"")
			fmt.Println("  t=0")
			fmt.Println("}")
			fmt.Println("trap 'echo -ne \"\\e[0m\"' DEBUG")
		case "show":
			prompt.Show()
		case "version":
			fmt.Printf("%v %v on %v/%v with %v\n", os.Args[0], version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		default:
			fmt.Println("Invalid command")
			doHelp()
		}
	}
}
