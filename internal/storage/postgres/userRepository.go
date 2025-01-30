package db

import (
	"context"
	"errors"
	"log"

	"github.com/ShayeGun/go-server/models"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepository struct {
	q   *Queries
	ctx context.Context
}

func NewUserRepository(ctx context.Context, q *Queries) *UserRepository {
	return &UserRepository{
		q:   q,
		ctx: ctx,
	}
}

func (r *UserRepository) GetById(id string) (
	models.User, error) {

	oid := pgtype.UUID{}

	if err := oid.Scan(id); err != nil {
		log.Println(err)
		return models.User{}, err
	}

	u, err := r.q.GetUserByID(r.ctx, oid)
	if err != nil {
		log.Println(err)
		return models.User{}, errors.New("error in finding user")
	}

	mu := models.User{}
	mu.SetEmail(u.Email)
	mu.SetPassword(u.Password)
	mu.SetID(id)

	return mu, nil
}

func (r *UserRepository) Add(user models.User) (models.User, error) {
	u, err := r.q.CreateUser(r.ctx, CreateUserParams{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	})

	if err != nil {
		log.Println(err)
		return models.User{}, errors.New("error in creating user")
	}

	strID, err2 := u.ID.Value()

	if err2 != nil {
		log.Println(err2)
		return models.User{}, errors.New("error in parsing user ID")
	}

	mu := models.User{}
	mu.SetEmail(u.Email)
	mu.SetPassword(u.Password)
	mu.SetID(strID.(string))

	return mu, nil
}

func (r *UserRepository) Delete(id string) error {

	oid := pgtype.UUID{}

	if err := oid.Scan(id); err != nil {
		log.Println(err)
		return err
	}

	err := r.q.DeleteUser(r.ctx, oid)
	if err != nil {
		log.Println(err)
		return errors.New("error in finding user")
	}

	return nil
}

func (r *UserRepository) Update(user models.User) (models.User, error) {

	oid := pgtype.UUID{}

	if err := oid.Scan(user.GetID()); err != nil {
		log.Println(err)
		return models.User{}, err
	}

	u, err := r.q.UpdateUser(r.ctx, UpdateUserParams{
		ID:       oid,
		Password: user.GetPassword(),
	})

	if err != nil {
		log.Println(err)
		return models.User{}, errors.New("error in updating user")
	}

	strID, err2 := u.ID.Value()

	if err2 != nil {
		log.Println(err2)
		return models.User{}, errors.New("error in parsing user ID")
	}

	mu := models.User{}
	mu.SetEmail(u.Email)
	mu.SetPassword(u.Password)
	mu.SetID(strID.(string))

	return mu, nil
}
