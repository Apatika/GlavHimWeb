package service

import "encoding/json"

type User struct {
	Name  string
	Tel   string
	Email string
}

func marshalUsers(users []User) ([]byte, error) {
	result, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return result, nil
}
