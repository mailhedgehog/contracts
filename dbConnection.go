package contracts

type DbConnectionConfig map[string]interface{}

type DbConnectionsConfig struct {
	Connections map[string]DbConnectionConfig `yaml:"connections"`
}
