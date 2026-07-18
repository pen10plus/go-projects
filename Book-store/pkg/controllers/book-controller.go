package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
    "net/http"
	"strconv"
	"github.com/pen10plus/Book-store/pkg/models"
	"github.com/pen10plus/Book-store/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks := models.GetAllBooks()
	res,_=json.Marshal(newBooks)
	w.Header().set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter,r *http.Request){
vars :=mux.Vars(r)
bookId :=vars["bookId"]
ID , err:= strconv.ParseInt(bookId,0,0)
if err != nil{
fmt.Println("error while parsing")
}

bookDetails,_ := models.GetBookById(ID)

res,_ := json.Marshal(bookDetails)
w.Header().Set("content-type", "pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){
CreateBook := &models.Book{}
utils.ParseBody(r,CreateBook)
b := CreateBook.CreateBook()
res,_ := json.Marshal(b)
w.WriteHeader(http.StatusOK)
w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request){
vars := mux.Vars(r)
bookId := vars["bookId"]
ID, err := strconv.ParseInt(bookId, 0, 0)


