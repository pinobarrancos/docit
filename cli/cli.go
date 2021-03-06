package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	// Version the version of the prog.
	Version = "No Version Provided"
	// Buildstamp the time the prog was built.
	Buildstamp = "No Buildstamp provided"
	// Githash the git commit hash
	Githash = "No Githash provided"
	prog    = "docit"
)

// Flags read from the command line.
var (
	index   string
	root    string
	port    string
	name    string
	help    string
	members string
)

// Default flag values when not supplied from the command line.
const (
	indexDefault   = "index.md"
	rootDefault    = "./"
	portDefault    = "8080"
	nameDefault    = "docit"
	helpDefault    = "help.md"
	MembersDefault = ""
)

// Arg contains the flag values read from the command line
// or the default flag values if they are not supplied.
type Arg struct {
	Index      string
	Root       string
	Port       string
	Name       string
	Help       string
	MemberFile string
}

// Usage prints how to use the program.
func Usage() {
	fmt.Println("Usage: ")
	fmt.Println("")
	fmt.Printf("    %s serve -index=Readme.md -help=help.md -root=doc -port=8080\n", prog)
	fmt.Printf("    %s serve -index=landing/index.md -root=home/Projects -port=8080\n", prog)
	fmt.Printf("    %s version\n", prog)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("    serve    - serves files from given root dir")
	fmt.Println("    version  - prints the version of the program")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Printf("    %s serve \n", prog)
	fmt.Printf("    %s serve -index=Readme.md\n", prog)
	fmt.Printf("    %s serve -help=help.md\n", prog)
	fmt.Printf("    %s serve -name=ProjectName\n", prog)
	fmt.Printf("    %s serve -members=member-file\n", prog)
	fmt.Printf("    %s serve -root=test\n", prog)
	fmt.Printf("    %s serve -index=index.md -root=doc\n", prog)
	fmt.Printf("    %s serve -index=index.md -root=doc -port=5000\n", prog)
	fmt.Printf("    %s serve -index=landing/index.md -root=test/Projects -help=landing/help.md", prog)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Version")
	PrintVersion()
}

// PrintVersion prints the version of the prog.
func PrintVersion() {
	fmt.Printf("    Version        : %s\n", Version)
	fmt.Printf("    UTC Build Time : %s\n", Buildstamp)
	fmt.Printf("    Git Commit Hash: %s\n", Githash)
}

func setFlags() *flag.FlagSet {
	f := flag.CommandLine
	f.Usage = Usage

	f.StringVar(&index, "index", indexDefault, "Page to show for '/'")
	f.StringVar(&root, "root", rootDefault, "Root directory to serve files from")
	f.StringVar(&port, "port", portDefault, "The port")
	f.StringVar(&name, "name", nameDefault, "The name of the project")
	f.StringVar(&help, "help", helpDefault,
		"The help page to show when a user clicks on the help link")
	f.StringVar(&members, "members", MembersDefault,
		"Path to a file containing emails")

	return f
}

// Parse returns the command and a context with
// the parsed flags from the command line.
func Parse() (string, Arg) {
	argLen := len(os.Args)

	f := setFlags()

	if argLen < 2 {
		return "", Arg{}
	}

	cmd := os.Args[1]
	if cmd != "serve" && cmd != "version" {
		return "", Arg{}
	}

	for _, flag := range os.Args[2:] {
		if !strings.HasPrefix(flag, "-") {
			fmt.Println("flag should be prefixed with a - sign: ", flag)
			return "", Arg{}
		}
	}

	_ = f.Parse(os.Args[2:])

	return cmd, Arg{
		Index:      index,
		Root:       root,
		Port:       port,
		Name:       name,
		Help:       help,
		MemberFile: members,
	}
}
