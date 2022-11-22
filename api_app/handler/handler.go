package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"streetfamily.com/model"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService    model.UserService
	DB             *pgxpool.Pool
	CompanyService model.CompanyService
	CatalogService model.CatalogService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R              *gin.Engine
	UserService    model.UserService
	CompanyService model.CompanyService
	CatalogService model.CatalogService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:    c.UserService,
		CompanyService: c.CompanyService,
		CatalogService: c.CatalogService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group("/api/")

	g.POST("/login", h.Login)
	g.POST("/companies", h.CompaniesSearch)
	g.POST("/catalog", h.GetCatalog)
	g.POST("/catalog/add", h.AddItem)
	g.POST("/catalog/delete", h.DeleteItem)
	g.POST("/catalog/find-by-id", h.FindItemById)
	g.PUT("/catalog/update-by-id", h.UpdateItemById)

}
