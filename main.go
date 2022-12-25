package go_product_api

import "os"

func main() {
	a := App{}
	a.Init(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	a.Run(":8010")
}
