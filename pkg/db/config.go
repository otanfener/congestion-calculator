package db

type Config struct {
	URI        string `required:"true"`
	Name       string `required:"true"`
	Collection string `required:"true"`
}
