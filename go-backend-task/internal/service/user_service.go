package service

import (
	"context"
	"go-backend-task/internal/models"
	"go-backend-task/internal/repository"
	"time"

	db "go-backend-task/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	store *repository.Store
}

func NewUserService(store *repository.Store) *UserService {
	return &UserService{store: store}
}

// Helper to calculate age
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	parsedDOB, _ := time.Parse("2006-01-02", req.DOB)

	// Fix 1: Convert time.Time to pgtype.Date
	pgDate := pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	}

	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  pgDate,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	// Fix 2: Use .Time to get the standard time back
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.store.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Fix 3: Access .Time for calculation
	age := calculateAge(user.Dob.Time)

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.store.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Time.Format("2006-01-02"),
			Age:  calculateAge(u.Dob.Time),
		})
	}
	return response, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	parsedDOB, _ := time.Parse("2006-01-02", req.DOB)

	pgDate := pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	}

	user, err := s.store.UpdateUser(ctx, db.UpdateUserParams{
		Name: req.Name,
		Dob:  pgDate,
		ID:   id,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.store.DeleteUser(ctx, id)
}