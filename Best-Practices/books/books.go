package books

import (
	"slices"
)

type Solution struct {
	books         map[string]string
	borrowStatus  map[string]bool
	borrowHistory map[string][]string
	// TODO: Add a field for tracking reservations using appropriate data structure
	reseveStatus map[string][]string
}

func NewSolution() *Solution {
	return &Solution{
		books:         make(map[string]string),
		borrowStatus:  make(map[string]bool),
		borrowHistory: make(map[string][]string),
		// TODO: Initialize the reservations field with the appropriate data structure
		reseveStatus: make(map[string][]string),
	}
}

func (s *Solution) AddBook(bookId string, title string) bool {
	if _, exists := s.books[bookId]; !exists {
		s.books[bookId] = title
		s.borrowStatus[bookId] = false
		return true
	}
	return false
}

func (s *Solution) CheckAvailability(bookId string) string {
	if title, exists := s.books[bookId]; exists && !s.borrowStatus[bookId] {
		return title
	}
	return ""
}

func (s *Solution) BorrowBook(userId string, bookId string) bool {
	if _, exists := s.books[bookId]; exists && !s.borrowStatus[bookId] {
		s.borrowStatus[bookId] = true
		s.UpdateBorrowHistory(userId, bookId)
		return true
	}
	return false
}

func (s *Solution) ReturnBook(bookId string) bool {
	if s.borrowStatus[bookId] {
		// TODO: Check if there are reservations for this book and update borrowing status if needed
		if len(s.reseveStatus[bookId]) > 0 {
			// last := len(s.reseveStatus[bookId])-1
			s.reseveStatus[bookId] = s.reseveStatus[bookId][1:]
			return true
		}
		s.borrowStatus[bookId] = false
		return true
	}
	return false
}

func (s *Solution) GetBorrowHistory(bookId string) []string {
	if history, exists := s.borrowHistory[bookId]; exists {
		return history
	}
	return []string{}
}

// TODO: Implement the ReserveBook function to add a user to the reservation queue for a borrowed book
func (s *Solution) ReserveBook(userId, bookId string) bool {
	if _, exists := s.books[bookId]; exists && s.borrowStatus[bookId] && !slices.Contains(s.reseveStatus[bookId], userId) {
		s.reseveStatus[bookId] = append(s.reseveStatus[bookId], userId)
		s.UpdateBorrowHistory(userId, bookId)
		return true
	}
	return false
}

// TODO: Implement the CheckReservation function to return the next user in line for a reserved book
func (s *Solution) CheckReservation(bookId string) string {
	if len(s.reseveStatus[bookId]) > 0 {
		return s.reseveStatus[bookId][len(s.reseveStatus[bookId])-1]
	}
	return ""
}

func (s *Solution) UpdateBorrowHistory(userId string, bookId string) {
	s.borrowHistory[bookId] = append(s.borrowHistory[bookId], userId)
}
