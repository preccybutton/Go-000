package server

import "net/http"

func FirstHandlec (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("first"))
}

func SecondHandlec (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("second"))
}