package store

import (
	"github.com/1deep1/deepCraft_web/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	// if err := r.store.db.QueryRow(
	// 	//"INSERT INTO `users` (`email`, `username`, `password`) VALUES ('$1', '$2', MD5('$3')) RETURNING `id`;",
	// 	/*"INSERT INTO `users` (`uuid`, `id`, `email`, `username`, `password`, `accessToken`, `serverID`, `hwidId`) VALUES (NULL, NULL, ?, ?, ?, NULL, NULL, NULL)",
	// 	u.Email,
	// 	u.Username,
	// 	u.Password,
	// 	*/
	// 	"INSERT INTO `users` (`uuid`, `id`, `email`, `username`, `password`, `accessToken`, `serverID`, `hwidId`) VALUES (NULL, NULL, 'email', 'user', 'pass', NULL, NULL, NULL)",
	// ).Scan(&u.ID); err != nil {
	// 	return nil, err
	// }

	if err := r.store.db.QueryRow("INSERT INTO `users` (`uuid`, `id`, `email`, `username`, `password`, `accessToken`, `serverID`, `hwidId`) VALUES (NULL, NULL, 'email', 'user', 'pass', NULL, NULL, NULL)"); err != nil {
		return nil, nil
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT `id`, `email`, `username`, `password` FROM `users` WHERE `email` = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.Password,
	); err != nil {
		return nil, err
	}

	return nil, nil
}
