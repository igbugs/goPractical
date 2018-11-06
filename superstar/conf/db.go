package conf

const DriverName = "mysql"

type DBConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

var MasterDbConfig = DBConf{
	User:     "gouser",
	Password: "123456",
	Host:     "192.168.247.133",
	Port:     3306,
	DBName:   "superstar",
}

var SlaveDbConfig = DBConf{
	User:     "gouser",
	Password: "123456",
	Host:     "192.168.247.133",
	Port:     3306,
	DBName:   "superstar",
}
