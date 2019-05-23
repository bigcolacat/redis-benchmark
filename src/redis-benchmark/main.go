package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var options = ProcessArguments()
	if(options.ShowHelp) {
		fmt.Println(options.HelpText)
		os.Exit(1)
	}
	PrintOptions(options)
	rand.Seed(time.Now().UnixNano())

	for _, test := range options.Tests {
		if(test == "PING") {
			Bench(test, &options, nil, func(buf []byte) []byte {
				return AppendCommand(buf, "PING")
			})
		}
		if(test == "SET") {
			Bench(test, &options, nil, func(buf []byte) []byte {
				return AppendCommand(buf, "SET", "key:string", "val")
			})
		}
		if(test == "GET") {
			Bench(test, &options, nil, func(buf []byte) []byte {
				return AppendCommand(buf, "GET", "key:string")
			})
		}
		if(test == "GEOADD") {
			Bench(test, &options, nil, func(buf []byte) []byte {
				return AppendCommand(buf, "GEOADD", "key:geo",
					strconv.FormatFloat(rand.Float64()*360-180, 'f', 7, 64),
					strconv.FormatFloat(rand.Float64()*170-85, 'f', 7, 64),
					strconv.Itoa(rand.Int()))
			})
		}
		if(test == "GEORADIUS") {
			Bench(test, &options, nil, func(buf []byte) []byte {
				return AppendCommand(buf, "GEORADIUS", "key:geo",
					strconv.FormatFloat(rand.Float64()*360-180, 'f', 7, 64),
					strconv.FormatFloat(rand.Float64()*170-85, 'f', 7, 64),
					"10", "km")
			})
		}
	}
}