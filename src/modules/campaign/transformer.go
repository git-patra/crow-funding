package campaign

import "time"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
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
