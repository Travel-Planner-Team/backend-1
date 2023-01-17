package backend

var (
	MySQL *MySQLBackend
)

type MySQLBackend struct {
	client *mysql.Client
}

func InitMySQLBackend(config *util.MySQLInfo) {
	// some implementation here
}

func (backend *MySQLBackend) someFunc() error {
	//some implementation here
	return nil
}
