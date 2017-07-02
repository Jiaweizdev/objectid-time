package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main() {

	// Expect an ObjectId passed as the first CLI parameter
	input := os.Args[1]

	// Check if it is a valid ObjectID
	if bson.IsObjectIdHex(input) {
		extractDateFromObjectId(bson.ObjectIdHex(input))
	} else {
		fmt.Println("The input is not a valid ObjectId.")
	}
}

func extractDateFromObjectId(id bson.ObjectId) {

	// Time in UTC
	fmt.Println(id.Time().UTC())
	// Time in Local Timezone
	fmt.Println(id.Time())
}
