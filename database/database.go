package database

type Service struct {
	Name          string
	DisplayName   string
	Executable    string
	Args          []string
	Description   string
	AppDirectory  string
	RestartPolicy string
}

type Config struct {
	LogLevel       string
	ServiceTimeout int
}
