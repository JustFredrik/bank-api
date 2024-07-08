package auth

type IAPIKey interface {
	Token() string
	Role() string
	AccountId() uint64
	CreatedTime() int64
}

type IAPIKeyPool interface {
	GetAPIKey() *IAPIKey
}

type KeyPool struct {
	APIKeys    map[string]IAPIKey
	TotalCount int
}

func (k KeyPool) GetAPIKey(apiSecretToken string) *IAPIKey {
	if key, ok := k.APIKeys[apiSecretToken]; ok {
		return &key
	}
	return nil
}

type BaseAPIKey struct {
	token       string
	role        string
	accountId   uint64
	createdTime int64 // Unix timestamp
}

func (t BaseAPIKey) Token() string {
	return t.token
}

func (t BaseAPIKey) Role() string {
	return t.role
}

func (t BaseAPIKey) AccountId() uint64 {
	return t.accountId
}

func (t BaseAPIKey) CreatedTime() int64 {
	return t.createdTime
}
