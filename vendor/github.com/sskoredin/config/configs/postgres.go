package configs

import consul "github.com/sskoredin/config"

type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DB       string `toml:"db"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func (c *Postgres) Read() error {
	client, err := consul.NewClient()
	if err != nil {
		return err
	}

	c.Host, err = client.Get(consul.PostgresHost)
	if err != nil {
		return err
	}
	c.Port, err = client.GetInt(consul.PostgresPort)
	if err != nil {
		return err
	}
	c.DB, err = client.Get(consul.PostgresDB)
	if err != nil {
		return err
	}
	c.User, err = client.Get(consul.PostgresUSER)
	if err != nil {
		return err
	}
	c.Password, err = client.Get(consul.PostgresPASSWORD)
	if err != nil {
		return err
	}

	return nil
}
