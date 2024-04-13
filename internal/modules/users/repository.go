package users

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUser(id uint64) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUser(id uint64, data UpdateUserDto) (User, error)
	DeleteUser(id uint64) error
}

/*
----------------------------------------------------------------------------

	Repository

----------------------------------------------------------------------------
*/
type UsersSQLRepository struct {
	db *gorm.DB
}

func NewUsersSQLRepository(db *gorm.DB) *UsersSQLRepository {
	db.AutoMigrate(&User{})

	return &UsersSQLRepository{db: db}
}

func (r *UsersSQLRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *UsersSQLRepository) GetAllUsers() ([]User, error) {
	users := []User{}
	r.db.Find(&users)

	return users, nil
}

func (r *UsersSQLRepository) GetUser(id uint64) (User, error) {
	user := User{}
	result := r.db.Find(&user, id)
	if result.RowsAffected == 0 {
		return User{}, fmt.Errorf("User not found")
	}

	return user, nil
}

func (r *UsersSQLRepository) GetUserByEmail(email string) (User, error) {
	user := User{}
	result := r.db.Where("email = ?", email).Find(&user)
	if result.RowsAffected == 0 {
		return User{}, fmt.Errorf("User not found")
	}

	return user, nil
}

func (r *UsersSQLRepository) UpdateUser(id uint64, data UpdateUserDto) (User, error) {
	user := User{}
	result := r.db.Find(&user, id)
	if result.RowsAffected == 0 {
		return User{}, fmt.Errorf("User not found")
	}

	user.Name = data.Name
	r.db.Save(&user)

	return user, nil
}

func (r *UsersSQLRepository) DeleteUser(id uint64) error {
	user := User{}
	result := r.db.Find(&user, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("User not found")
	}

	r.db.Delete(&user)

	return nil
}

/*
----------------------------------------------------------------------------

	Mock Repository

----------------------------------------------------------------------------
*/
type UsersMockRepository struct {
	table map[uint64]User
	m     sync.RWMutex
}

func NewUsersMockRepository() *UsersMockRepository {
	return &UsersMockRepository{table: make(map[uint64]User)}
}

func (r *UsersMockRepository) CreateUser(user User) (User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	id := uint(rand.Uint64())

	user.Model.ID = id
	r.table[uint64(id)] = user

	return user, nil
}

func (r *UsersMockRepository) GetAllUsers() ([]User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	users := make([]User, 0, len(r.table))

	for _, v := range r.table {
		users = append(users, v)
	}

	return users, nil
}

func (r *UsersMockRepository) GetUser(id uint64) (User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	user, ok := r.table[id]
	if !ok {
		return User{}, fmt.Errorf("User not found")
	}

	return user, nil
}

func (r *UsersMockRepository) GetUserByEmail(email string) (User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, user := range r.table {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User not found")
}

func (r *UsersMockRepository) UpdateUser(id uint64, data UpdateUserDto) (User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	user, ok := r.table[id]
	if !ok {
		return User{}, fmt.Errorf("User not found")
	}

	user.Name = data.Name
	user.UpdatedAt = time.Now()
	r.table[id] = user

	return user, nil
}

func (r *UsersMockRepository) DeleteUser(id uint64) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.table[id]; !ok {
		return fmt.Errorf("User not found")
	}

	delete(r.table, id)
	return nil
}
