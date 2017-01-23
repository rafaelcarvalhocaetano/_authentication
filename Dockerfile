FROM golang:1.17

WORKDIR /auth

RUN go get github.com/gofiber/fiber/v2 && \
go get -u gorm.io/gorm && \
go get -u gorm.io/driver/postgres && \
go get golang.org/x/crypto/bcrypt && \
go get github.com/dgrijalva/jwt-go

COPY . .

CMD [ "go", "run", "main.go" ]