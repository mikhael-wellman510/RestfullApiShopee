// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// type Data struct {
// 	name string
// }

// func FindData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	w.Header().Add("Content-Type", "application/json")
// 	response := map[string]string{
// 		"Name": "Mikhael Wellman",
// 	}

// 	err := json.NewEncoder(w).Encode(response)
// 	fmt.Println(err)
// }

// func AddData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// Ubah Json ke Map
// 	data := &map[string]string{}
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(data)
// 	fmt.Println(err)

// 	fmt.Println("Data req : ", data)
// 	w.Header().Add("Content-Type", "application/json")
// 	errData := json.NewEncoder(w).Encode(data)
// 	fmt.Println(errData)

// }

// type Profile struct {
// 	Name string
// 	Age  int
// }

// type ListProfile struct {
// 	ListProfile []Profile
// }

// func (list *ListProfile) AddData(data *Profile) {
// 	list.ListProfile = append(list.ListProfile, *data)
// }

// // Instance Global , karena klo d dalam fungsi . akan reset ulang
// var profileList *ListProfile = &ListProfile{}

// func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	profile := &Profile{}
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(profile)
// 	fmt.Println(err)

// 	profileList.AddData(profile)

// 	w.Header().Add("Content-Type", "application/json")
// 	errData := json.NewEncoder(w).Encode(profile)
// 	fmt.Println(errData)
// }

// func GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	fmt.Println(profileList)
// }

// func main() {

// 	router := httprouter.New()

// 	router.GET("/", FindData)
// 	router.POST("/create", AddData)
// 	router.POST("/profil", Create)
// 	router.GET("/findAll", GetList)

// 	log.Fatal(http.ListenAndServe(":8000", router))

// }
