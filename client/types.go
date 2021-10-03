package client

// A SessionIdentificationResponse object allows your client to identify the session it's using and prove that it is the owner of the session,
// it proves as an identification effort and security effort.
// Most methods that requires authentication or some sort of identity will require you to pass on this object as a parameter under session_identification,
// missing parameters or invalid security values will cause the request to fail as it's validated upon request.
type SessionIdentificationResponse struct {
	// The ID of the session obtained when establishing a session
	SessionId string `json:"session_id"`
	// The Public Hash of the client used when establishing the session
	ClientPublicHash string `json:"client_public_hash"`
	// The session challenge answer revolving around the client's private hash, the same client used to establish the session
	ChallengeAnswer string `json:"challenge_answer"`
}

// SessionRegisterResponse (session.register) Registers a new user to the network.
// This method is intended to be simple (for now), this will be changed in the future to add additional security like requiring the user to complete a captcha before proceeding.
// The client must present the network's terms of service document and ask if the user agrees to the terms of service.
type SessionRegisterResponse struct {
	// The Session Identification object.
	SessionIdentification SessionIdentificationResponse `json:"session_identification"`
	// The ID of the HelpDocument for the Terms of Service.
	TermsOfServiceId string `json:"terms_of_service_id"`
	// Indicates if the user has agreed to the terms of service, if this parameter is false then the method will throw an error.
	TermsOfServiceAgree bool `json:"terms_of_service_agree"`
	// The username to register to the network
	Username string `json:"username"`
	// The password used to authenticate to the network
	Password string `json:"password"`
	// The first name of the user to set as a display name
	FirstName string `json:"first_name"`
	// The last name of the user to set as a display name (optional)
	LastName string `json:"last_name"`
}

// CreateSessionResponse Establishes a new session to the network, new and unauthenticated sessions expires after 10 minutes of inactivity,
// authenticating to the session will increase the expiration time to 72 hours of inactivity.
// This timer is reset whenever the session is used in one way or another.
type CreateSessionResponse struct {
	// The Public Hash of the client that's establishing the session
	PublicHash string `json:"public_hash"`
	// The Private Hash of the client used when establishing the session
	PrivateHash string `json:"private_hash"`
	// The platform that the client is running on, eg; Linux, Windows, Android
	Platform string `json:"platform"`
	// The name of the client
	Name string `json:"name"`
	// The version of the client
	Version string `json:"version"`
}

// GetSessionResponse The Session Identification object
type GetSessionResponse SessionIdentificationResponse
