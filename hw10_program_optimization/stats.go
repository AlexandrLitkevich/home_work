package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"
)

//easyjson:json
type User struct {
	ID       int    `json:"Id"`       //nolint
	Name     string `json:"Name"`     //nolint
	Username string `json:"Username"` //nolint
	Email    string `json:"Email"`    //nolint
	Phone    string `json:"Phone"`    //nolint
	Password string `json:"Password"` //nolint
	Address  string `json:"Address"`  //nolint
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var user User
	lowDomain := strings.ToLower(domain)
	result := make(DomainStat)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		err := user.UnmarshalJSON(scanner.Bytes())
		if err != nil {
			return nil, err
		}

		currentEmail := strings.ToLower(user.Email)

		ok := strings.Contains(currentEmail, lowDomain)
		if ok {
			em := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[em]++
		}
	}
	return result, nil
}
