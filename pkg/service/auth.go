package service

import (
  "crypto/sha1"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/repository"
  "time"
)

const (
  salt ="sdnlbdlkfb423156sad"
  signingKey = "sdjkfbskdjbfcowb456464"
  tokenTTL = 12 * time.Hour
)
type AuthService struct {
  repo repository.Authorization
}

type tokenClaims struct{
  jwt.StandardClaims
  UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
  return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser( user todo.User) (int, error) {
  user.Password = generatePasswordHash(user.Password)
  return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error){
  user, err := s.repo.GetUser(username, generatePasswordHash(password))
  if err != nil {
    return "", err
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
    jwt.StandardClaims{
    ExpiresAt: time.Now().Add(tokenTTL).Unix(),
    IssuedAt: time.Now().Unix(),
  },
  user.Id,
})
  return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
  hash := sha1.New()
  hash.Write([]byte(password))

  return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
