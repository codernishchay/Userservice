# User Service

Setup 

```go
git clone https://github.com/codernishchay/Userservice 

cd UserService 

// create a .env file and add 
// MONGOURI= "databse url"  

go run main.go  
```

 

User Model

```go
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	DOB         string             `json:"dob" bson:"dob"`
	Address     string             `json:"address" bson:"address"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   string               
}
```

/get 

```go
To get all the users from database. 
```

![Readme1.png](User%20Service%20ae9bd4f32c74436fbac7b609ba96e524/Readme1.png)

/create

```go
To create a new user to the database. 
```

![Readme2.png](User%20Service%20ae9bd4f32c74436fbac7b609ba96e524/Readme2.png)

/update

```go
To update a user data using its user ID 
```

![Readme3.png](User%20Service%20ae9bd4f32c74436fbac7b609ba96e524/Readme3.png)

/delete

```go
To delete user from database using the userID 
```

![Readme4.png](User%20Service%20ae9bd4f32c74436fbac7b609ba96e524/Readme4.png)