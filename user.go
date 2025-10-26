package cloudreve

import "net/http"

func (c *Client) Login(usr, pwd string) (*User, error) {
	// prepare payload
	payload := LoginRequest{
		Username: usr,
		Password: pwd,
	}

	// prepare receiving data
	var user User

	// call private do method
	// API path: /api/v3/user/session
	// HTTP method: post
	err := c.do(http.MethodPost, "/api/v3/user/session", payload, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
