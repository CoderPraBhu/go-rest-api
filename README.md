Reference: https://go.dev/doc/tutorial/web-service-gin
Gin Web Framework: https://pkg.go.dev/github.com/gin-gonic/gin
Gin Docs: https://gin-gonic.com/docs/


```
go mod init coderprabhu.con/go-rest-api
go get .
go run .
curl localhost:8080/albums
[
    {
        "id": "1",
        "title": "One Love",
        "artist": "Blue",
        "price": 10.99
    },
    {
        "id": "2",
        "title": "Legend",
        "artist": "Bob Marley",
        "price": 12.99
    },
    {
        "id": "3",
        "title": "Nothing but the best",
        "artist": "Bob Dylan",
        "price": 12.99
    }
] 

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Scientist","artist": "Coldplay","price": 10.99}'

curl localhost:8080/albums/2
{
    "id": "2",
    "title": "Legend",
    "artist": "Bob Marley",
    "price": 12.99
} 

```

Git 
```
git init
git status
git checkout -b main
git add .
git commit -m "Go Rest Api Using Gin"
git remote add origin https://github.com/CoderPraBhu/go-rest-api.git
git push -u origin main


```