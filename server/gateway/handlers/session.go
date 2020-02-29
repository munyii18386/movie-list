package handlers

import(
	"errors"
	"io"
	"encoding/base64"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha256"
	"net/http"

)

// CreateSessionID generate a new session token
func CreateSessionID(sessionKey string) (string, error)  {
	if len(sessionKey) < 1 {
		return "", errors.New("signing key must not be empty")
	}
	b := make([]byte, 32)
	 if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return "", errors.New("unable to generate session message")
    }
	h := hmac.New(sha256.New, []byte(sessionKey))
	h.Write(b)
	signature := h.Sum(nil)
	res := base64.URLEncoding.EncodeToString(signature)

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