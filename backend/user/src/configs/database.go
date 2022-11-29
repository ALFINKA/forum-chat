package configs

import (
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	//err := godotenv.Load("../../.env")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(os.Getenv("DB_HOST"))
	//	log.Fatalln("error loading .env file")
	//}

	//host := os.Getenv("DB_HOST") || "localhost"
	//user := os.Getenv("DB_USER") || "root"
	//password := os.Getenv("DB_PASSWORD") ||| "root"
	//database := os.Getenv("DB_NAME") || "group_chat"
	//port := os.Getenv("DB_PORT") || 9905
	//host := "localhost"
	//user := "root"
	//password := "root"
	//database := "group_chat"
	//port := 9905

	//dsn := "host=" + host + "user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"
	dsn := "host=localhost user=postgres password=root dbname=group_chat port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		//fmt.Println(err)
		log.Fatalln(" database connection failed")
	}
	return db
}
