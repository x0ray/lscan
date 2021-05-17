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
	BundleID = "sonder-log-icu"

	FlowOrchestratorExpectedStatusError          = BundleID + ".flow.expected.status.error.log"
	FlowOrchestratorInitialIDNotValid2           = BundleID + ".flow.initial.id.not.valid.log"
	FlowOrchestratorUnmarshalFailedError         = BundleID + ".flow.unmarshal.failed.error.log"
	FlowOrchestratorClientExtractFailed          = BundleID + ".flow.client.extract.failed.log"
	FlowOrchestratorFileUriShortError            = BundleID + ".flow.file.uri.short.error.log"
	FlowOrchestratorUriPrefixNotValid            = BundleID + ".flow.uri.prefix.not.valid.log"
	FlowOrchestratorInvalidDependencyID          = BundleID + ".flow.invalid.dependency.ID.log"
	FlowOrchestratorFlowNameNotValid             = BundleID + ".flow.name.not.valid.log"
	FlowOrchestratorInitialIDNotValid            = BundleID + ".flow.initial.id.not.valid.log"
	FlowOrchestratorIDNotValidType               = BundleID + ".flow.id.not.valid.type.log"
	FlowOrchestratorIDNotValidNumberDependencies = BundleID + ".flow.id.not.valid.number.dependencies.log"
)
`

var scanThis string = `// main entry to lscan a go language literal scanner utility
package main

// Some easy to fine test constants so this source file can be used as test input :-)
const TestKeys = "
	sonder-log-icu.flow.expected.status.error.log=Expected status: {expect}, got status: {status}, for \"{uri}\" error code: {errorCode} {message}
	sonder-log-icu.flow.unmarshal.failed.error.log=JSON unmarshal for \"{uri}\" failed with error {error}
	sonder-log-icu.flow.client.extract.failed.log=Failed to extract rest client from context.
	sonder-log-icu.flow.uri.prefix.not.valid.log=File URI prefix \"{fup}\" is not \"{prefix}\"
	sonder-log-icu.flow.invalid.dependency.ID.log=Invalid dependency ID \"{id}\"
	sonder-log-icu.flow.name.notvalid.log=Invalid flow name \"{name}\" is not a valid SAS name.
	sonder-log-icu.flow.initial.id.not.valid.log=Initial flow ID \"{id}\" not valid GUID format.
	sonder-log-icu.flow.id.not.valid.type.log=Flow ID \"{id}\" has invalid node type \"{type}\".
	sonder-log-icu.flow.id.not.valid.number.dependencies.log=Flow ID {id} has invalid number of dependencies {numdep}.
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
