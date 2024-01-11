package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Exams struct {
	ID       primitive.ObjectID `bson:"_id"`
	Subject  string             `bson:"subject"`
	Teacher  string             `bson:"teacher"`
	Time     string             `bson:"time"`
	Room     string             `bson:"room"`
	Date     string             `bson:"date"`
	Group    string             `bson:"group"`
	Subgroup string             `bson:"subgroup"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tg_id    string             `bson:"tg_id"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Username string             `bson:"username"`
	Group    string             `bson:"group,omitempty"`
	Subgroup string             `bson:"subgroup"`
	Role     string             `bson:"role"`
}

type Monday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

type Tuesday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
	Lesson2 Teacher_Lesson `bson:"lesson2"`
}

type Wednesday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

type Thursday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

type Friday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
	Lesson2 Teacher_Lesson `bson:"lesson2"`
	Lesson3 Teacher_Lesson `bson:"lesson3"`
	Lesson4 Teacher_Lesson `bson:"lesson4"`
}

type Teacher_Lesson struct {
	Message string `bson:"message,omitempty"`
	Room    string `bson:"room,omitempty"`
	Time    string `bson:"time,omitempty"`
	Type    string `bson:"type,omitempty"`
	Group   string `bson:"group"`
}

type Teachers_Lessons struct {
	Monday    Monday    `bson:"monday,omitempty"`
	Tuesday   Tuesday   `bson:"tuesday,omitempty"`
	Wednesday Wednesday `bson:"wednesday,omitempty"`
	Thursday  Thursday  `bson:"thursday"`
	Friday    Friday    `bson:"friday,omitempty"`
}

type Teachers_Schedule struct {
	ID      primitive.ObjectID `bson:"_id"`
	Lessons Teachers_Lessons   `bson:"lessons"`
	Name    string             `bson:"name"`
	Surname string             `bson:"surname"`
	Subject string             `bson:"subject"`
}

type Teacher struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tg_id    string             `bson:"tg_id"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Username string             `bson:"username"`
	Role     string             `bson:"role"`
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
	Lesson5 Lesson `bson:"lesson5,omitempty"`
}

type Schedule struct {
	Day      string  `bson:"day"`
	Group    string  `bson:"group"`
	Lessons  Lessons `bson:"lessons"`
	Message1 string  `bson:"message1"`
}

type Connection struct {
	Schedule          *mongo.Collection
	Users             *mongo.Collection
	Exams             *mongo.Collection
	Teachers_Schedule *mongo.Collection
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

func (connection Connection) GetScheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Schedule
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Schedule.Find(ctx, bson.M{})
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

func (connection Connection) GetTeachersScheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Teachers_Schedule
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Teachers_Schedule.Find(ctx, bson.M{})
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

func (connection Connection) GetScheduleGroupEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Schedule
	params := mux.Vars(request)
	group, _ := params["group"]
	json.NewDecoder(request.Body).Decode(&group)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Schedule.Find(ctx, bson.M{"group": group})
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

func (connection Connection) GetScheduleDayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Schedule
	params := mux.Vars(request)
	day, _ := params["day"]
	json.NewDecoder(request.Body).Decode(&day)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Schedule.Find(ctx, bson.M{"day": day})
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

func (connection Connection) GetScheduleGroupDayEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Schedule
	params := mux.Vars(request)
	group, _ := params["group"]
	day, _ := params["day"]
	json.NewDecoder(request.Body).Decode(&group)
	json.NewDecoder(request.Body).Decode(&day)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Schedule.Find(ctx, bson.M{"group": group, "day": day})
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

func (connection Connection) GetExamsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var exams []Exams
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Exams.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &exams); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(exams)
}

func (connection Connection) GetExamsGroupEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var exams []Exams
	params := mux.Vars(request)
	group, _ := params["group"]
	json.NewDecoder(request.Body).Decode(&group)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Exams.Find(ctx, bson.M{"group": group})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &exams); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(exams)
}

func (connection Connection) GetExamsGroupSubGroupEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var exams []Exams
	params := mux.Vars(request)
	group, _ := params["group"]
	subgroup, _ := params["subgroup"]
	json.NewDecoder(request.Body).Decode(&group)
	json.NewDecoder(request.Body).Decode(&subgroup)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := connection.Exams.Find(ctx, bson.M{"group": group, "subgroup": subgroup})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &exams); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(exams)
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
	result, err := connection.Schedule.UpdateOne(ctx, bson.M{"day": day}, bson.D{{"$set", bson.D{{messageN, message}}}})
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
	opts := options.Client().ApplyURI("mongodb+srv://myatak:0968715133@cluster0.0zmnueu.mongodb.net/").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	connection := Connection{
		Schedule:          client.Database("project").Collection("schedule"),
		Users:             client.Database("project").Collection("users"),
		Exams:             client.Database("project").Collection("exams"),
		Teachers_Schedule: client.Database("project").Collection("teachers_schedule"),
	}

	router := mux.NewRouter()

	router.HandleFunc("/user", connection.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/teacher", connection.CreateTeacherEndpoint).Methods("POST")

	router.HandleFunc("/users", connection.GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/teachers", connection.GetTeachersEndpoint).Methods("GET")
	router.HandleFunc("/admins", connection.GetAdminsEndpoint).Methods("GET")
	router.HandleFunc("/students", connection.GetStudentsEndpoint).Methods("GET")

	router.HandleFunc("/schedule", connection.GetScheduleEndpoint).Methods("GET")
	router.HandleFunc("/schedule/{group}", connection.GetScheduleGroupEndpoint).Methods("GET")
	router.HandleFunc("/scheduled/{day}", connection.GetScheduleDayEndpoint).Methods("GET")
	router.HandleFunc("/schedule/{group}/{day}", connection.GetScheduleGroupDayEndpoint).Methods("GET")

	router.HandleFunc("/teachers_schedule", connection.GetTeachersScheduleEndpoint).Methods("GET")

	router.HandleFunc("/exams", connection.GetExamsEndpoint).Methods("GET")
	router.HandleFunc("/exams/{group}", connection.GetExamsGroupEndpoint).Methods("GET")
	router.HandleFunc("/exams/{group}/{subgroup}", connection.GetExamsGroupSubGroupEndpoint).Methods("GET")

	router.HandleFunc("/user/{id}", connection.UpdateUserEndpoint).Methods("PUT")
	router.HandleFunc("/schedule/{day}/{lesson}/{message}", connection.UpdateMessageEndpoint).Methods("PUT")

	router.HandleFunc("/user/{id}", connection.DeleteUserEndpoint).Methods("DELETE")

	http.ListenAndServe("localhost:2000", router)
}
