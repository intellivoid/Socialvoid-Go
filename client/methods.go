package client

import (
	"crypto/sha1"
	"fmt"
	"github.com/pquerna/otp/totp"
	"time"
)

func NewSession(publicHash string, privateHash string) (c *CreateSessionStruct) {
	c.PrivateHash = privateHash
	c.PublicHash = publicHash
	c.Name = VERSION
	c.Platform = ""
	return c
}

func (c *CreateSessionStruct) GetChallengeAnswer(challenge string) (string, error) {
	totpCode, err := totp.GenerateCode(challenge, time.Now())
	if err != nil {
		return "", nil
	}

	hash := sha1.New()
	hash.Write([]byte(totpCode + c.PrivateHash))
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
