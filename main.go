package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simpel-app-auth/auth"
	"simpel-app-auth/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type GormStudent struct {
	Stud_id       uint64 `json:"stud_id" binding:"required"`
	Stud_name     string `json:"stud_name" binding:"required"`
	Stud_age      uint64 `json:"stud_age" binding:"required"`
	Stud_address  string `json:"stud_address" binding:"required"`
	Stud_phonenum string `json:"stud_phonenum" binding:"required"`
}

func postHandler(ctx *gin.Context, db *gorm.DB) {
	var data GormStudent

	// if ctx.Bind(&data) == nil {
	// 	_, err := db.Exec("insert into students values ($1,$2,$3,$4,$5)", data.Stud_id, data.Stud_name, data.Stud_age, data.Stud_address, data.Stud_phonenum)

	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"massage": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "created ok",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusBadRequest, gin.H{
	// 	"message": "error",
	// })

	//Dengan Gorm
	//Bind adalah metode yang digunakan dalam konteks HTTP yang bertujuan untuk mengambil dan memparsing data dari permintaan HTTP yang masuk.
	if ctx.Bind(&data) == nil { //data yang dikirim saat http tidak kosong
		db.Create(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"massage": "succes created",
			"data":    data,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"massage": "error",
	})

}

func getHandler(ctx *gin.Context, db *gorm.DB) {
	// var data []GormStudent
	// studId := ctx.Param("stud_id")                                             //ini dapet nya nanti saat kita masukin param di uri
	// rows, err := db.Query("select * from students where stud_id = $1", studId) //$1 itu kitamsukan nilai dari studId

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// rowTostuct(rows, &data)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"data":   data,
	// })

	//dengan gorm
	var data GormStudent

	studId := ctx.Param("stud_id")
	// id, _ := strconv.ParseUint(studId, 10, 64)

	// data := GormStudent{Stud_id: id}
	// Dalam contoh di atas, data akan memiliki nilai yang sama dengan nilai id
	if db.Find(&data, "stud_id=?", studId).RecordNotFound() {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data, //karena variabel data ini telah memiliki nilai baru setelah di prosen oleh find()
	})

}

func getAllHandler(ctx *gin.Context, db *gorm.DB) {
	// var data []GormStudent
	// rows, err := db.Query("select * from students")

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// rowTostuct(rows, &data)

	// if data == nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"massage": "data not found",
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"data":   data,
	// })

	//Getallstudent Dengan Gorm

	var data []GormStudent
	//kenapa harus pakai &data karena untuk daat fungsi find dijalankan dia bisa memproses dan merubah nilai
	//di variabel data diatasnya
	db.Find(&data) //untuk mencari /men get semua data yang ada di tabel di database
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})

}

func putHandler(c *gin.Context, db *gorm.DB) {
	// var data GormStudent
	// studId := c.Param("stud_id")

	// if c.Bind(&data) == nil {
	// 	_, err := db.Exec("update students set stud_name=$1 where stud_id=$2", data.Stud_name, studId)

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "success",
	// 		"massage": "succes update",
	// 	})
	// }

	//dengan gorm
	var data GormStudent

	studId := c.Param("stud_id")

	if db.Find(&data, "stud_id=?", studId).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"massage": "tidak ada data dengan id :" + studId,
		})
		return
	}

	reqStud := data

	if c.Bind(&reqStud) == nil { //jika nill artinya proses biding berhasil
		db.Model(&data).Where("stud_id=?", studId).Update(reqStud)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"data":    data,
			"massage": "data berhasil di update",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"massage": "error gagal binding",
	})

}

func delHandler(c *gin.Context, db *gorm.DB) {
	// studId := c.Param("stud_id")

	// _, err := db.Exec("delete from students where stud_id=$1", studId)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"massage": err.Error(),
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  "succes",
	// 	"massage": "delete success",
	// })
	var data GormStudent
	studId := c.Param("stud_id")

	db.Delete(&data, "stud_id=?", studId)
	c.JSON(http.StatusOK, gin.H{
		"massage": "delete sucess",
	})
}

func setupRouter() *gin.Engine {
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatal("Error load env")

	}
	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)

	}
	Migrate(db)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "succes",
		})
	})

	r.POST("/login", auth.LoginHandler)

	v1 := r.Group("v1")

	v1.POST("/student", func(ctx *gin.Context) {
		postHandler(ctx, db)
	})

	v1.GET("/student", middleware.AuthValidate, func(ctx *gin.Context) {
		getAllHandler(ctx, db)
	})

	v1.GET("/student/:stud_id", middleware.AuthValidate, func(ctx *gin.Context) {
		getHandler(ctx, db)
	})

	v1.PUT("/student/:stud_id",middleware.AuthValidate, func(ctx *gin.Context) {
		putHandler(ctx, db)
	})
	v1.DELETE("/student/:stud_id",middleware.AuthValidate, func(ctx *gin.Context) {
		delHandler(ctx, db)
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")

}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&GormStudent{})

	data := GormStudent{}
	if DB.Find(&data).RecordNotFound() {
		fmt.Println("=============run seeder user ==================")
		seederUser(DB)
	}

}

func seederUser(DB *gorm.DB) {
	data := GormStudent{
		Stud_id:       1,
		Stud_name:     "fahrel",
		Stud_age:      20,
		Stud_address:  "jakarta",
		Stud_phonenum: "09876543214",
	}

	DB.Create(&data)

}

// Penjelasan Kode:

// Kita mengimpor pustaka-pustaka yang diperlukan seperti github.com/gin-gonic/gin untuk framework Gin, github.com/jinzhu/gorm untuk ORM GORM, dan _ "github.com/lib/pq" untuk driver PostgreSQL.

// Kami mendefinisikan struktur data GormStudent yang mencerminkan entitas mahasiswa dengan atribut-atribut seperti stud_id, stud_name, stud_age, stud_address, dan stud_phonenum.

// Fungsi main adalah fungsi utama yang menjalankan aplikasi dan mengatur server web untuk mendengarkan pada port 8080.

// Fungsi setupRouter digunakan untuk mengonfigurasi router HTTP menggunakan Gin. Ini mendefinisikan rute-rute yang digunakan untuk menangani permintaan HTTP seperti POST, GET, PUT, dan DELETE.

// Fungsi Migrate digunakan untuk melakukan migrasi otomatis tabel ke database PostgreSQL dan menambahkan data awal jika tidak ada data mahasiswa dalam database.

// Fungsi seederUser digunakan untuk menambahkan data awal (seeder) ke dalam database. Dalam contoh ini, satu data mahasiswa ditambahkan.

// Fungsi postHandler menangani permintaan POST untuk menambahkan data mahasiswa baru ke dalam database.

// Fungsi getHandler menangani permintaan GET untuk mendapatkan data mahasiswa berdasarkan stud_id yang diberikan sebagai bagian dari URL.

// Fungsi getAllHandler menangani permintaan GET untuk mendapatkan semua data mahasiswa dari database.

// Fungsi putHandler menangani permintaan PUT untuk memperbarui data mahasiswa berdasarkan stud_id yang diberikan sebagai bagian dari URL.

// Fungsi delHandler menangani permintaan DELETE untuk menghapus data mahasiswa berdasarkan stud_id yang diberikan sebagai bagian dari URL.

// Catatan Penting:
// Pastikan Anda sudah memiliki PostgreSQL yang dijalankan dengan benar dan telah mengonfigurasi sesuai dengan koneksi yang Anda tentukan. Selain itu, pastikan semua pustaka yang digunakan telah diinstal.

// Dengan
