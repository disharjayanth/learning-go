package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func login(uid, pwd string) error {
	// return nil
	return errors.New("cannot login")
}

func getData(file string) ([]byte, error) {
	return []byte{}, nil
	// return []byte{}, errors.New("cannot get data")
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func main() {
	sb, err := LoginAndGetData("43566", "somePWD", "someFile")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Slice of byte:", sb)
}
