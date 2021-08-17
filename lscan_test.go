// main entry to lscan a Go language source file literal scanner utility
package main

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/x0ray/lscan/test"
)

var indexThis string = `
// Some easy to fine test constants so this source file can be used as test input :-)
const (
	BundleID = "log"

	TestExpectedStatusError          = BundleID + ".test.expected.status.error.log"
	TestInitialIDNotValid2           = BundleID + ".test.initial.id.not.valid.log"
	TestUnmarshalFailedError         = BundleID + ".test.unmarshal.failed.error.log"
	TestClientExtractFailed          = BundleID + ".test.client.extract.failed.log"
	TestFileUriShortError            = BundleID + ".test.file.uri.short.error.log"
	TestUriPrefixNotValid            = BundleID + ".test.uri.prefix.not.valid.log"
	TestInvalidDependencyID          = BundleID + ".test.invalid.dependency.ID.log"
	TestFlowNameNotValid             = BundleID + ".test.name.not.valid.log"
	TestInitialIDNotValid            = BundleID + ".test.initial.id.not.valid.log"
	TestIDNotValidType               = BundleID + ".test.id.not.valid.type.log"
	TestIDNotValidNumberDependencies = BundleID + ".test.id.not.valid.number.dependencies.log"
)
`

var scanThis string = `// main entry to lscan a go language literal scanner utility
package main

// Some easy to fine test constants so this source file can be used as test input :-)
const TestKeys = "
	test.expected.status.error.log=Expected status: {expect}, got status: {status}, for \"{uri}\" error code: {errorCode} {message}
	test.unmarshal.failed.error.log=JSON unmarshal for \"{uri}\" failed with error {error}
	test.client.extract.failed.log=Failed to extract rest client from context.
	test.uri.prefix.not.valid.log=File URI prefix \"{fup}\" is not \"{prefix}\"
	test.invalid.dependency.ID.log=Invalid dependency ID \"{id}\"
	test.name.notvalid.log=Invalid test name \"{name}\" is not a valid name.
	test.initial.id.not.valid.log=Initial test ID \"{id}\" not valid GUID format.
	test.id.not.valid.type.log=Flow ID \"{id}\" has invalid node type \"{type}\".
	test.id.not.valid.number.dependencies.log=Flow ID {id} has invalid number of dependencies {numdep}.
"
`

func Test_main(t *testing.T) {
	var ttMutex = &sync.RWMutex{}
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "LscanVersion",
			args: []string{"-v"},
			want: 0,
		},
		{
			name: "LscanHelp",
			args: []string{"-h"},
			want: 0,
		},
		{
			name: "LscanMsgKeyScan",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-s", test.MakeFile("scan", "go", scanThis)},
			want: 2,
		},
		{
			name: "LscanMsgKeyIndexList",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-list"},
			want: 1,
		},
		{
			name: "LscanMsgKeyIndexListVerbose",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-list", "-verbose"},
			want: 1,
		},
		{
			name: "LscanMsgKeyIndexListDebug",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-debug"},
			want: 1,
		},
		{
			name: "LscanMsgKeyIndexListVerboseDebug",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-list", "-verbose", "-debug"},
			want: 1,
		},
		{
			name: "LscanMsgKeys",
			args: []string{"-x", test.MakeFile("index", "go", indexThis), "-s", test.MakeFile("scan", "go", scanThis), "-list", "-verbose", "-debug"},
			want: 2,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			ttMutex.RLock()
			fmt.Printf("Start: %s\n", tt.name)
			os.Args = os.Args[:1]
			for _, v := range tt.args {
				os.Args = append(os.Args, v)
			}
			fmt.Printf("Args: %v\n", os.Args)
			got := body()
			fmt.Printf("Stop: %s RC: %d\n", tt.name, got)
			if got != tt.want {
				t.Errorf("%s body() got: %v want %v\n", tt.name, got, tt.want)
			}
			ttMutex.RUnlock()
		})

	}
}
