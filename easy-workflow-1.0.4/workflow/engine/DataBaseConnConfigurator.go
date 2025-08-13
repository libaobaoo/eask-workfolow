package engine

// DataBaseConnConfigurator holds redis client configuration options.
type DataBaseConnConfigurator struct {
	Addr         string // redis server address
	Password     string // redis password
	DB           int    // redis database index
	PoolSize     int    // maximum number of socket connections
	MinIdleConns int    // minimum number of idle connections
}

// DBConnConfigurator provides default redis connection settings.
var DBConnConfigurator = DataBaseConnConfigurator{
	Addr:         "localhost:6379",
	DB:           0,
	PoolSize:     10,
	MinIdleConns: 2,
}
