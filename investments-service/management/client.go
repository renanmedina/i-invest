package management

type Client struct {
	Id    string
	Name  string
	Email string
}

func NewClient(id string, name string, email string) Client {
	return Client{id, name, email}
}

func (c *Client) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":    c.Id,
		"name":  c.Name,
		"email": c.Email,
	}
}
