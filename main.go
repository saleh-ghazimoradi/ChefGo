package main

import "github.com/saleh-ghazimoradi/ChefGo/cmd"

func main() {
	//client, err := utils.ConnectMongoDB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	cmd.Execute()
}
