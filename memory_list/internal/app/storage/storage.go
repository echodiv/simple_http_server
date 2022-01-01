package storage

import (
	"errors"
	"log"
)

type LinkedList struct {
	Name  string
	Value string
	Next  *LinkedList
}

type Storage struct {
	empty bool
	root  *LinkedList
}

func NewStorage() *Storage {
	log.Println("Create new storage")
	return &Storage{
		empty: true,
	}
}

func (s *Storage) AddNewElement(name string) error {
	if _, err := s.GetElementByName(name); err == nil {
		return errors.New("element exist")
	}
	new := &LinkedList{
		Name: name,
	}
	new.Next = s.root
	s.root = new
	if s.empty {
		s.empty = false
	}
	log.Printf("Created with %v", s.root.Name)
	return nil
}

func (s *Storage) GetElementByName(name string) (*LinkedList, error) {
	log.Println("search element:", name, len(name))
	if s.empty {
		log.Println("storage is empty :(")
		return nil, errors.New("storage is empty")
	}
	searchPoint := s.root
	for searchPoint.Next != nil {
		log.Println(">> find by name: ", searchPoint.Name, len(searchPoint.Name))
		if searchPoint.Name == name {
			log.Println(">>>> SUCCESS")
			return searchPoint, nil
		}
		searchPoint = searchPoint.Next
	}
	return nil, errors.New("cant find object")
}
