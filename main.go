package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err =
		gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/anekazoo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
	db.AutoMigrate(&animal{})
}

type (
	animal struct {
		gorm.Model
		Name  string `json:"name"`
		Class string `json:"class"`
		Legs  int    `json:"legs"`
	}
	transformedAnimal struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Class string `json:"class"`
		Legs  int    `json:"legs"`
	}
)

func cretedAnimal(c *gin.Context) {
	var std transformedAnimal
	var model animal
	c.Bind(&std)
	validasi := validatorCreated(std)
	model = transferVoToModel(std)
	if validasi != "" {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": validasi})
	} else {
		db.Create(&model)
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
	}
}

func fetchAllAnimal(c *gin.Context) {
	var model []animal
	var vo []transformedAnimal

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Not Found"})
	}

	for _, item := range model {
		vo = append(vo, transferModelToVo(item))
	}
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": vo})
}

func fetchSingleAnimal(c *gin.Context) {
	var model animal
	var vo transformedAnimal

	modelID := c.Param("id")
	db.Find(&model, modelID)

	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Not Found"})
	}
	vo = transferModelToVo(model)
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": vo})
}

func updateAnimal(c *gin.Context) {
	var model animal
	var vo transformedAnimal
	modelID := c.Param("id")
	db.First(&model, modelID)

	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "No Data"})
	}
	c.Bind(&vo)

	validasi := validatorCreated(vo)
	if validasi != "" {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": validasi})
	} else {
		db.Model(&model).Update(transferVoToModel(vo))
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
	}
}

func deleteAnimal(c *gin.Context) {
	var model animal
	modelID := c.Param("id")

	db.First(&model, modelID)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data not found"})
	}
	db.Delete(model)
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": "Data has been successfully deleted"})
}

func transferModelToVo(model animal) transformedAnimal {
	var vo transformedAnimal
	vo = transformedAnimal{
		ID:    model.ID,
		Name:  model.Name,
		Class: model.Class,
		Legs:  model.Legs,
	}
	return vo
}

func transferVoToModel(vo transformedAnimal) animal {
	var model animal
	model = animal{
		Name:  vo.Name,
		Class: vo.Class,
		Legs:  vo.Legs,
	}
	return model
}

func validatorCreated(vo transformedAnimal) string {

	var kosong string = "Can not be empty"

	if vo.Name == "" {
		return "Nama" + kosong
	}

	if vo.Class == "" {
		return "Class" + kosong
	}

	if vo.Legs == 0 {
		return "Legs" + kosong
	}

	return ""
}

func main() {

	router := gin.Default()
	v1 := router.Group("/v1/animal")
	{
		v1.POST("", cretedAnimal)
		v1.GET("", fetchAllAnimal)
		v1.GET("/:id", fetchSingleAnimal)
		v1.PUT("/:id", updateAnimal)
		v1.DELETE("/:id", deleteAnimal)
	}
	router.Run(":8080")
}
