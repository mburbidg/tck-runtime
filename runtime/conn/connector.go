package conn

type Connector interface {
	// Authenticate authenticates the given principal identified by principalID using the specified password.
	// On success an authentication ID is returned.
	Authenticate(principalID, password string) (authID string, err error)

	// NewSession creates a new GQL-session, for the authenticated principal identified by authID. On success
	// a session ID is returned.
	NewSession(authID string) (sessionID string, err error)

	// Request submits a GQL-request to the server. The returned outcome is a byte array containing a JSON
	// representation of the outcome. See the README for format details of the JSON.
	Request(authID, sessionID, program string) (outcome []byte, err error)
}
