package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
)

type SignedDetail struct {
	FirstName string
	LastName  string
	Email     string
	Uid       uint
	UserType  string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateToken(user *models.User) (models.Token, string, error) {
	claims := &SignedDetail{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Uid:       user.ID,
		UserType:  user.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token_value, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return models.Token{}, token_value, err
	}

	temp := models.GetRefreshTokenByUserId(user.ID)
	if temp.ID == 0 {
		refresh_token_value, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
		if err != nil {
			return models.Token{}, token_value, err
		}

		token := &models.Token{
			Refresh: refresh_token_value,
			UserId:  user.ID,
		}

		token = token.CreateToken()
		return *token, token_value, nil
	} else {
		token := &models.Token{
			Refresh: temp.Refresh,
			UserId:  user.ID,
		}
		return *token, token_value, nil
	}
}

func GetUserFromToken(token string) (*models.User, error) {
	var user *models.User = &models.User{}
	t, err := jwt.ParseWithClaims(token, &SignedDetail{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if claims, ok := t.Claims.(*SignedDetail); ok && t.Valid {
		user.FirstName = claims.FirstName
		user.LastName = claims.LastName
		user.Email = claims.Email
		user.UserType = claims.UserType
		user.ID = claims.Uid

		return user, nil
	} else {
		return user, err
	}
}
