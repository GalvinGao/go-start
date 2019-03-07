package main

import (
	"net/http"
	"database/sql"
	
	"github.com/labstack/gommon/log"
	_	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type BaseResponse struct {
	Success	bool	`json:"success"`
	Description	string	`json:"description"`
	Version	string	`json:"version"`
}

type BoolResponse struct {
	Success	bool	`json:"success"`
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	db, err := sql.Open("mysql", "uv:U7UiZCK5IYZsNwaF@/uv_pages_dev")
	if err != nil {
		e.Logger.Fatal("MySQL Connect Error: " + err.Error())
	}
	
	err = db.Ping()
	if err != nil {
		e.Logger.Fatal("MySQL Connect Error: " + err.Error())
	}
	
	e.GET("/", func (c echo.Context) error {
		return c.JSON(http.StatusOK, &BaseResponse{
			Success: true,
			Description: "Universal Village Page CMS",
			Version: "0.1.0",
		})
	})
	
	/* e.GET("/check-db", func (c echo.Context) error {
		ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
		if err := db.PingContext(ctx); err != nil {
			e.Logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, &BoolResponse{
				Success: false,
			})
		}
		return c.JSON(http.StatusOK, &BoolResponse{
			Success: true,
		})
	}) */
	
	// Get the post contents; scope:guest
	e.GET("/:id", func (c echo.Context) error {
		// Parse the params
		paramid := c.Param("id")
		highlight := c.QueryParam("highlight")
		
		var (
			id	uint64
			created	uint8
			modified	uint8
			title	string
			content	string
		)
		
		stmt, err := db.Prepare("SELECT * FROM pages WHERE id = ? LIMIT 1")
		if err != nil {
			e.Logger.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(paramid)
		if err != nil {
			e.Logger.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&id, &created, &modified, &title, &content); err != nil {
				e.Logger.Fatal(err)
			}
		}
		if err = rows.Err(); err != nil {
			e.Logger.Fatal(err)
		}
		return c.String(http.StatusNotFound, "NotImplemented yet. params:id:"+paramid+",highlight:"+highlight)
	})
	
	// Admin group
	admin := e.Group("/admin")
	
	// Get the edit post contents UI page; scope:admin; type:html
	admin.GET("/:id", func (c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusNotFound, "NotImplemented yet. params:id:"+id)
	})
	
	// Post a post update; scope:admin; type:api
	admin.POST("/:id", func (c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusNotFound, "NotImplemented yet. params:id:"+id)
	})
	
	// Put a new post; scope:admin; type:api
	admin.PUT("/:id", func (c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusNotFound, "NotImplemented yet. params:id:"+id)
	})
	
	// Delete a post; scope:admin; type:api
	admin.DELETE("/:id", func (c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusNotFound, "NotImplemented yet. params:id:"+id)
	})
	
	e.Logger.Fatal(e.Start(":8087"))
}

