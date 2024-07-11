// Package auth provides authentication utilities and moddleware for the server.
package auth

// IAPIKey represents an API key in the service.
type IAPIKey interface {
	Token() string
	Role() string
	AccountId() uint64
	CreatedTime() int64
}

// IAPIKeyPool represents a pool of API keys used to track and fetch tokens.
type IAPIKeyPool interface {
	GetAPIKey() *IAPIKey
}

// KeyPool is a pool of API keys used to track and fetch tokens.
type KeyPool struct {
	APIKeys    map[string]IAPIKey
	TotalCount int
}

// GetAPIKey gets an API key from the key pool.
func (k KeyPool) GetAPIKey(apiSecretToken string) *IAPIKey {
	if key, ok := k.APIKeys[apiSecretToken]; ok {
		return &key
	}
	return nil
}

// BaseAPIKey is the standard API key in system and implements the IAPIKey interface.
type BaseAPIKey struct {
	token       string
	role        string
	accountId   uint64
	createdTime int64 // Unix timestamp
}

// Token returns the API key token.
func (t BaseAPIKey) Token() string {
	return t.token
}

// Role returns the API key role.
func (t BaseAPIKey) Role() string {
	return t.role
}

// AccountId returns the API keys associated accounts accountId.
func (t BaseAPIKey) AccountId() uint64 {
	return t.accountId
}

// CreatedTime returns a UNIX timestamp of when API key was created.
func (t BaseAPIKey) CreatedTime() int64 {
	return t.createdTime
}
