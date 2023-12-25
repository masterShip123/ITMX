package utilities

import "os"

func SetEnv() {
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_USER", "root")
	_ = os.Setenv("DB_PASSWORD", "my-secret-pw")
	_ = os.Setenv("DB_DATABASE", "itmx")

}
