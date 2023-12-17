package main

// configuration is the configuration for the service.
// for now, we are using a single configuration for the whole service.
// in the future, we can split the configuration and move to the base folder
type configuration struct {
	HostAndPort        string `yaml:"hostAndPort" env:"DB_HOST_AND_PORT,overwrite,default=localhost:5432"`
	Username           string `yaml:"username" env:"DB_USER,overwrite,default=bego"`
	Password           string `yaml:"password" env:"DB_PASSWORD,overwrite,default=password"`
	Database           string `yaml:"database" env:"DB_DATABASE,overwrite,default=bego"`
	SSLMode            string `yaml:"sslMode" env:"DB_SSLMODE,overwrite,default=disable"`
	StatementCacheMode string `yaml:"statementCacheMode" env:"DB_STATEMENT_CACHE_MODE,overwrite,default=prepare"`
	Migrate            bool   `yaml:"migrate" env:"DB_MIGRATE,overwrite,default=true"`
}
