package handler

import (
	"middleware/model"
	"middleware/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHanlder interface {
	CreateData(c *gin.Context)
	GetData(c *gin.Context)
	GetDataById(c *gin.Context)
	UpdateData(c *gin.Context)
	DeleteData(c *gin.Context)
}

type handler struct {
	Service service.IService
}

func NewHandler(service service.IService) IHanlder {
	return &handler{
		Service: service,
	}
}

func (h *handler) CreateData(c *gin.Context) {
	var data model.Data

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	err = h.Service.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"message": "Created new data"})

}

func (h *handler) GetData(c *gin.Context) {
	datas := h.Service.Get()

	c.JSON(http.StatusOK, datas)
}

// DeleteData implements IHanlder.
func (*handler) DeleteData(c *gin.Context) {
	panic("unimplemented")
}

// GetDataById implements IHanlder.
func (*handler) GetDataById(c *gin.Context) {
	panic("unimplemented")
}

// UpdateData implements IHanlder.
func (*handler) UpdateData(c *gin.Context) {
	panic("unimplemented")
}
