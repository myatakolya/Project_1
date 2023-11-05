package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

/*{
"_id": {
"$oid": "65467941405d337d4cc2f9ca"
},
"day": "Пятница",
"group": 231,
"lessons": [
{
"subject" : "Английский",
"teacher" : "Фабрина А.В.",
"room" : "306_В",
"type" : "Практика",
"time" : "11:30-13:00",
"message": ""
}
],
"message": "229"
}

type Role struct {
	Naming        string `bson:"naming"`
	Sheduleaccess bool   `bson:"sheduleaccess"`
	Fullaccess    bool   `bson:"fullaccess"`
	Comments      bool   `bson:"comments"`
}*/

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	id         int                `bson:"id"`
	Name       string             `bson:"name"`
	Secondname string             `bson:"secondname"`
	Username   string             `bson:"username"`
	Groupe     int                `bson:"groupe,omitempty"`
	Roles      bson.A             `bson:"roles"`
}

type Teacher struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	id         int                `bson:"id"`
	Name       string             `bson:"name"`
	Secondname string             `bson:"secondname"`
	Username   string             `bson:"username"`
	Roles      bson.A             `bson:"roles"`
}

type Shedule struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Day              string             `bson:"day"`
	Groupe           int                `bson:"group"`
	Lessons          bson.M             `bson:"lessons"`
	Messagebyteacher bson.M             `bson:"messagebyteacher"`
}

type Connection struct {
	Shedule *mongo.Collection
	Users   *mongo.Collection
	Roles   *mongo.Collection
}

func (connection Connection) CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Users.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) CreateTeacherEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var teacher Teacher
	if err := json.NewDecoder(request.Body).Decode(&teacher); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Users.InsertOne(ctx, teacher)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Users.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func (connection Connection) GetSheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Shedule
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &shed); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(shed)
}

func (connection Connection) UpdateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Users.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", user}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) UpdateSheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var day Shedule
	json.NewDecoder(request.Body).Decode(&day)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Shedule.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", day}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) DeleteUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Users.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://myatak:0968715133@cluster0.0zmnueu.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	connection := Connection{
		Shedule: client.Database("project").Collection("shedule"),
		Users:   client.Database("project").Collection("users"),
		Roles:   client.Database("project").Collection("roles"),
	}

	router := mux.NewRouter()
	router.HandleFunc("/user", connection.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user", connection.CreateTeacherEndpoint).Methods("POST")
	router.HandleFunc("/users", connection.GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/shedule", connection.GetSheduleEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", connection.UpdateUserEndpoint).Methods("PUT")
	router.HandleFunc("/shedule/{id}", connection.UpdateSheduleEndpoint).Methods("PUT")
	router.HandleFunc("/user/{id}", connection.DeleteUserEndpoint).Methods("DELETE")
	http.ListenAndServe("localhost:8080", router)
}
