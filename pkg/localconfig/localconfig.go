package localconfig

import (
	"log"
)

// Configurations exported
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	IsSecure     bool
	Debug        bool `yml:"DEBUG,required"`
}

type Configurations struct {
	Server      ServerConfigurations
	Database    DatabaseConfigurations
	SERVER_NAME string
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}
