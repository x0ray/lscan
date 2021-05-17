// main entry to lscan a Go language source file literal scanner utility
package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
Program:   	lscan.go
Component: 	dm - demo programs
Language:  	go
Support:   	david - nzkiwi1g@gmail.com
Author:    	David A. Fahey - whome God preserve.
Purpose:   	To scan a Go language source file literal scanner utility.

History:   	14May2021 Initial coding                                    DAF

Notes:      Regular expression solver: https://regexr.com/
			Golang regexp: https://golang.org/pkg/regexp/
			Golang token: https://golang.org/pkg/go/token/

Build and run:
			cd .../src/github.com/x0ray/lscan
			go mod init github.com/x0ray/lscan
			go: creating new go.mod: module github.com/x0ray/lscan

			go build ./..
			./lscan

Testing:
			cd .../src/github.com/x0ray/lscan
			go test . -v


Output:
	❯ ./lscan -x xxx.go -s yyy.go -list
	2021/05/14 02:51:28 WARN index key: '.flow.initial.id.not.valid.log' at: xxx.go:70:60 was not added, already occurs at: xxx.go:63:60
	2021/05/14 02:51:28 Index file keys from: 'lscan.go' follow
	Num     Location        Matched Literal
	0001    lscan.go:65:60  .flow.client.extract.failed.log
	0002    lscan.go:66:60  .flow.file.uri.short.error.log
	0003    lscan.go:67:60  .flow.uri.prefix.not.valid.log
	0004    lscan.go:69:60  .flow.name.not.valid.log
	0005    lscan.go:71:60  .flow.id.not.valid.type.log
	0006    lscan.go:63:60  .flow.initial.id.not.valid.log
	0007    lscan.go:64:60  .flow.unmarshal.failed.error.log
	0008    lscan.go:72:60  .flow.id.not.valid.number.dependencies.log
	0009    lscan.go:62:60  .flow.expected.status.error.log
	0010    lscan.go:68:60  .flow.invalid.dependency.ID.log
	2021/05/14 02:51:28 Scan file keys from: 'testinput.go' follow
	Num     Location        Matched Literal
	0001    testinput.go:5:18+0     .flow.expected.status.error.log
	0002    testinput.go:5:18+3     .flow.uri.prefix.not.valid.log
	0003    testinput.go:5:18+4     .flow.invalid.dependency.ID.log
	0004    testinput.go:5:18+5     .flow.name.notvalid.log
	0005    testinput.go:5:18+6     .flow.initial.id.not.valid.log
	0006    testinput.go:5:18+8     .flow.id.not.valid.number.dependencies.log
	0007    testinput.go:5:18+1     .flow.unmarshal.failed.error.log
	0008    testinput.go:5:18+2     .flow.client.extract.failed.log
	0009    testinput.go:5:18+7     .flow.id.not.valid.type.log
	2021/05/14 02:51:28 WARN index key: '.flow.file.uri.short.error.log' at: xxx.go:66:60 not found in scan file: 'yyy.go'
	2021/05/14 02:51:28 WARN index key: '.flow.name.not.valid.log' at: xxx.go:69:60 not found in scan file: 'yyy.go'
*/

const (
	note = `
	This command compares selected literal strings between two Go language source
	code files. The literals are selected using a regular expression pattern. If 
	the literal strings from the first file are all found in the second files no
	messages are produced and the return code is zero, otherwise a non zero return
	code is produced along with messages on stderr.
`
	pgm     = "LScan"
	ver     = "0.0.8"
	verdate = "16May2021"
)

var opt struct {
	indexFile string // input Go source file to create literal index from
	scanFile  string // file to scan for occurrences of index values
	indexRx   string // regular expression used to identify index file literals
	scanRx    string // regular expression used to identify scan file literals
	debug     bool   // debug flag
	listIndex bool   // only list the index
	verbose   bool   // display progress information
	version   bool   // display program version and stop
	help      bool   // display the help text and stop
}

func init() {
	// setup flag options
	flag.StringVar(&opt.indexFile, "index", "", "Input Go source file to used to create literal index")
	flag.StringVar(&opt.indexFile, "x", "", "Alias for -index option")
	flag.StringVar(&opt.scanFile, "scan", "", "File to scan for occurrences of index values")
	flag.StringVar(&opt.scanFile, "s", "", "Alias for -scan option")
	flag.StringVar(&opt.indexRx, "indexrx", `\A\.[a-zA-Z0-9\.]+`, "Regular expression used to identify index file literals")
	flag.StringVar(&opt.indexRx, "i", `\A\.[a-zA-Z0-9\.]+`, "Alias for -indexrx option")
	flag.StringVar(&opt.scanRx, "scanrx", `\.[a-zA-Z0-9\.]+`, "Regular expression used to identify scan file literals")
	flag.StringVar(&opt.scanRx, "c", `\.[a-zA-Z0-9\.]+`, "Alias for -scanrx option")
	flag.BoolVar(&opt.debug, "debug", false, "Provide debugging trace output")
	flag.BoolVar(&opt.listIndex, "list", false, "List the contents of the index and scan index")
	flag.BoolVar(&opt.verbose, "verbose", false, "Display progress information on the log")
	flag.BoolVar(&opt.version, "version", false, "Display program version information")
	flag.BoolVar(&opt.version, "v", false, "Alias for -version option")
	flag.BoolVar(&opt.help, "help", false, "Display program help text")
	flag.BoolVar(&opt.help, "h", false, "Alias for -help option")
}

func initOpt() {
	opt.indexFile = ""
	opt.scanFile = ""
	opt.indexRx = ""
	opt.scanRx = ""
	opt.debug = false
	opt.listIndex = false
	opt.verbose = false
	opt.version = false
	opt.help = false
}

func main() {
	rc := body()
	os.Exit(rc) // us RC outside body to allow lscan_test.go to test body()
}

func body() int {
	var (
		err       error
		scanRx    *regexp.Regexp
		indexRx   *regexp.Regexp
		file      *os.File
		inBytes   []byte
		scnr      scanner.Scanner
		fset      *token.FileSet
		tokenFile *token.File
		index     map[string]string
		scan      map[string]string
		rc        int
	)
	initOpt()
	flag.Parse()

	if opt.version {
		fmt.Printf("Program %s version %s as of %s\n", pgm, ver, verdate)
		rc = 0
		goto theend
	}
	if opt.help {
		fmt.Printf("Program %s \n", pgm)
		fmt.Printf("%s\n", note)
		flag.Usage()
		rc = 0
		goto theend
	}

	if opt.verbose {
		log.Printf("LScan started.")
	}

	indexRx, err = regexp.Compile(opt.indexRx)
	if err != nil {
		log.Printf("ERROR Invalid option regular expression value: '%s' for -indexrx, %v\n", opt.indexRx, err)
		rc = 4
		goto theend
	}

	scanRx, err = regexp.Compile(opt.scanRx)
	if err != nil {
		log.Printf("ERROR Invalid option regular expression value: '%s' for -scanrx, %v\n", opt.scanRx, err)
		rc = 8
		goto theend
	}

	// tokenize index file and create a map of the selected literal tokens
	file, err = os.Open(opt.indexFile) // For read access.
	if err != nil {
		log.Printf("ERROR Open failed for index file: '%s', %v\n", opt.indexFile, err)
		rc = 12
		goto theend
	}
	inBytes, err = io.ReadAll(file)
	if err != nil {
		log.Printf("ERROR Read failed for index file: '%s', %v\n", opt.indexFile, err)
		rc = 16
		goto theend
	}
	file.Close()
	index = make(map[string]string)
	fset = token.NewFileSet()                                          // positions are relative to fset
	tokenFile = fset.AddFile(opt.indexFile, fset.Base(), len(inBytes)) // register input "file"
	scnr.Init(tokenFile, inBytes, nil /* no error handler */, scanner.ScanComments)
	if opt.debug {
		log.Printf("Token stream from index file: '%s' follows\n", opt.indexFile)
		log.Printf("Location\tToken\tValue\n")
	}
	for {
		// Repeated calls to Scan yield the token sequence found in the input.
		pos, tok, lit := scnr.Scan()
		if tok == token.EOF {
			break
		}
		if opt.debug {
			log.Printf("%s\t%s\t%s\n", fset.Position(pos), tok, lit)
		}
		if tok == token.STRING {
			lit = strings.Trim(lit, "\"") // strip "" of literal tokens
			if matched := indexRx.MatchString(lit); matched {
				posn := fset.Position(pos).String()
				if ep, ok := index[lit]; ok { // check already there == duplicate
					log.Printf("WARN index key: '%s' at: %s was not added, already occurs at: %s \n", lit, posn, ep)
					rc = 1
				} else {
					index[lit] = posn
				}
			}
		}
	}

	// optionally list out the map of literals found in the index file
	if opt.listIndex {
		log.Printf("Index file keys from: '%s' follow\n", opt.indexFile)
		fmt.Printf("Num\tLocation\tMatched Literal\n")
		i := 0
		for k, v := range index {
			i++
			fmt.Printf("%04d\t%s\t%s\n", i, v, k)
		}
	}

	// optionally tokenize scan scan file and create a map of the selected literal tokens
	if opt.scanFile != "" {
		file, err = os.Open(opt.scanFile) // For read access.
		if err != nil {
			log.Printf("ERROR Open failed for scan file: '%s', %v\n", opt.scanFile, err)
			rc = 20
			goto theend
		}
		inBytes, err = io.ReadAll(file)
		if err != nil {
			log.Printf("ERROR Read failed for scan file: '%s', %v\n", opt.scanFile, err)
			rc = 24
			goto theend
		}
		file.Close()
		scan = make(map[string]string)
		fset = token.NewFileSet()                                         // positions are relative to fset
		tokenFile = fset.AddFile(opt.scanFile, fset.Base(), len(inBytes)) // register input "file"
		scnr.Init(tokenFile, inBytes, nil /* no error handler */, scanner.ScanComments)

		// Repeated calls to Scan yield the token sequence found in the input.
		if opt.debug {
			log.Printf("Token stream from scan file: '%s' follows\n", opt.scanFile)
			log.Printf("Location\tToken\tValue\n")
		}
		for {
			pos, tok, lit := scnr.Scan()
			if tok == token.EOF {
				break
			}
			if opt.debug {
				log.Printf("%s\t%s\t%s\n", fset.Position(pos), tok, lit)
			}

			if tok == token.STRING {
				lit = strings.Trim(lit, "\"")
				if matched := scanRx.MatchString(lit); matched {
					// check literal to determine if it has multiple matches
					all := scanRx.FindAllString(lit, -1)
					for i, v := range all {
						scan[v] = fset.Position(pos).String() + "+" + strconv.Itoa(i)
					}
				}
			}
		}

		// optionally list out the map of literals found in the scan file
		if opt.listIndex {
			log.Printf("Scan file keys from: '%s' follow\n", opt.scanFile)
			fmt.Printf("Num\tLocation\tMatched Literal\n")
			i := 0
			for k, v := range scan {
				i++
				fmt.Printf("%04d\t%s\t%s\n", i, v, k)
			}
		}

		// check scan file contains all values from index file
		for k, v := range index {
			if _, ok := scan[k]; !ok {
				log.Printf("WARN index key: '%s' at: %s not found in scan file: '%s'\n", k, v, opt.scanFile)
				rc = 2
			}
		}
	}

theend:
	if opt.verbose {
		log.Printf("LScan ended.")
	}
	return rc
}
