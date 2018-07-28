package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/awoodbeck/strftime"

	"github.com/fatih/color"
	pflag "github.com/janmir/pflag"
)

const (
	_date    = "Jan 2 15:04 MST 2006"
	_build   = ""
	_version = ""
	_release = ""
	_git     = ""

	_name = "Buildy"
)

var (
	rfg = color.New(color.FgRed, color.Bold).SprintfFunc()
	gfg = color.New(color.FgGreen).SprintfFunc()

	cflag = pflag.CustomFlagSet(rfg("%s", _name), false, errors.New("\n "))

	date    = cflag.String("date", _date, "Date string format. [Mon Jan 2 15:04:05 -0700 MST 2006]")
	build   = cflag.String("build", _build, "Build string. [Alpha, Beta, Dev ...]")
	version = cflag.String("version", _version, "Version string SemVer/Custom string.[v0.0.0, 0.0.0-%m%d%Y]")
	release = cflag.String("release", _release, "Release name.[\"Oneric Ocelot\"]")
	git     = cflag.String("git", _git, "Git command to execute and append.")

	sname    = ""
	sdate    = "Date: "
	sbuild   = "Build: "
	sversion = "Version: "
	srelease = "Release name: "
	sgit     = ""

	errLogger, stdLogger *log.Logger

	/*
	* ╭──App Name───────────────────────────────╮
	* │                                         │
	* │  Version:                               │
	* │  Build:                                 │
	* │  Author:                                │
	* │                                         │
	* │  ░░░░░░░░░░░░░░░░ Git ░░░░░░░░░░░░░░░░  │
	* │  ahasjds7 jp Message                    │
	* │  asdasd8w jp Message                    │
	* │  asd78qws jp Another Message            │
	* │                                         │
	* │ ┌─Revision History (3 mos)────────────┐ │
	* │ │▪▪▪▪▫▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▫▪▪▪▫▫▪│ │
	* │ │▪▪▪▪▪▫▫▪▪▫▪▫▪▪▪▪▪▫▫▪▪▪▪▫▪▫▪▪▪▪▪▪▪▪▫▫▪│ │
	* │ │▪▪▫▪▪▫▪▪▫▪▫▪▪▪▫▪▫▫▫▪▪▫▪▪▪▫▪▪▫▪▪▫▪▪▫▫▪│ │
	* │ │▪▪▫▪▪▪▫▪▫▪▪▫▪▫▫▫▪▫▪▪▪▪▪▪▫▪▪▫▪▪▫▪▪▪▫▫▪│ │
	* │ │▪▪▪▪▪▫▫▪▪▪▪▪▪▪▪▪▪▫▪▪▪▪▪▪▪▫▪▪▫▪▫▪▪▪▫▫▪│ │
	* │ │▪▪▪▪▪▪▪▫▪▫▪▪▫▪▪▪▪▪▫▪▪▪▫▪▪▫▪▪▪▪▪▪▫▪▪▫▪│ │
	* │ │▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪▫▪▫▫▪▪▪▪▪▫│ │
	* │ └─────────────────────────────────────┘ │
	* ╰─────────────────────────────────────────╯
	 */
	mainString  = ""
	top         = "─"
	topleft     = "\n\t╭"
	bottomleft  = "\n\t╰"
	topright    = "╮"
	bottomright = "╯"
	leftside    = "\n\t│"
	rightside   = "│"
	space       = " "
	padleft     = 3
	padright    = 3
	mainMax     = padleft + padright
)

func init() {
	cflag.Parse(os.Args[1:])

	errLogger = log.New(os.Stderr, rfg("%s", " ▶ "), log.Ldate|log.Lshortfile)
	stdLogger = log.New(os.Stdout, "", 0)
}

func main() {
	logger(rfg("%s", "Debug Mode"))

	//Add name
	sname += _name

	//Measure length first before appending
	now := time.Now().Local()
	if *date != "" {
		sdate += now.Format(*date)

		logger(sdate)
	}

	if *build != "" {
		sbuild += *build

		logger(sbuild)
	}

	if *version != "" {
		sversion += strftime.Format(&now, *version)

		logger(sversion)
	}

	if *release != "" {
		srelease += *release

		logger(srelease)
	}

	if *git != "" {

	}

	//get max width
	mainMax = max(len(sname), len(sdate),
		len(sbuild), len(sversion),
		len(srelease), len(sgit))

	logger("len: %d\n", mainMax)

	sname = makeTop(_name)
	appendString(sname)

	sdate = makeBody(sdate)
	appendString(sdate)

	sversion = makeBody(sversion)
	appendString(sversion)

	srelease = makeBody(srelease)
	appendString(srelease)

	appendString(makeBottom())

	//print all to stdOutput
	stdLogger.Println(mainString)
}

func makeBody(str string) string {
	s := ""
	l := padleft + (mainMax - len(str)) + padright
	for i := 0; i <= l; i++ {
		switch {
		case i == 0:
			s += leftside
		case i < padleft:
			s += space
		case i == padleft:
			s += str
		case i == l:
			s += rightside
		default:
			s += " "
		}
	}
	return s
}

func makeBottom() string {
	s := ""
	l := padleft + (mainMax) + padright

	for i := 0; i < l; i++ {
		switch {
		case i == 0:
			s += leftside
		case i == l-1:
			s += rightside
		default:
			s += space
		}
	}

	for i := 0; i <= l; i++ {
		switch {
		case i == 0:
			s += bottomleft
		case i == padleft:
		case i < padleft:
			s += top
		case i == l:
			s += bottomright + "\n"
		default:
			s += top
		}
	}
	return s
}

func makeTop(name string) string {
	s := ""
	l := padleft + (mainMax - len(name)) + padright
	for i := 0; i <= l; i++ {
		switch {
		case i == 0:
			s += topleft
		case i < padleft:
			s += top
		case i == padleft:
			s += rfg("%s", name)
		case i == l:
			s += topright
		default:
			s += top
		}
	}

	l += len(name)
	for i := 0; i < l; i++ {
		switch {
		case i == 0:
			s += leftside
		case i == l-1:
			s += rightside
		default:
			s += space
		}
	}
	return s
}

func max(vals ...int) int {
	mx := vals[0]
	for _, v := range vals {
		if v > mx {
			mx = v
		}
	}

	return mx
}

func appendString(str ...string) {
	for _, v := range str {
		mainString += v
	}
}

func logger(args ...interface{}) {
	if _debug {
		if len(args) > 1 && (strings.HasPrefix(args[0].(string), "%") ||
			strings.Contains(args[0].(string), "%")) {
			//format print
			errLogger.Printf(args[0].(string), args[1:]...)
		} else {
			//line print
			errLogger.Println(args...)
		}
	}
}

func catch(err error) {
	if err != nil {
		_, file, no, _ := runtime.Caller(1)
		ss := strings.Split(file, "/")
		file = ss[len(ss)-1]
		report := fmt.Sprintf("%s:%d, %+v", file, no, err)

		log.Fatal(report)
	}
}
