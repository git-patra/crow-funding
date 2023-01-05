package campaign

import (
	"strings"
	"time"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	return CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		CreatedAt:        campaign.CreatedAt,
		UpdatedAt:        campaign.UpdatedAt,
	}
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	results := []CampaignFormatter{}

	for _, campaign := range campaigns {
		results = append(results, FormatCampaign(campaign))
	}

	return results
}

type CampaignDetailFormatter struct {
	ID               int                             `json:"id"`
	UserID           int                             `json:"user_id"`
	Name             string                          `json:"name"`
	ShortDescription string                          `json:"short_description"`
	Description      string                          `json:"description"`
	ImageURL         string                          `json:"image_url"`
	GoalAmount       int                             `json:"goal_amount"`
	CurrentAmount    int                             `json:"current_amount"`
	Slug             string                          `json:"slug"`
	Perks            []string                        `json:"perks"`
	User             CampaignDetailUserFormatter     `json:"user"`
	Images           []CampaignDetailImagesFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignDetailImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	perks := []string{}
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	images := []CampaignDetailImagesFormatter{}
	for _, image := range campaign.CampaignImages {
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		images = append(images, CampaignDetailImagesFormatter{
			ImageURL:  image.FileName,
			IsPrimary: isPrimary,
		})
	}

	return CampaignDetailFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		Perks:            perks,
		Images:           images,
		User: CampaignDetailUserFormatter{
			Name:     campaign.User.Name,
			ImageURL: campaign.User.AvatarFileName,
		},
	}
}
