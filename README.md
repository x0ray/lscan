# lscan - Literal string scanner for Go language programs

This command compares selected literal strings between two Go language source
code files. The literals are selected using a regular expression pattern. If 
the literal strings from the first file are all found in the second files no
messages are produced and the return code is zero, otherwise a non zero return
code is produced along with messages on stderr.

## Installing 

go get github.com/x0ray/lscan 

## Method

This program scans a Go language file, referred to as the index file. 
Scanning creates a token stream. The token stream is then examined and the literal 
string tokens are matched with a regular expression. If the token matches
its value is added to a match index map.

If specified the scan file is also indexed in the same way as the index file.  
The index is then used to examine the occurrence of index values in the scan
file, and report if all index values do not appear in the scan file.

## Usage
``` 
Usage of lscan.exe:
  -c string
        Alias for -scanrx option (default "\\.[a-zA-Z0-9\\.]+")
  -debug
        Provide debugging trace output
  -h    Alias for -help option
  -help
        Display program help text
  -i string
        Alias for -indexrx option (default "\\A\\.[a-zA-Z0-9\\.]+")
  -index string
        Input Go source file to used to create literal index
  -indexrx string
        Regular expression used to identify index file literals (default "\\A\\.[a-zA-Z0-9\\.]+")
  -list
        List the contents of the index and scan index
  -s string
        Alias for -scan option
  -scan string
        File to scan for occurrences of index values
  -scanrx string
        Regular expression used to identify scan file literals (default "\\.[a-zA-Z0-9\\.]+")
  -v    Alias for -version option
  -verbose
        Display progress information on the log
  -version
        Display program version information
  -x string
        Alias for -index option
```

## Examples

### Version

Display the program version using `lscan -v` or `lscan --version`

``` sh
> lscan --version
Program LScan version 0.0.10 as of 16Aug2021
```

### Help

Display lscan help information using `lscan --help` or `lscan -h`

``` sh
> lscan --help
Program LScan

        This command compares selected literal strings between two Go language source
        code files. The literals are selected using a regular expression pattern. If
        the literal strings from the first file are all found in the second files no
        messages are produced and the return code is zero, otherwise a non zero return
        code is produced along with messages on stderr.

Usage of D:\usr\GoApp\gopath\bin\lscan.exe:
  -c string
        Alias for -scanrx option (default "\\.[a-zA-Z0-9\\.]+")
  -debug
        Provide debugging trace output
  -h    Alias for -help option
  -help
        Display program help text
  -i string
        Alias for -indexrx option (default "\\A\\.[a-zA-Z0-9\\.]+")
  -index string
        Input Go source file to used to create literal index
  -indexrx string
        Regular expression used to identify index file literals (default "\\A\\.[a-zA-Z0-9\\.]+")
  -list
        List the contents of the index and scan index
  -s string
        Alias for -scan option
  -scan string
        File to scan for occurrences of index values
  -scanrx string
        Regular expression used to identify scan file literals (default "\\.[a-zA-Z0-9\\.]+")
  -v    Alias for -version option
  -verbose
        Display progress information on the log
  -version
        Display program version information
  -x string
        Alias for -index option 
```

### Index a file

Create an index of a Go language file and list the contents using `lscan -x index.go --list`. Input Go language file:
``` go// main entry to lscan a go language literal scanner utility
package main

// Some easy to fine test constants so this source file can be used as test input :-)
const (
        BundleID = "log"

        TestExpectedStatusError          = BundleID + ".log.expected.status.error.log"
        TestInitialIDNotValid2           = BundleID + ".log.initial.id.not.valid.log"
        TestUnmarshalFailedError         = BundleID + ".log.unmarshal.failed.error.log"
        TestClientExtractFailed          = BundleID + ".log.client.extract.failed.log"
        TestFileUriShortError            = BundleID + ".log.file.uri.short.error.log"
        TestUriPrefixNotValid            = BundleID + ".log.uri.prefix.not.valid.log"
        TestInvalidDependencyID          = BundleID + ".log.invalid.dependency.ID.log"
        TestFlowNameNotValid             = BundleID + ".log.name.not.valid.log"
        TestInitialIDNotValid            = BundleID + ".log.initial.id.not.valid.log"
        TestIDNotValidType               = BundleID + ".log.id.not.valid.type.log"
        TestIDNotValidNumberDependencies = BundleID + ".log.id.not.valid.number.dependencies.log"
)
```

``` sh
> lscan --index test_keys.go -list
2021/08/16 21:55:45 WARN Key: '.log.initial.id.not.valid.log' Loc: test_keys.go:16:48 already indexed. Not adding new key Loc: test_keys.go:9:48
2021/08/16 21:55:45 Index file keys from: 'test_keys.go' follow
Num     Location        Matched Literal
0001    test_keys.go:6:13       log
0002    test_keys.go:8:48       .log.expected.status.error.log
0003    test_keys.go:13:48      .log.uri.prefix.not.valid.log
0004    test_keys.go:15:48      .log.name.not.valid.log
0005    test_keys.go:14:48      .log.invalid.dependency.ID.log
0006    test_keys.go:17:48      .log.id.not.valid.type.log
0007    test_keys.go:18:48      .log.id.not.valid.number.dependencies.log
0008    test_keys.go:9:48       .log.initial.id.not.valid.log
0009    test_keys.go:10:48      .log.unmarshal.failed.error.log
0010    test_keys.go:11:48      .log.client.extract.failed.log
0011    test_keys.go:12:48      .log.file.uri.short.error.log
```

lscan -x index.go -s scan.go
    
    lscan -x index.go -list -verbose
    lscan -x index.go -debug
    lscan -x index.go -list -verbose -debug
    lscan -x index.go -s scan.go -list -verbose -debug
