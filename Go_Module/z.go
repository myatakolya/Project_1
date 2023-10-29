package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Подключение к базе данных
	// Устанавливаем URI базы данных
	uri := "mongodb+srv://myatakolyajr:0968715133@cluster0.ebfmjow.mongodb.net/?retryWrites=true&w=majority"

	// Создаем клиент MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string("Dot"))
	// Подключаемся к базе данных
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Подключение к базе данных MongoDB успешно установлено")

	// Открываем коллекцию "users"
	collection := client.Database("proj").Collection("shedule")

	var results []bson.M
	cur, err := collection.Find(ctx, options.Find())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Выводим результаты
	for _, result := range results {
		fmt.Printf("%#v\n", result)
	}
}
