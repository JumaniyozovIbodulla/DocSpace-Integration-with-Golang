package main

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" 
)


const (
	JwtSecret     = "mysecret123" 
	DocSpaceURL   = "http://localhost:8090" 
)

type DocSpaceConfig struct {
	DocSpaceURL string
	Token       string 
}


func generateToken(userID, roomID string) (string, error) {

	claims := jwt.MapClaims{
		"userid": userID,
		"roomid": roomID, 
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 1).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}


func main() {
	r := gin.Default()
	r.LoadHTMLGlob(filepath.Join("templates", "*"))

	r.GET("/", func(c *gin.Context) {
        
        currentUserID := "user-123"
        currentRoomID := "default-room"
        
		token, err := generateToken(currentUserID, currentRoomID)
		if err != nil {
			log.Println("JWT generatsiyasida xatolik:", err)
			c.String(http.StatusInternalServerError, "Token yaratilmadi.")
			return
		}
        
		config := DocSpaceConfig{
			DocSpaceURL: DocSpaceURL,
			Token:       token, 
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"DocSpaceURL": config.DocSpaceURL,
			"Token":       config.Token,
		})
	})

	log.Println("Go/Gin serveri http://localhost:8080 da ishlamoqda")
	log.Fatal(r.Run(":8080"))
}