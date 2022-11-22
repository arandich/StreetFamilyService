package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"streetfamily.com/model"
	"streetfamily.com/storage"
)

func (h *Handler) CompaniesSearch(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body model.CompanySearch

	if err := c.BindJSON(&body); err != nil {
		return
	}

	var nameQuery string
	if body.Name != "" {
		nameQuery = body.Name + "%"
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM shop WHERE shop_name LIKE  LOWER($1)", nameQuery)
	if err != nil {
		log.Panic(err)
		return
	}

	var values []interface{}
	var resp model.User
	respAr := map[int32]string{}

	for rows.Next() {

		values, err = rows.Values()
		if err != nil {
			log.Panic("error while irrerating dataset")
			return
		}

		resp.ID = values[0].(int32)
		resp.Name = values[1].(string)

		respAr[resp.ID] = resp.Name

	}

	if len(respAr) == 0 {
		c.JSON(http.StatusBadRequest, "companies not found")
		pgpool.Close()
	} else {
		c.JSON(http.StatusOK, respAr)
	}

	if err := rows.Err(); err != nil {
		log.Panic(err)
	}

}
