package main

const (
	// JSON Errors
	JSONUNMARSHAL = "Error Unmarshalling json"

	// JWY Auth Errors
	SCOPETYPEERROR     = "Error Parsing Token Scope"
	TOKENNOTVALID      = "The Token Given is not Valid"
	PERMISSIONNOTGIVEN = "User is not Authorized to Use This Route"

	// Credential Auth Errors
	INVALIDCREDENTIALS     = "Missing Credentials"
	INVALIDEMAILORPASSWORD = "Invalid Email Address or Password"

	// Generic Errors
	INTERNALSERVICEERROR = "Internal Service Error"

	// DB Errors
	DBCONNECTIONERROR = "Error connecting to the database"

)
