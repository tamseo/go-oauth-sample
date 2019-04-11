*How to run*

```
go get github.com/gorilla/mux
go get github.com/auth0/go-jwt-middleware
go get github.com/dgrijalva/jwt-go
go run main.go

```

*Deploy to GAE*

```
gcloud app deploy app.yaml --project <PROJECT_KEY>
```

*To access private API*
```
curl -H "Authorization:Bearer <token>" http://localhost:8080/private
```
