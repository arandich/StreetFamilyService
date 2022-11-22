package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"streetfamily.com/model"
	"streetfamily.com/storage"
)

type Settings struct {
	Catalog []model.Catalog
}

func (h *Handler) GetCatalog(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	//var body model.Catalog
	//
	//if err := c.BindJSON(&body); err != nil {
	//	log.Fatal(err)
	//}
	//
	//nameQuery := body.Name + "%"

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM shop_name_catalog")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successful query")

	var values []interface{}

	var resp Settings

	for rows.Next() {

		values, err = rows.Values()
		if err != nil {
			log.Println("error while irrerating dataset")
		}

		item := model.Catalog{
			ID:          values[0].(int32),
			Category:    values[1].(int32),
			Name:        values[2].(string),
			Value:       values[3].(int32),
			Src:         values[4].(string),
			Weight:      values[5].(int32),
			Description: values[6].(string),
		}

		resp.Catalog = append(resp.Catalog, item)

	}

	if len(resp.Catalog) == 0 {
		c.JSON(http.StatusBadRequest, "catalog not found")
		pgpool.Close()
	} else {
		c.JSON(http.StatusOK, resp.Catalog)
	}

	if err := rows.Err(); err != nil {
		log.Panic(err)
	}

}

func (h *Handler) AddItem(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body model.Catalog

	if err := c.BindJSON(&body); err != nil {
		log.Panic(err)
		return
	}

	item := model.Catalog{
		ID:          0,
		Category:    body.Category,
		Name:        body.Name,
		Value:       body.Value,
		Src:         body.Src,
		Weight:      body.Weight,
		Description: body.Description,
	}

	rows, err := pgpool.Query(context.Background(), "INSERT INTO shop_name_catalog (category_id, name, value, src, weight, description) VALUES ($1,$2,$3,$4,$5,$6)",
		item.Category, item.Name, item.Value, item.Src, item.Weight, item.Description)
	if err != nil {
		log.Panic(err)
	}

	var resp Settings

	for rows.Next() {

	}

	resp.Catalog = append(resp.Catalog, item)

	if len(resp.Catalog) == 0 {
		c.JSON(http.StatusBadRequest, "items add error")
		pgpool.Close()
	} else {
		c.JSON(http.StatusOK, resp.Catalog)
	}

	if err := rows.Err(); err != nil {
		log.Panic(err)
		return
	}

}

func (h *Handler) DeleteItem(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body model.Catalog

	if err := c.BindJSON(&body); err != nil {
		log.Panic(err)
		return
	}

	item := model.Catalog{
		ID: body.ID,
	}

	rows, err := pgpool.Query(context.Background(), "DELETE FROM shop_name_catalog  WHERE id = ($1)", body.ID)
	if err != nil {
		log.Panic(err)
	}

	var resp Settings

	for rows.Next() {

	}

	resp.Catalog = append(resp.Catalog, item)

	if len(resp.Catalog) == 0 {
		c.JSON(http.StatusBadRequest, "items delete error")
		pgpool.Close()
	} else {
		c.JSON(http.StatusOK, "Item deleted")
	}

	if err := rows.Err(); err != nil {
		log.Panic(err)
		return
	}

}

func (h *Handler) FindItemById(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body model.Catalog

	if err := c.BindJSON(&body); err != nil {
		log.Panic(err)
		return
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM shop_name_catalog  WHERE id = ($1)", body.ID)
	if err != nil {
		log.Panic(err)
	}

	var values []interface{}

	var resp Settings

	for rows.Next() {

		values, err = rows.Values()
		if err != nil {
			log.Println("error while irrerating dataset")
		}

		item := model.Catalog{
			ID:          values[0].(int32),
			Category:    values[1].(int32),
			Name:        values[2].(string),
			Value:       values[3].(int32),
			Src:         values[4].(string),
			Weight:      values[5].(int32),
			Description: values[6].(string),
		}

		resp.Catalog = append(resp.Catalog, item)

	}

	if len(resp.Catalog) == 0 {
		c.JSON(http.StatusBadRequest, "item not found")
		pgpool.Close()
	} else {
		c.JSON(http.StatusOK, resp.Catalog)
	}

	if err := rows.Err(); err != nil {
		log.Panic(err)
	}

}

func (h *Handler) UpdateItemById(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body model.Catalog

	if err := c.BindJSON(&body); err != nil {
		log.Panic(err)
		return
	}

	item := model.Catalog{
		ID:          body.ID,
		Category:    body.Category,
		Name:        body.Name,
		Value:       body.Value,
		Src:         body.Src,
		Weight:      body.Weight,
		Description: body.Description,
	}
	query := fmt.Sprintf("UPDATE shop_name_catalog SET category_id = %d,name = '%s',value = %d,src = '%s',weight = %d,description = '%s' WHERE id = %d", item.Category, item.Name, item.Value, item.Src, item.Weight, item.Description, item.ID)
	rows, err := pgpool.Query(context.Background(), query)
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {

	}

	c.JSON(http.StatusOK, item)

	if err := rows.Err(); err != nil {
		log.Panic(err)
	}

}
