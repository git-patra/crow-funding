package handlers

import (
	"CrowFundingV2/src/auth"
	"CrowFundingV2/src/helper"
	"CrowFundingV2/src/modules/campaign"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	campaignService campaign.Service
	authService     auth.Service
}

func NewCampaignHandler(service campaign.Service, authService auth.Service) *campaignHandler {
	return &campaignHandler{service, authService}
}

// api get list
func (h campaignHandler) GetListCampaign(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.FindCampaigns(userID)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Get data failed!", "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "List of campaigns.", "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetDetailCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Get data failed!", "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.FindCampaignByID(input)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Get data failed!", "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "Campaign Detail.", "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}
