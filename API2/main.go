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

type Role struct {
	Naming   string `bson:"naming"`
	Sched    bool   `bson:"sched"`
	Full     bool   `bson:"full"`
	Comments bool   `bson:"comments"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tg_id    string             `bson:"tg_id"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Username string             `bson:"username"`
	Groupe   int                `bson:"groupe,omitempty"`
	Role     string             `bson:"role"`
}

type Teacher struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tg_id    string             `bson:"tg_id"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Username string             `bson:"username"`
	Role     string             `bson:"role"`
}

type Mess struct {
	Message1 string `bson:"message1"`
}

type Lesson struct {
	Subject string `bson:"subject,omitempty"`
	Teacher string `bson:"teacher,omitempty"`
	Room    string `bson:"room,omitempty"`
	Type    string `bson:"type,omitempty"`
	Time    string `bson:"time,omitempty"`
}

type Lessons struct {
	Lesson1 Lesson `bson:"lesson1,omitempty"`
	Lesson2 Lesson `bson:"lesson2,omitempty"`
	Lesson3 Lesson `bson:"lesson3,omitempty"`
	Lesson4 Lesson `bson:"lesson4,omitempty"`
}

type LessonsT struct {
	Lesson1 Lesson `bson:"lesson1,omitempty"`
	Lesson2 Lesson `bson:"lesson2,omitempty"`
	Lesson3 Lesson `bson:"lesson3,omitempty"`
}

type LessonsF struct {
	Lesson1 Lesson `bson:"lesson1,omitempty"`
}

type Monday struct {
	Day      string  `bson:"day"`
	Group    int     `bson:"group"`
	Lessons  Lessons `bson:"lessons"`
	Message1 string  `bson:"message1"`
	Message2 string  `bson:"message2"`
	Message3 string  `bson:"message3"`
	Message4 string  `bson:"message4"`
}

type Tuesday struct {
	Day      string   `bson:"day"`
	Group    int      `bson:"group"`
	Lessons  LessonsT `bson:"lessons"`
	Message1 string   `bson:"message1"`
	Message2 string   `bson:"message2"`
	Message3 string   `bson:"message3"`
}

type Friday struct {
	Day      string   `bson:"day"`
	Group    int      `bson:"group"`
	Lessons  LessonsF `bson:"lessons"`
	Message1 string   `bson:"message1"`
}

type Shedule struct {
	Day      string  `bson:"day"`
	Group    int     `bson:"group"`
	Lessons  Lessons `bson:"lessons"`
	Message1 string  `bson:"message1"`
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

func (connection Connection) GetMondayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var day []Monday
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{"day": "Понедельник"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &day); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(day)
}

func (connection Connection) GetTuesdayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var day []Tuesday
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{"day": "Вторник"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &day); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(day)
}

func (connection Connection) GetWednesdayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var day []Tuesday
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{"day": "Среда"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &day); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(day)
}

func (connection Connection) GetThursdayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var day []Tuesday
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{"day": "Четверг"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &day); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(day)
}

func (connection Connection) GetFridayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var day []Friday
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Shedule.Find(ctx, bson.M{"day": "Пятница"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &day); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(day)
}

func (connection Connection) GetAdminsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var admins []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Users.Find(ctx, bson.M{"role": "Админ"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &admins); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(admins)
}

func (connection Connection) GetStudentsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var admins []User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Users.Find(ctx, bson.M{"role": "Студент"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &admins); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(admins)
}

func (connection Connection) GetTeachersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var teachers []Teacher
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Users.Find(ctx, bson.M{"role": "Учитель"})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &teachers); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(teachers)
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

func (connection Connection) UpdateMessageEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	day, _ := params["day"]
	messageN, _ := params["lesson"]
	message, _ := params["message"]
	json.NewDecoder(request.Body).Decode(&message)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := connection.Shedule.UpdateOne(ctx, bson.M{"day": day}, bson.D{{"$set", bson.D{{messageN, message}}}})
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
	router.HandleFunc("/teacher", connection.CreateTeacherEndpoint).Methods("POST")

	router.HandleFunc("/users", connection.GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/teachers", connection.GetTeachersEndpoint).Methods("GET")
	router.HandleFunc("/admins", connection.GetAdminsEndpoint).Methods("GET")
	router.HandleFunc("/students", connection.GetStudentsEndpoint).Methods("GET")
	router.HandleFunc("/schedule", connection.GetSheduleEndpoint).Methods("GET")
	router.HandleFunc("/monday", connection.GetMondayEndpoint).Methods("GET")
	router.HandleFunc("/tuesday", connection.GetTuesdayEndpoint).Methods("GET")
	router.HandleFunc("/wednesday", connection.GetWednesdayEndpoint).Methods("GET")
	router.HandleFunc("/thursday", connection.GetThursdayEndpoint).Methods("GET")
	router.HandleFunc("/friday", connection.GetFridayEndpoint).Methods("GET")

	router.HandleFunc("/user/{id}", connection.UpdateUserEndpoint).Methods("PUT")
	router.HandleFunc("/shedule/{day}/{lesson}/{message}", connection.UpdateMessageEndpoint).Methods("PUT")

	router.HandleFunc("/user/{id}", connection.DeleteUserEndpoint).Methods("DELETE")

	http.ListenAndServe("localhost:2000", router)
}
