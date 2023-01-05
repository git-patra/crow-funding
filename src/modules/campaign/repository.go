package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := repo.db.Find(&campaigns).Preload("CampaignImages").Error

	return campaigns, err
}

func (repo *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := repo.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	return campaigns, err
}
