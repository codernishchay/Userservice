# User Service

Database used : Mongodb 

Language : Golang 


User Model 

```go
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	DOB         string             `json:"dob" bson:"dob"`
	Address     string             `json:"address" bson:"address"`
	Description string             `json:"description" bson:"description"`
}
```

this app consists of 4 endpoints : 

     create : to create a user to database.  

     update : to update a user using their user id

     delete : to delete user using their user id 

     get  : to get list of all users 

Update

![s4.png](User%20Service%20db98ab5a8c844f07ae8025765cdc2cb4/s4.png)

Delete 

![s1.png](User%20Service%20db98ab5a8c844f07ae8025765cdc2cb4/s1.png)

GET 

![s44.png](User%20Service%20db98ab5a8c844f07ae8025765cdc2cb4/s44.png)

POST

![s2.png](User%20Service%20db98ab5a8c844f07ae8025765cdc2cb4/s2.png)
