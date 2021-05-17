// main entry to lscan a go language literal scanner utility
package main

// Some easy to fine test constants so this source file can be used as test input :-)
const TestKeys = `
	sonder-log-icu.flow.expected.status.error.log=Expected status: {expect}, got status: {status}, for "{uri}" error code: {errorCode} {message}
	sonder-log-icu.flow.unmarshal.failed.error.log=JSON unmarshal for "{uri}" failed with error {error}
	sonder-log-icu.flow.client.extract.failed.log=Failed to extract rest client from context.
	sonder-log-icu.flow.uri.prefix.not.valid.log=File URI prefix "{fup}" is not "{prefix}"
	sonder-log-icu.flow.invalid.dependency.ID.log=Invalid dependency ID "{id}"
	sonder-log-icu.flow.name.notvalid.log=Invalid flow name "{name}" is not a valid SAS name.
	sonder-log-icu.flow.initial.id.not.valid.log=Initial flow ID "{id}" not valid GUID format.
	sonder-log-icu.flow.id.not.valid.type.log=Flow ID "{id}" has invalid node type "{type}".
	sonder-log-icu.flow.id.not.valid.number.dependencies.log=Flow ID {id} has invalid number of dependencies {numdep}".
`
