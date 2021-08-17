// main entry to lscan a go language literal scanner utility
package main

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
