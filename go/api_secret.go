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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type secretRequestBody struct {
	Secret           string `json:"secret"`
	ExpireAfterViews int32  `expireAfterViews:""`
	ExpireAfter      int32  `expireAfter:"secret"`
}

// AddSecret comment like this?
func AddSecret(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Code goes here 1")

	body := secretRequestBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	if validErrs := (&body).validate(); len(validErrs) > 0 {
		fmt.Println("Code goes here")
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	randStr := RandomString(10)

	a := randStr + body.Secret
	secretText := []byte(a)
	hash := hex.EncodeToString(md5.New().Sum(secretText)) //generate hash for secret
	createdAt := time.Now()
	expiresAt := time.Now().Add(time.Minute * time.Duration(body.ExpireAfter)) // The secret will be expired after X minutes.
	remainingViews := body.ExpireAfterViews                                    // 10 times for default

	secret := Secret{}
	secret.Hash = string(hash)
	secret.SecretText = body.Secret
	secret.CreatedAt = createdAt
	secret.ExpiresAt = expiresAt
	secret.RemainingViews = remainingViews
	returnHash := Create(secret)
	response, _ := json.Marshal(returnHash)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func GetSecretByHash(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path[11:]
	if hash == "" {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Secret hash is invalid.")
	}
	secretText := Show(hash)
	if secretText == "not found" {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Secret not found")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, secretText)
}

func (a *secretRequestBody) validate() url.Values {
	errs := url.Values{}

	// check if the Secret empty
	if a.Secret == "" {
		errs.Add("Secret", "Secret is required.")
	}

	// check the ExpireAfterViews field is > 0
	if a.ExpireAfterViews <= 0 {
		errs.Add("ExpireAfterViews", "ExpireAfterViews must be greater than 0.")
	}

	return errs
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
