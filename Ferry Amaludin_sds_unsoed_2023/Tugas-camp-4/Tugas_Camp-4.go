//Ferry Amaludin
//menginstal paket "gin" terlebih dahulu dengan menggunakan perintah:
//go get -u github.com/gin-gonic/gin
//setelah itu kode dibawah dijanakan terlebih dahulu dengan menekan crl+alt+n jika menggunakan vscode.

//Buka terminal atau command prompt pada komputer Anda.
//Salin perintah cURL yang dibawah dan copy ke command prompt.
//kemudian jalankan perintah tersebut diwah ini ya:

//curl -X POST -H "Content-Type: application/json" -d '{
//    "jari-jari-lingkaran": 5,
//    "sisi-persegi": 4,
//    "alas-segitiga": 6,
//    "tinggi-segitiga": 8
//}' http://localhost:8080/hitung

//

package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/hitung", func(c *gin.Context) {
		var inputData struct {
			JariJariLingkaran float64 `json:"jari-jari-lingkaran"`
			SisiPersegi       float64 `json:"sisi-persegi"`
			AlasSegitiga      float64 `json:"alas-segitiga"`
			TinggiSegitiga    float64 `json:"tinggi-segitiga"`
		}

		if err := c.ShouldBindJSON(&inputData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var hasil struct {
			LuasLingkaran     float64 `json:"luas-lingkaran"`
			KelilingLingkaran float64 `json:"keliling-lingkaran"`
			LuasPersegi       float64 `json:"luas-persegi"`
			KelilingPersegi   float64 `json:"keliling-persegi"`
			LuasSegitiga      float64 `json:"luas-segitiga"`
			KelilingSegitiga  float64 `json:"keliling-segitiga"`
		}

		// Menghitung luas dan keliling lingkaran
		if inputData.JariJariLingkaran > 0 {
			hasil.LuasLingkaran = math.Pi * inputData.JariJariLingkaran * inputData.JariJariLingkaran
			hasil.KelilingLingkaran = 2 * math.Pi * inputData.JariJariLingkaran
		}

		// Menghitung luas dan keliling persegi
		if inputData.SisiPersegi > 0 {
			hasil.LuasPersegi = inputData.SisiPersegi * inputData.SisiPersegi
			hasil.KelilingPersegi = 4 * inputData.SisiPersegi
		}

		// Menghitung luas dan keliling segitiga
		if inputData.AlasSegitiga > 0 && inputData.TinggiSegitiga > 0 {
			hasil.LuasSegitiga = 0.5 * inputData.AlasSegitiga * inputData.TinggiSegitiga
			sisiSegitiga := math.Sqrt(math.Pow(inputData.AlasSegitiga, 2) + math.Pow(inputData.TinggiSegitiga, 2))
			hasil.KelilingSegitiga = inputData.AlasSegitiga + inputData.TinggiSegitiga + sisiSegitiga
		}

		c.JSON(http.StatusOK, hasil)
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
