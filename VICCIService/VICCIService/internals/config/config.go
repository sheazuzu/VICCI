package config

type VICCIConfig struct {
	Basic_url string
	User string
	Password string

}

func LoadVICCIConfig(folder string) (*VICCIConfig, error){


	if err := loadEnv(folder); err != nil {
	return nil, err
	}

	return &VICCIConfig{
		getEnv("VICCI_BASIC_URL","default"),
		getEnv("AUTH_USER","default"),
		getEnv("AUTH_PASSWORD","default"),
	}, nil

}