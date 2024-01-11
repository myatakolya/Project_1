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

// Структура повторяет структуру базы данных, для корректной работы с данными
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

// Структура повторяет структуру базы данных, для корректной работы с данными
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

// Структура повторяет структуру базы данных, для корректной работы с данными
type Monday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Tuesday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
	Lesson2 Teacher_Lesson `bson:"lesson2"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Wednesday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Thursday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Friday struct {
	Lesson1 Teacher_Lesson `bson:"lesson1"`
	Lesson2 Teacher_Lesson `bson:"lesson2"`
	Lesson3 Teacher_Lesson `bson:"lesson3"`
	Lesson4 Teacher_Lesson `bson:"lesson4"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Teacher_Lesson struct {
	Message string `bson:"message,omitempty"`
	Room    string `bson:"room,omitempty"`
	Time    string `bson:"time,omitempty"`
	Type    string `bson:"type,omitempty"`
	Group   string `bson:"group"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Teachers_Lessons struct {
	Monday    Monday    `bson:"monday,omitempty"`
	Tuesday   Tuesday   `bson:"tuesday,omitempty"`
	Wednesday Wednesday `bson:"wednesday,omitempty"`
	Thursday  Thursday  `bson:"thursday"`
	Friday    Friday    `bson:"friday,omitempty"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Teachers_Shedule struct {
	ID      primitive.ObjectID `bson:"_id"`
	Lessons Teachers_Lessons   `bson:"lessons"`
	Name    string             `bson:"name"`
	Surname string             `bson:"surname"`
	Subject string             `bson:"subject"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Teacher struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tg_id    string             `bson:"tg_id"`
	Name     string             `bson:"name"`
	Surname  string             `bson:"surname"`
	Username string             `bson:"username"`
	Role     string             `bson:"role"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Lesson struct {
	Subject string `bson:"subject,omitempty"`
	Teacher string `bson:"teacher,omitempty"`
	Room    string `bson:"room,omitempty"`
	Type    string `bson:"type,omitempty"`
	Time    string `bson:"time,omitempty"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Lessons struct {
	Lesson1 Lesson `bson:"lesson1,omitempty"`
	Lesson2 Lesson `bson:"lesson2,omitempty"`
	Lesson3 Lesson `bson:"lesson3,omitempty"`
	Lesson4 Lesson `bson:"lesson4,omitempty"`
	Lesson5 Lesson `bson:"lesson5,omitempty"`
}

// Структура повторяет структуру базы данных, для корректной работы с данными
type Shedule struct {
	Day      string  `bson:"day"`
	Group    string  `bson:"group"`
	Lessons  Lessons `bson:"lessons"`
	Message1 string  `bson:"message1"`
}

// Структура для подключения к коллекциям в базе данных
type Connection struct {
	Shedule           *mongo.Collection
	Users             *mongo.Collection
	Exams             *mongo.Collection
	Teachers_Schedule *mongo.Collection
}

// Функция которая обрабатывает метод POST для добавления пользователя в базу данных
func (connection Connection) CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User                                                       //создаётся объект по типу структуры User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil { //обработка ошибок
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // объявляем контекс
	result, err := connection.Users.InsertOne(ctx, user)                //выполняем действие
	if err != nil {                                                     //обработка ошибок
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result) //влзвращаем результат
}

// Функция которая обрабатывает метод POST для добавления учителя в базу данных
func (connection Connection) CreateTeacherEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var teacher Teacher //создаётся объект по типу структуры Teacher
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

// Функция для получения пользователей
func (connection Connection) GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []User // создаётся лист объектов по типу структуры User
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

// функция для получения расписания из базы данных
func (connection Connection) GetSheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Shedule // создаётся объект по типу структуры Shedule
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

// Функция для получения расписания преподавателя из соответсвующей коллекции
func (connection Connection) GetTeachersScheduleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var shed []Teachers_Shedule
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

// функция для получения всех пользователей с ролью Студент
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

// функция для получения экзаменов
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

// функция для получения всех пользователей с ролью Учитель
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

// функция для даления пользователя из базы данных
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
	opts := options.Client().ApplyURI("mongodb+srv://myatakolyajr:0968715133@cluster0.ebfmjow.mongodb.net/").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	connection := Connection{ // реализация структуры Connection
		Shedule:           client.Database("proj").Collection("shedule"),           // подключение к коллекции
		Users:             client.Database("proj").Collection("users"),             // подключение к коллекции
		Exams:             client.Database("proj").Collection("exams"),             // подключение к коллекции
		Teachers_Schedule: client.Database("proj").Collection("teachers_schedule"), // подключение к коллекции
	}

	router := mux.NewRouter()

	router.HandleFunc("/user", connection.CreateUserEndpoint).Methods("POST")                      //обработчик эндпоинта /user в методе POST
	router.HandleFunc("/teacher", connection.CreateTeacherEndpoint).Methods("POST")                //обработчик эндпоинта /teacher в методе POST
	router.HandleFunc("/users", connection.GetUsersEndpoint).Methods("GET")                        //обработчик эндпоинта /users в методе GET
	router.HandleFunc("/teachers", connection.GetTeachersEndpoint).Methods("GET")                  //обработчик эндпоинта /teachers в методе GET
	router.HandleFunc("/students", connection.GetStudentsEndpoint).Methods("GET")                  //обработчик эндпоинта /students в методе GET
	router.HandleFunc("/schedule", connection.GetSheduleEndpoint).Methods("GET")                   //обработчик эндпоинта /schedule в методе GET
	router.HandleFunc("/teachers_schedule", connection.GetTeachersScheduleEndpoint).Methods("GET") //обработчик эндпоинта /teachers_schedule в методе GET
	router.HandleFunc("/exams", connection.GetExamsEndpoint).Methods("GET")                        //обработчик эндпоинта /exams в методе GET
	router.HandleFunc("/user/{id}", connection.DeleteUserEndpoint).Methods("DELETE")               //обработчик эндпоинта /user/{id} в методе DELETE

	http.ListenAndServe("localhost:8000", router) //Активация нашего API
}
