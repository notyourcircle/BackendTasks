## Instruction
### How to run storage system
This REST API used MySQL for storage system, edit code 
```
gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/anekazoo?charset=utf8&parseTime=True&loc=Local")
```
with
```
gorm.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<database>?charset=utf8&parseTime=True&loc=Local")
```
Don't forget to open  XAMPP and activate MySQL
Don't forget to make new database which will be used

### How to run the application
This REST API used GIN and GORM framwork, so dont forget to do
```
go get gopkg.in/gin-gonic/gin.v1
go get -u github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
```
if you not copy go.mod and go.sum file
To run this application just do 
```
go run main.go
```
after that we must use postman to do POST, PUT, DELETE, GET
- to POST, set postman to POST and fill that with ```localhost:8080/v1/animal/```
then set body to raw then select JSON and fill the body like this
```
{
    "Name"  :"cat",
    "Class" :"mammal",
    "Legs"  :4
}
```
id will be filled automatically(auto-increment)


- to PUT, set postman to PUT and fill that with ```localhost:8080/v1/animal/{id}``` example ```localhost:8080/v1/animal/1```
after that we can update data use same method like POST


- to GET all currently existing data, set postman to GET and fill that with ```localhost:8080/v1/animal/```

- to GET single data, set postman to GET and fill that with ```localhost:8080/v1/animal/{id}```

- to DELETE, set postman to DELETE and fill that with ```localhost:8080/v1/animal/{id}```

### Addresses of the API
Addresses of the API are ```localhost:8080/v1/animal``` and ```localhost:8080/v1/animal/{id}``` for PUT,GET single data, and DELETE function
