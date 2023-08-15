package dbConfig

import (
	firestore "cloud.google.com/go/firestore"
	"go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database
var DATABASE_FireBase *firestore.Client

const DATABASE_URL = "mongodb+srv://chamith:123@cluster0.ujlq82i.mongodb.net/?retryWrites=true&w=majority"

//const DATABASE_URL = "mongodb+srv://pasinduruwantha12:1234@cgaasui.amademj.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME = "AWSLambda"

