package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Options represents various options used by the Bench() function.
type Options struct {
	ShowHelp    bool
	Host        string
	Port        int
	Password    string

	Requests int
	Clients  int
	Pipeline int
	Quiet    bool
	CSV      bool
	Stdout   io.Writer
	Stderr   io.Writer

	Tests       []string
	HelpText    string
}

// DefaultsOptions are the default options used by the Bench() function.
var defaultOptions = Options{
	ShowHelp:    false,
	Host:        "127.0.0.1",
	Port:        6379,

	Requests: 100000,
	Clients:  50,
	Pipeline: 1,
	Quiet:    false,
	CSV:      false,
	Stdout:   os.Stdout,
	Stderr:   os.Stderr,

	Tests:       []string{"PING"},
	HelpText: 	 helpText,
}

var helpText = `Usage: redis-benchmark [options]

  -h, --help          	displays help
  -H, --host String   	Server hostname - default: 127.0.0.1
  -a, --password String Password to use when connecting to the server
  -P, --numreq Int 	    Pipeline <numreq> requests. Default 1 (no pipeline).
  -n, --requests Int  	Total number of requests - default: 100000
  -c, --clients Int   	Number of parallel connections - default: 50
  -t, --tests Array   	Only run the comma separated list of tests. The test names are the same as the ones produced as output. - default: PING
  -p, --port Int      	Server port - default: 6379
  
Version 1.0.0`

func buildHelp(message string) Options {
	return Options{ShowHelp: true, HelpText: message + `

` + helpText}
}

var emptyOptions = Options{}

func parseNumber(args []string, index *int) (int, Options) {
	*index++
	if *index >= len(args) {
		return -1, buildHelp("Error: Incorrect parameters specified")
	}
	number, err := strconv.Atoi(args[*index])
	if err != nil {
		return -1, buildHelp("Error: Invalid type for parameter")
	}

	if number < 0 {
		return -1, buildHelp("Error: Parameter value must be non-negative")
	}

	return number, emptyOptions
}

func ParseArguments(argument []string) Options {
	options := defaultOptions
	var errOptions Options

	args := argument[1:]
	for i:=0; i<len(args); i++ {
		if args[i] == "--help" || args[i] == "-h" {
			options.ShowHelp = true
			return options
		} else if(args[i] == "--host" || args[i] == "-H") {
			i++
			if(i >= len(args)) {
				return buildHelp("No host is provided")
			}
			options.Host = args[i]
		} else if(args[i] == "--password" || args[i] == "-a") {
			i++
			if(i >= len(args)) {
				return buildHelp("No password is provided")
			}
			options.Password = args[i]
		} else if(args[i] == "--numreq" || args[i] == "-P") {
			options.Pipeline, errOptions = parseNumber(args, &i)
			if(errOptions.ShowHelp) {
				return errOptions
			}
		} else if(args[i] == "--requests" || args[i] == "-n") {
			options.Requests, errOptions = parseNumber(args, &i)
			if(errOptions.ShowHelp) {
				return errOptions
			}
		} else if(args[i] == "--clients" || args[i] == "-c") {
			options.Clients, errOptions = parseNumber(args, &i)
			if(errOptions.ShowHelp) {
				return errOptions
			}
		} else if(args[i] == "--tests" || args[i] == "-t") {
			i++
			if(i >= len(args)) {
				return buildHelp("No test is provided")
			}
			options.Tests = strings.Split(args[i], ",")
			for i := range options.Tests {
				options.Tests[i] = strings.ToUpper(options.Tests[i])
			}
		} else if(args[i] == "--port" || args[i] == "-p") {
			options.Port, errOptions = parseNumber(args, &i)
			if(errOptions.ShowHelp) {
				return errOptions
			}
		}
	}

	return options
}

// ProcessArguments gets arguments from the command line and parses them
func ProcessArguments() Options {
	return ParseArguments(os.Args)
}

func PrintOptions(options Options) {
	fmt.Printf("Host: %s, Port: %d, PipelinedRequests: %d, " +
		"Requests: %d, Connections: %d, Tests: %s\n",
		options.Host, options.Port,
		options.Pipeline, options.Requests,
		options.Clients, options.Tests)
}