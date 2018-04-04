package config

//MollyConfig global var
var MollyConfig = DefaultConfiguration()

//Configuration structure
type Configuration struct {
	//Server configuration
	Server Server
	//StorageList configuration
	StorageList []Storage
}

//Server structure
type Server struct {
	//Address structure
	Address  string `json:"address"`
	Graphiql string `json:"graphiql"`
}

//Storage structure
type Storage struct {
	//Name definition
	Name string `json:"name"`
	//Path definition
	Path string `json:"localPath"`
}

//DefaultConfiguration creates a default instance
func DefaultConfiguration() *Configuration {
	return &Configuration{
		Server: Server{
			Address:  "0.0.0.0:9090",
			Graphiql: "resources/graphiql",
		},
		StorageList: []Storage{
			{
				Name: "i18n",
				Path: "./resources/data/i18n",
			},
			{
				Name: "db",
				Path: "./resources/data/databases",
			},
			{
				Name: "ms",
				Path: "./resources/data/microservices",
			},
		},
	}
}
