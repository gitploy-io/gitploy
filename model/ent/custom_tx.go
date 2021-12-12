package ent

func (tx *Tx) GetDriverDialect() string {
	return tx.config.driver.Dialect()
}
