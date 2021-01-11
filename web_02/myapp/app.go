package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	// json annotation
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"created_at"`
}

type fooHandler struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	// response의 body를 받아서 user struct에 맞게 디코딩
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request : ", err)
		return
	}
	user.CreateAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// name argument 넘기기
	// http://localhost:3000/bar?name=222
	name := r.URL.Query().Get("name") // name argument 가져오기
	if name == "" {
		name = "default name"
	}
	fmt.Fprintf(w, "%s", name)
}

// NewHttpHandler using mux
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	// 직접 func 형태로 등록
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	// 인스턴스 형태로 등록
	mux.Handle("/foo", &fooHandler{})

	return mux
}
