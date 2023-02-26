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

## Running the app

 **local.env** - The app will require a `local.env` file to be created to work locally. It was not pushed to the repo for security purposes as it can contain sensitve data. The variables required are in the sample file below. This can be copied into a file called `local.env` and updated with values necessary.

 sample `.env` file:
 ```
PG_HOST: <host>:<port>
PG_USER: <postgres_user>
PG_PASSWORD: <postgres_password>
PG_DB: <postgres_database_name>
PG_SSL: <postgres_ssl_setting>
BEARER_API_KEY: <hard_coded_API_key>
 ```

Once the env file is in place in the root of the project, the app can be run locally either by using `go run` or alternatively building and running the binary. Navigate to the root directory of the project, run `go mod tidy` to get the dependencies. then either:
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

### Endpoints

I have included a Postman collection for testing, with all the relevant endpoints. an overview.

 - `GET <host>/api/v1/jokes` - returns a single random joke
 - `GET <host>/api/v1/jokes/page/<page_number>` - returns a paginated list of jokes, along with the total jokes and current page. Also accepts a optional query param of `records` to specify the number of records per page, if something other than the default 10 is required.
 - `POST <host>/api/v1/jokes` - accepts and JSON array of joke objects to save

 **NOTE** - POST  endpoint requires `Authorization` header, with a value of `Bearer <secret_key>` as defined in the `env` file.

## Design Decisions and Assumptions

 **GIN Framework** - Used this framework to speed up development, since given the time, it would not have been possible to build something as robust and well tested as a 3rd party framework. It also comes with a lot of features letteing me focus on the business logic. Other ideas considered:
 - `gorilla/mux` - Less features would have required more build time but offered more control of the handlers and works similar to the standard library package.
 - built in `net/http` package` - would have required a longer devlopment time as a lot of the features such as middlewares would have to be build from scratch. The implementation is a short time would not be as robust as when using a framework.

 **Database** - For the assessment, I assumed that a database would be available on your side to test the app with. the credentials will need to be configured in the `local.env` file. Other ideas considered:
  - `Docker` file with Postgres service, felt it would add unnecessary complexity to just run the app for assessment purposes.

 **API Key** - An assumption was made that have the API key in the env file would be suffcient for this assesment, in a real world case, I realise this would never be the case and would be database and user driven.

 **Pagination/Random** - The method to get the total jokes and to return the randomized joke both rely on retreiveing a full list of ids from the database, An assumption was made that this would be adequate for the assessement, but in a database that grows to large number of rows, thia would be ineffecient.

 **Duplicates** - An assumption was made to allow duplicate entries, as this would be very hard to verify in the time given, as there is no unique key, though of using a hash value, but this is not guaranteed to work as expected, since a single change would result in a different hash.

 **Request timeouts** - An assumption was made that for this assessment, not timeouts on the requests would be necessary.
 
 ---

## Potential improvements

 **Migrations** - No work was done to able database migrations, so this would be a potential improvement

 **ORM** - Although for this assessment, i wanted to demonstrate the SQL queries, in a real world app, it ould make sense to simplfy the datbase layer by using a ORM package such as `gorm.io/gorm`

 **CRUD** - Given more time, additional endpoints could have been added to create a full update and delete option for the jokes to provide full CRUD capability

 **Better test coverage** - Additional tests could have been written, such as tests on the individual endpoints to confirm the are configured correctly, sonce the underlying logic has already been tested.