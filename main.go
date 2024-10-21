package main

import (
	"fmt"
	app "golang-shopeeMart-api/App"
	implController "golang-shopeeMart-api/Controller/impl"
	exception "golang-shopeeMart-api/Exception"
	helper "golang-shopeeMart-api/Helper"
	implRepository "golang-shopeeMart-api/Repository/impl"
	implService "golang-shopeeMart-api/Service/impl"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	validate := validator.New()
	db := app.NewDb()

	matakuliahRepository := implRepository.NewMatakuliahRepositori()
	matakuliahService := implService.NewMatakuliahService(matakuliahRepository, db, validate)
	matakuliahController := implController.NewMatakuliahController(matakuliahService)

	router := httprouter.New()

	router.POST("/api/matakuliah", matakuliahController.Create)
	router.PUT("/api/matakuliah/:matakuliahId", matakuliahController.Update)
	router.GET("/api/matakuliah/:matakuliahId", matakuliahController.FindById)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server Started , Listening on : ", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		helper.PanicIfErr(err)
	}
}
