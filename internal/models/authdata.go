package models

// AuthData is a struct that contains the authentication data for some
// entity within the system.  This allows us to over time support
// things like social login without needing to break everyone's
// password.
type AuthData struct {
	// We point to a user here to allow more than one form of
	// authentication to point to the same user.  This allows for
	// things like offline authentication as well as federated
	// cloud logins.
	User User 
	
	// SecretHash is the most basic authenticator, it is used to
	// store the hashed version of a secret which can be presented
	// for authentication purposes.
	SecretHash string

	// Capabilities are things that allow users to do things.
	// Users can have capabilities directly, but the more
	// appropriate way to grant them is via groups.
	Capabilities []string
}
