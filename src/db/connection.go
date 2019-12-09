package db

// import (
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// /*Create mysql connection*/
// func CreateCon() *mongo.Client {

// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("db is connected")
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)

// 	return client
// }
