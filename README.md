# HearX technical assessment

- language: Go 1.19

---

## External Packages used:
```
github.com/joho/godotenv
github.com/gin-gonic/gin
github.com/lib/pq
github.com/stretchr/testify

```

---

## Design Decisions and Assumptions

 **GIN Framework** - Used this framework to speed up development, since it is well tested, robust and comes with a lot of features built in. Other ideas considered:
 - `gorilla/mux` - Less features would have required more build time but offered more control of the handlers and works similar to the standard librtary package
 - built in `net/http` package` - would have required a longer devlopment time as a lot of the features such as middlewares would have to be build from scratch. The implementation is a short time would not be as robust as when using a framework.

 **AWS DB** - 
 Other ideas considered:
  - `Docker` file with Postgres service, felt it would add unnecessary complexity to just run the app for assessment purposes

 **local.env** - 

 **migrations** - 

 **CRUD** - 

 **Pagination** - 

 **ORM** - 

 **Duplicates** - 

 ---

### Running the app

App can be run locally either by using `go run` or alternatively building and running the binary. Navigate to the root directory of the project, run `go mod tidy` to get the dependencies. then either:
```
go run main.go
```
or
```
go build .
./hearxtest
```

**Running unit tests**

```
go test -v ./...
```
---

**NOTE** - API calls require `Authorization` header, with a value of `Bearer bWFyY2xhdWRlcg==` (can be configured in env file)

