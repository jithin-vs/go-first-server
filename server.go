package main

import (
	"context"
	"fmt"
	"net/http"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	mux := http.NewServeMux()
    uri := "mongodb+srv://jithinvs1045:122333@cluster0.563o5.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	mux.HandleFunc("/",handleHome)
	connectMongoDB(uri)
	fmt.Println("server listening on port 8000")
	http.ListenAndServe(":8000",mux)
}

func handleHome(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"welcome to go server")
}

func connectMongoDB(uri string){
   client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI(uri))
   if err!= nil {
	 panic(err)
   }
   fmt.Println("Connected to MongoDB")
   defer func () {
	 if err = client.Disconnect(context.TODO());err!=nil{
		panic(err)
	 }
   }()
}