package postgres

import (
	"database/sql"
	"log"
	"task/model"

	"github.com/google/uuid"
)

type UserRepo interface {
	Register(req *model.RegisterReq) (*model.RegisterResp, error) 
	CheckUser(email string) (*model.UserInfo, error)
	SaveToken(req *model.SaveTokenReq) error
}

type userImpl struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userImpl{
		Db: db,
	}
}

func (U *userImpl) Register(req *model.RegisterReq) (*model.RegisterResp, error) {
	id := uuid.NewString()

	query := `
				INSERT INTO users(
					id, username, email, password_hash)
				VALUES
					($1, $2, $3, $4)`
	_, err := U.Db.Exec(query, id, req.Username, req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.RegisterResp{
		UserId: id,
	}, nil
}

func (U *userImpl) CheckUser(email string) (*model.UserInfo, error) {
	resp := model.UserInfo{}
	query := `
				SELECT 
					id, username, password_hash
				FROM 
					users
				WHERE
					email = $1`
	err := U.Db.QueryRow(query, email).Scan(&resp.Id, &resp.Username, &resp.PasswordHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &resp, nil
}

func (U *userImpl) SaveToken(req *model.SaveTokenReq) error {
	query := `
				INSERT INTO refresh_token(
					user_id, refresh_token, expires_at)
				VALUES
					($1, $2, $3)`
	_, err := U.Db.Exec(query, req.UserId, req.RefreshToken, req.ExpiresAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
