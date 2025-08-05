package userservice

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/shatwik7/polycrate/lib/db"
)

type UserRepository struct {
	database *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{database: db}
}

func (repo *UserRepository) InsertUser(input CreateUserInput) (*User, error) {
	query := `INSERT INTO users (username, email, full_name, profile_picture_url, bio, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, now(), now())
	          RETURNING id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at`
	user := &User{}
	err := repo.database.QueryRow(query, input.Username, input.Email, input.FullName, input.ProfilePictureUrl, input.Bio).
		Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.ProfilePictureUrl, &user.Bio, &user.Website, &user.Location, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (repo *UserRepository) FindUserById(id uuid.UUID) (*User, error) {
	query := `SELECT id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at FROM users WHERE id = $1`
	user := &User{}
	err := repo.database.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.ProfilePictureUrl, &user.Bio, &user.Website, &user.Location, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repo *UserRepository) UpdateUser(input UpdateUserInput) (*User, error) {
	query := `UPDATE users SET full_name = $1, profile_picture_url = $2, bio = $3, updated_at = now()
	          WHERE id = $4
	          RETURNING id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at`
	user := &User{}
	err := repo.database.QueryRow(
		query,
		input.FullName,
		input.ProfilePictureUrl,
		input.Bio, input.ID).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.FullName,
			&user.ProfilePictureUrl,
			&user.Bio,
			&user.Website,
			&user.Location,
			&user.CreatedAt,
			&user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, err
}

func (repo *UserRepository) DeleteUser(id uuid.UUID) (bool, error) {
	res, err := repo.database.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return false, errors.New("no user deleted")
	}
	count, err := res.RowsAffected()
	return count > 0, err
}

func (repo *UserRepository) ListUsers(limit, offset int) ([]User, error) {
	query := `SELECT id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at FROM users LIMIT $1 OFFSET $2`
	rows, err := repo.database.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.ProfilePictureUrl, &user.Bio, &user.Website, &user.Location, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) FindUserByEmail(email string) (*User, error) {
	query := `SELECT id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at FROM users WHERE email = $1`
	user := &User{}
	err := repo.database.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.ProfilePictureUrl, &user.Bio, &user.Website, &user.Location, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (repo *UserRepository) FindUsersByUsernamePartial(partial string, limit, offset int) ([]User, error) {
	query := `SELECT id, username, email, full_name, profile_picture_url, bio, website, location, created_at, updated_at 
			  FROM users WHERE username ILIKE $1 LIMIT $2 OFFSET $3`
	rows, err := repo.database.Query(query, "%"+partial+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.ProfilePictureUrl, &user.Bio, &user.Website, &user.Location, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) InsertCredential(cred UserCredential) (bool, error) {
	query := `INSERT INTO user_credentials (user_id, password_hash, last_login, is_active) VALUES ($1, $2, $3, $4)`
	_, err := repo.database.Exec(query, cred.UserID, cred.PasswordHash, cred.LastLogin, cred.IsActive)
	return err == nil, err
}

func (repo *UserRepository) GetCredential(userID uuid.UUID) (*UserCredential, error) {
	query := `SELECT user_id, password_hash, last_login, is_active FROM user_credentials WHERE user_id = $1`
	cred := &UserCredential{}
	err := repo.database.QueryRow(query, userID).Scan(&cred.UserID, &cred.PasswordHash, &cred.LastLogin, &cred.IsActive)
	return cred, err
}

func (repo *UserRepository) UpdateCredential(cred UserCredential) (bool, error) {
	query := `UPDATE user_credentials SET password_hash = $1, last_login = $2, is_active = $3 WHERE user_id = $4`
	res, err := repo.database.Exec(query, cred.PasswordHash, cred.LastLogin, cred.IsActive, cred.UserID)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	return count > 0, err
}
