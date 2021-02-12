//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"

	userM "dogen/once_pkg/domain/model"
	userR "dogen/once_pkg/domain/repository"
)

type userRepoImpl struct {
	db *sql.DB
}

func NewUserRepoImpl(db *sql.DB) userR.UserRepository {
	return &userRepoImpl{
		db,
	}
}

// Select
func (userI userRepoImpl) Select(ctx context.Context) ([]userM.User, error) {
	// rows, err := userI.db.Query("SELECT * FROM user")
	if err != nil {
		return err
	}
	return convertToUser(rows)
}

// Insert
func (userI userRepoImpl) Insert(ctx context.Context, entity *userM.User) error {
	// stmt, err := userI.db.Prepare("INSERT INTO user () VALUES ())
	if err != nil {
		return err
	}
	_, err := stmt.Exec()
	return err
}

// Update
func (userI userRepoImpl) Update(ctx context.Context, entity *userM.User) error {
	// stmt, err := userI.db.Prepare("UPDATE user SET WHERE ")
	if err != nil {
		return err
	}
	_, err := stmt.Exec()
	return err
}

// Delete
func (userI userRepoImpl) Delete(ctx context.Context, entity *userM.User) error {
	// stmt, err := userI.db.Prepare("DELETE FROM user WHERE ")
	if err != nil {
		return err
	}
	_, err := stmt.Exec()
	return err
}

// convertToUser
func convertToUser(rows *sql.rows) ([]userM.User, error) {
	var users []userM.User
	for rows.Next() {
		var user userM.User
		err := rows.Scan(
		// Need to scan field
		)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	return &users, nil
}
