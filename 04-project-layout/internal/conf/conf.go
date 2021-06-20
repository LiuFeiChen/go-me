package conf

type Grpc struct {
	Addr string
	Timeout string
}

type Server struct {
	Grpc  *Grpc
}

type Database struct  {
	Driver string
	Source string
}

type Data struct {
	Database *Database
}

type Config struct {
	Server *Server
	Data *Data
}
