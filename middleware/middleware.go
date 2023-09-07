package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	SECRET = "secret"
)

func AuthValidate(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token required",
		})
		c.Abort()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		})
		c.Abort()
	}

}

//Dokumentasi
// Program yang Anda sertakan adalah contoh middleware dalam bahasa Go (Golang) yang berfungsi untuk memvalidasi token JWT (JSON Web Token) yang dikirimkan dalam header permintaan HTTP. Middleware ini digunakan dalam kerangka kerja web Gin untuk memeriksa apakah pengguna yang mengakses suatu endpoint memiliki token yang valid sebelum membiarkan mereka mengakses endpoint tersebut. Berikut adalah penjelasan dan dokumentasi dari middleware tersebut:

// ### Header Imports
// ```go
// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )
// ```

// Pada bagian ini, Anda mengimpor paket-paket yang diperlukan. Ini mencakup paket Gin untuk kerangka kerja web, paket JWT untuk memanipulasi token JWT, dan paket `fmt` yang digunakan untuk mencetak pesan.

// ### Konstanta
// ```go
// const (
// 	SECRET = "secret"
// )
// ```

// Anda mendefinisikan konstanta `SECRET` yang berisi kunci rahasia yang digunakan untuk menandatangani token JWT. Kunci ini harus sesuai dengan kunci yang digunakan pada saat pembuatan token.

// ### Fungsi `AuthValidate()`
// ```go
// func AuthValidate(c *gin.Context) {
// 	tokenString := c.Request.Header.Get("Authorization")

// 	if tokenString == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "token required",
// 		})
// 		c.Abort()
// 	}
// ```

// - Fungsi `AuthValidate()` adalah middleware yang akan dipanggil sebelum suatu permintaan HTTP mencapai handler endpoint yang sesuai.

// - Middleware ini pertama-tama mencoba untuk mengambil token JWT dari header permintaan HTTP dengan menggunakan `c.Request.Header.Get("Authorization")`. Jika tidak ada token yang ditemukan (token kosong), maka middleware mengirimkan respons JSON dengan status kode 400 (Bad Request) dan pesan "token required". Selanjutnya, middleware menghentikan pemrosesan permintaan dengan `c.Abort()`.

// ```go
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
// 			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
// 		}
// 		return []byte(SECRET), nil
// 	})
// ```

// - Selanjutnya, middleware ini mencoba untuk mem-parse (menguraikan) token JWT yang ditemukan dengan menggunakan `jwt.Parse()`. Dalam fungsi ini, Anda memeriksa metode penandatanganan token untuk memastikan bahwa token menggunakan metode HMAC-SHA256 (`jwt.SigningMethodHMAC`), yang sesuai dengan kunci rahasia yang digunakan saat pembuatan token. Jika metode tidak valid, middleware mengembalikan pesan kesalahan.

// - Jika metode valid, middleware ini mengembalikan kunci rahasia yang digunakan untuk memverifikasi tanda tangan token.

// ```go
// 	if token != nil && err == nil {
// 		fmt.Println("token verified")
// 		c.Next()
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "not authorized",
// 			"error":   err.Error(),
// 		})
// 		c.Abort()
// 	}
// ```

// - Terakhir, middleware memeriksa apakah token telah berhasil di-parse (tidak nil) dan tidak ada kesalahan saat parsing (`err == nil`). Jika token valid, maka middleware mencetak pesan "token verified" dan membiarkan permintaan melanjutkan ke handler endpoint berikutnya dengan `c.Next()`.

// - Jika token tidak valid, middleware mengirimkan respons JSON dengan status kode 401 (Unauthorized) dan pesan "not authorized". Middleware juga mencetak pesan kesalahan jika ada (`err.Error()`) dan menghentikan pemrosesan permintaan dengan `c.Abort()`.

// Middleware ini digunakan untuk melindungi endpoint yang memerlukan autentikasi dengan token JWT. Jika pengguna tidak memiliki token yang valid atau token tidak valid, mereka tidak diizinkan mengakses endpoint tersebut. Middleware serupa sering digunakan dalam pengembangan aplikasi web untuk melindungi sumber daya yang memerlukan otorisasi.
