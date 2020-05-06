package dbl_go

// User represents a user on top.gg
type User struct {
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	ID            string `json:"id"`
	Username      string `json:"username"`
	DefAvatar     string `json:"defAvatar"`
	Banner        string `json:"banner"`
	Bio           string `json:"bio"`
	Color         string `json:"color"`
	Admin         bool   `json:"admin"`
	WebMod        bool   `json:"webMod"`
	Mod           bool   `json:"mod"`
	CertifiedDev  bool   `json:"certifiedDev"`
	Supporter     bool   `json:"supporter"`
	Social        Social `json:"social"`
}

// Social contains a users social connections on top.gg
type Social struct {
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Reddit    string `json:"reddit"`
	Twitter   string `json:"twitter"`
	Youtube   string `json:"youtube"`
}

// Get user gets a User for the given id from top.gg
func (c *DBL) GetUser(id string) (*User, error) {
	req, err := c.newRequest("GET", "/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	_, err = c.do(req, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
