package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	io.WriteString(w,"hello")

}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	io.WriteString(w,p.ByName("user_name"))
}