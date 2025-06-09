package user

import "time"

type User struct {
	id          UserID
	email       Email
	password    PasswordHash
	displayName DisplayName
	createdAt   time.Time
	updatedAt   time.Time
}

func NewUser(
	email Email,
	password PasswordHash,
	displayName DisplayName,
) *User {
	now := time.Now()
	return &User{
		id:          NewUserID(),
		email:       email,
		password:    password,
		displayName: displayName,
		createdAt:   now,
		updatedAt:   now,
	}
}

// Reconstruct は永続化層からユーザーを再構築します。
func Reconstruct(
	id UserID,
	email Email,
	password PasswordHash,
	displayName DisplayName,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		id:          id,
		email:       email,
		password:    password,
		displayName: displayName,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

// ID はユーザーIDを返します。
func (u *User) ID() UserID {
	return u.id
}

// Email はメールアドレスを返します。
func (u *User) Email() Email {
	return u.email
}

// PasswordHash はパスワードハッシュを返します。
func (u *User) PasswordHash() PasswordHash {
	return u.password
}

// DisplayName は表示名を返します。
func (u *User) DisplayName() DisplayName {
	return u.displayName
}

// CreatedAt は作成日時を返します。
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt は更新日時を返します。
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
