package handlers

import(
	"errors"
	"os"
	"encoding/base64"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"fmt"

)

// CreateSessionID generate a new session token
func CreateSessionID(sessionKey string) (string, error)  {
	if len(sessionKey) < 1 {
		return "", errors.New("signing key must not be empty")
	}

	token := make([]byte, 0)

	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Printf("error generating salt: %v\n", err)
		os.Exit(1)
	}

	hash := hmac.New(sha256.New, []byte(sessionKey))
	hash.Write(salt)
	signature := hash.Sum(nil)

	for _, byte := range salt {
		token = append(token, byte)
	}

	for _, byte := range signature {
		token = append(token, byte)
	}

	res := base64.URLEncoding.EncodeToString(token)

	return res, nil
}

//StartSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func StartSession(sessionKey string, rs RedisStore, sessionState interface{}, w http.ResponseWriter) (string, error) {
	sid, err := CreateSessionID(sessionKey)
	if err != nil {
		return "", err
	}

	rs.Save(sid, sessionState)
	w.Header().Add("Authorization", sid)
	return sid, nil
}


//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(sessionKey string, rs RedisStore, sid string) (string, error) {
	
	_, err := ValidateID(sid, sessionKey)

	if err != nil {
		return "", errors.New("Invalid Session ID")
	}

	rs.Delete(sid)

	return sid, nil
}


//ValidateID validates the string in the `id` parameter
//using the `signingKey` as the HMAC signing key
//and returns an error if invalid, or a SessionID if valid
func ValidateID(sid string, sessionKey string) (string, error) {

	token, _ := base64.URLEncoding.DecodeString(sid)

	var salt []byte
	var signature []byte

	for i, byte := range token {
		if i < 32 {
			salt = append(salt, byte)
		} else {
			signature = append(signature, byte)
		}
	}

	//If they match,
	//return the entire `id` parameter as a SessionID type.
	hash := hmac.New(sha256.New, []byte(sessionKey))
	hash.Write(salt)
	hashSignature := hash.Sum(nil)

	if hmac.Equal(signature, hashSignature) {
		return sid, nil
	}

	//If not, return empty string  and new invalid id error.
	return "", errors.New("Invalid Session ID")
}
