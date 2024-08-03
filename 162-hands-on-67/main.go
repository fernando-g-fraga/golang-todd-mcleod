package main

import (
	"fmt"
)

type User struct {
	ID    int
	first string
}

type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]

	if !ok {
		return User{}, fmt.Errorf("User not fund: %v", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("o usuário já existe")
	}

	md.Users[u.ID] = u
	return nil

}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	s := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		first: "Fernando",
	}

	err := s.SaveUser(u1)
	if err != nil {
		fmt.Println("Erro ao salvar o usuário", err)
	}

	returnedUser, err := s.GetUser(u1.ID)

	if err != nil {
		fmt.Println("Usuário inexistente na base de dados", err)
	}

	fmt.Println(u1)
	fmt.Println(returnedUser)

}
