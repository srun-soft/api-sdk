package sdk

type SDK interface {
	GetToken() (string, error)
}

type APIClient struct {
}
