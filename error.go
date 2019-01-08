package main

const (
	// JSON Errors
	JSONUNMARSHAL = "Error Unmarshalling json"

	// JWY Auth Errors
	SCOPETYPEERROR     = "Error Parsing Token Scope"
	TOKENNOTVALID      = "The Token Given is not Valid"
	PERMISSIONNOTGIVEN = "User is not Authorized to Use This Route"

	// Generic Errors
	INTERNALSERVICEERROR = "Internal Service Error"
)
