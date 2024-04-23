package main

import "github.com/go-sql-driver/mysql"

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		panic(err)
	}

	store := NewStorage(db)
	api := NewAPIServer(":3000", store)
	api.ListenAndServe()
}
