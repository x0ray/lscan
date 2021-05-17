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

    lscan -v
    lscan -h
    lscan -x index.go -s scan.go
    lscan -x index.go -list
    lscan -x index.go -list -verbose
    lscan -x index.go -debug
    lscan -x index.go -list -verbose -debug
    lscan -x index.go -s scan.go -list -verbose -debug
