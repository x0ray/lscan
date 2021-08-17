// main entry to lscan a go language literal scanner utility
package main

// Some easy to fine test constants so this source file can be used as test input :-)
const TestKeys = `
	base.test.expected.status.error.log=Expected status: {expect}, got status: {status}, for "{uri}" error code: {errorCode} {message}
	base.test.unmarshal.failed.error.log=JSON unmarshal for "{uri}" failed with error {error}
	base.test.client.extract.failed.log=Failed to extract rest client from context.
	base.test.uri.prefix.not.valid.log=File URI prefix "{fup}" is not "{prefix}"
	base.test.invalid.dependency.ID.log=Invalid dependency ID "{id}"
	base.test.name.notvalid.log=Invalid test name "{name}" is not a valid name.
	base.test.initial.id.not.valid.log=Initial test ID "{id}" not valid GUID format.
	base.test.id.not.valid.type.log=Flow ID "{id}" has invalid node type "{type}".
	base.test.id.not.valid.number.dependencies.log=Flow ID {id} has invalid number of dependencies {numdep}".
`
