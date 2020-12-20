package main
import(
	"context"
	"log"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Users struct{
	ID primitive.ObjectID //`json:"id" bson:"_id,omitempty"`
	Name string //`json:"name"`
	Date_of_birth time.Time //`json:"dob"`
	Phone_Number int //`json:"phone_no"`
	Email_Address string //`json:"email"`
	CreatedDate time.Time //`json:"createddate"`
}
func main(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(context.TODO(), clientOptions)
	if err !=nil{
		log.Fatal(err)
	}
	err=client.Ping(context.TODO(),nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongodb!")
	collection := client.Database("test").Collection("Users")
	err=client.Disconnect(context,TODO())
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Disconnect")
}