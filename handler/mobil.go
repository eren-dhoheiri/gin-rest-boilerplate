package handler

import (
	"go-gin-api/helper"
	"go-gin-api/mobil"
	"net/http"

	"github.com/gin-gonic/gin"
)


type mobilHandler struct {
	mobilService mobil.Service
}

func NewMobilHandler(mobilService mobil.Service) *mobilHandler {
	return &mobilHandler{mobilService}
}

func (h *mobilHandler) GetBrand(c *gin.Context) {
	brands, err := h.mobilService.GetBrand()
	if err != nil {
		response := helper.APIResponse("Error get brand", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of brand", http.StatusOK, "Success", mobil.FormatBrands(brands))
	c.JSON(http.StatusOK, response)
}

func (h *mobilHandler) GetBrandByID(c *gin.Context) {

	var input mobil.GetBrandDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	brandDetail, err := h.mobilService.GetBrandByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	response := helper.APIResponse("brand detail", http.StatusOK, "succes", mobil.FormatBrand(brandDetail))
	c.JSON(http.StatusOK, response)

}

func (h *mobilHandler) CreateBrand(c *gin.Context) {
	
	var input mobil.CreateBrandInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create mobil", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBrand, err := h.mobilService.CreateBrand(input)
	if err != nil {
		response := helper.APIResponse("Failed to create brand", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create brand", http.StatusOK, "Success", mobil.FormatBrand(newBrand))
	c.JSON(http.StatusOK, response)

}

func (h *mobilHandler) DeleteBrand(c *gin.Context) {

	var input mobil.GetBrandDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	brandDetail, err := h.mobilService.DeleteBrand(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete brand", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	response := helper.APIResponse("delete brand", http.StatusOK, "success", mobil.FormatBrand(brandDetail))
	c.JSON(http.StatusOK, response)

}

// 

func (h *mobilHandler) GetMobil(c *gin.Context) {
	mobils, err := h.mobilService.GetMobil()
	if err != nil {
		response := helper.APIResponse("Error get mobil", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of mobil", http.StatusOK, "Success", mobil.FormatMobils(mobils))
	c.JSON(http.StatusOK, response)
}

func (h *mobilHandler) GetMobilByID(c *gin.Context) {

	var input mobil.GetMobilDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of mobil", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mobilDetail, err := h.mobilService.GetMobilByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of mobil", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	response := helper.APIResponse("mobil detail", http.StatusOK, "succes", mobil.FormatMobil(mobilDetail))
	c.JSON(http.StatusOK, response)

}

func (h *mobilHandler) CreateMobil(c *gin.Context) {
	
	var input mobil.CreateMobilInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create mobil", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMobil, err := h.mobilService.CreateMobil(input)
	if err != nil {
		response := helper.APIResponse("Failed to create mobil", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create mobil", http.StatusOK, "Success", mobil.FormatMobilSave(newMobil))
	c.JSON(http.StatusOK, response)

}


func (h *mobilHandler) UpdateMobil(c *gin.Context) {
	
	var inputID mobil.GetMobilDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mobil", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData mobil.CreateMobilInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update mobil", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedMobil, err := h.mobilService.UpdateMobil(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update mobil", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update mobil", http.StatusOK, "Success", mobil.FormatMobil(updatedMobil))
	c.JSON(http.StatusOK, response)

}

func (h *mobilHandler) DeleteMobil(c *gin.Context) {

	var input mobil.GetMobilDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete mobil", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mobilDetail, err := h.mobilService.DeleteMobil(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete mobil", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	response := helper.APIResponse("delete mobil", http.StatusOK, "succes", mobil.FormatMobil(mobilDetail))
	c.JSON(http.StatusOK, response)

}