package store

import (
	"github.com/jinzhu/gorm"
	"github.com/faozimipa/golang-echo-realworld-example-app/model"

)

//UserStore struct
type UserStore struct {
	db *gorm.DB
}

//NewUserStore func 
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

//GetByID func 
func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//GetByEmail func 
func (us *UserStore) GetByEmail(e string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//GetByUsername func 
func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Username: username}).Preload("Followers").First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//Create func 
func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

//Update func 
func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

//AddFollower func 
func (us *UserStore) AddFollower(u *model.User, followerID uint) error {
	return us.db.Model(u).Association("Followers").Append(&model.Follow{FollowerID: followerID, FollowingID: u.ID}).Error
}

//RemoveFollower func 
func (us *UserStore) RemoveFollower(u *model.User, followerID uint) error {
	f := model.Follow{
		FollowerID:  followerID,
		FollowingID: u.ID,
	}
	if err := us.db.Model(u).Association("Followers").Find(&f).Error; err != nil {
		return err
	}
	if err := us.db.Delete(f).Error; err != nil {
		return err
	}
	return nil
}

//IsFollower func 
func (us *UserStore) IsFollower(userID, followerID uint) (bool, error) {
	var f model.Follow
	if err := us.db.Where("following_id = ? AND follower_id = ?", userID, followerID).Find(&f).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
