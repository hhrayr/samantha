package configs

type SmtpConfigs struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func GetSmtpConfigs() *SmtpConfigs {
	return &SmtpConfigs{
		Username:    "autopartsrialto@gmail.com",
		Password:    "__123456",
		EmailServer: "smtp.gmail.com",
		Port:        587,
	}
}
