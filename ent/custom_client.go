package ent

func (c *Client) GetDriverDialect() string {
	return c.config.driver.Dialect()
}
