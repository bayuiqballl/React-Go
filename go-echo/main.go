package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bayuiqballl/go-echo/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func app(e *echo.Echo, store models.SiswaStore) {
	e.GET("/siswa", func(c echo.Context) error {
		siswas := store.All()
		return c.JSON(http.StatusOK, siswas)
	})

	e.GET("/siswa/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		siswas := store.Find(id)
		return c.JSON(http.StatusOK, siswas)
	})

	e.POST("/siswa", func(c echo.Context) error {
		nama := c.FormValue("nama")
		nisn := c.FormValue("nisn")
		pendidikan := c.FormValue("pendidikan")

		siswas, _ := models.CreateSiswa(nama, nisn, pendidikan)
		store.Save(siswas)

		return c.JSON(http.StatusOK, siswas)
	})

	e.PUT("/siswa/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		siswa := store.Find(id)
		siswa.Nama = c.FormValue("nama")
		siswa.Nisn = c.FormValue("nisn")
		siswa.Pendidikan = c.FormValue("pendidikan")

		store.Update(siswa)

		return c.JSON(http.StatusOK, siswa)
	})

	e.DELETE("/siswa/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		siswa := store.Find(id)
		store.Delete(siswa)
		return c.JSON(http.StatusOK, siswa)
	})

}

func main() {
	godotenv.Load()
	store := models.NewSiswaStoreMysql()
	e := echo.New()

	app(e, store)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
