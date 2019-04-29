/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func AddSecret(w http.ResponseWriter, r *http.Request) {
	secret := Secret{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&secret)
	if err != nil {
		panic(err)
	}
	data := []byte(secret.SecretText)
	hash := md5.New().Sum(data)
	createdAt := time.Now()
	expiresAt := time.Now().AddDate(0, 0, 1)
	remainingViews := 10
	secret.Hash = string(hash)
	secret.CreatedAt = createdAt
	secret.ExpiresAt = expiresAt
	secret.RemainingViews = int32(remainingViews)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs, _ := json.Marshal(secret)
	log.Printf(string(rs))
	Create()
	// w.WriteHeader(http.StatusOK)

}

func GetSecretByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
