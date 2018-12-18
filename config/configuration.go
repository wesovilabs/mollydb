package config

//MollyConfig global var
var MollyConfig = DefaultConfiguration()

//Configuration structure
type Configuration struct {
	//Server configuration
	Server Server
}

//Server structure
type Server struct {
	//Address structure
	Address  string `json:"address"`
	Graphiql string `json:"graphiql"`
}

//DefaultConfiguration creates a default instance
func DefaultConfiguration() *Configuration {
	return &Configuration{
		Server: Server{
			Address:  "0.0.0.0:9090",
			Graphiql: "resources/graphiql",
		},
	}
}
