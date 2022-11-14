package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"streetfamily.com/storage"
)

type userReq struct {
	ID int32 `json:"id"`
}

func (h *Handler) Login(c *gin.Context) {

	log.Printf("connecting to postgres...")
	pgpool, err := storage.ConnectToPostgres(context.Background(), "postgres://admin:password123@postgres:5432/golang_postgres")
	log.Println("Success connect to pool")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("successfully connected to postgres")
	defer pgpool.Close()

	var body userReq

	if err := c.BindJSON(&body); err != nil {
		return
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM shop WHERE id=($1)", body.ID)
	if err != nil {
		log.Fatal(err)
	}

	var values []interface{}
	var resp userReq

	for rows.Next() {

		values, err = rows.Values()
		if err != nil {
			log.Println("error while irrerating dataset")
		}

		resp.ID = values[0].(int32)

		c.JSON(http.StatusOK, gin.H{
			"id": resp.ID,
		})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if resp.ID == 0 {
		c.JSON(http.StatusBadRequest, "user not found")
	}

}
