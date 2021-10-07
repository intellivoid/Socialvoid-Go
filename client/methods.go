package client

func NewSession(publicHash string, privateHash string) (c *CreateSessionStruct) {
	c.PrivateHash = privateHash
	c.PublicHash = publicHash
	c.Name = VERSION
	c.Platform = ""
	return c
}

func (c *CreateSessionStruct) GetChallengeAnswer(secret string) (string, error) {
	// pvt_hash := c.PrivateHash
	return "", nil
}
