package auth

import (
	"net/http"
	"simpel-app-auth/models"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

const (
	USER     = "admin"
	PASSWORD = "cungkring"
	SECRET   = "secret"
)

func LoginHandler(c *gin.Context) {
	var user models.Credential

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	if user.Username != USER { //jika user yang didapat dari req http tidak sama dengan user yang kita telah tetapkan di const
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User Invalid",
		})
		return
	} else {
		if user.Password != PASSWORD {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "password invalid",
			})
			return
		}
	}

	//token
	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(), //token ini bertahan selama 1 menit setelah 1 menit akan di drop
		Issuer:    "test",
		IssuedAt:  time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(SECRET))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})

}

//Dokumentasi
// Program yang Anda sertakan adalah bagian dari sebuah aplikasi yang berfungsi sebagai mekanisme autentikasi (login) untuk pengguna. Program ini menggunakan kerangka kerja web Gin untuk membuat endpoint HTTP yang memproses permintaan login dan mengeluarkan token JWT (JSON Web Token) sebagai hasilnya. Berikut adalah penjelasan dan dokumentasi dari program tersebut:

// ### Header Imports
// ```go
// import (
// 	"net/http"
// 	"simpel-app-auth/models"
// 	"time"

// 	jwt "github.com/golang-jwt/jwt/v4"

// 	"github.com/gin-gonic/gin"
// )
// ```
// Pada bagian ini, Anda mengimpor paket-paket yang diperlukan. Ini termasuk paket `net/http` untuk komunikasi HTTP, paket-paket terkait JWT, dan paket Gin untuk kerangka kerja web.

// ### Konstanta
// ```go
// const (
// 	USER     = "admin"
// 	PASSWORD = "cungkring"
// 	SECRET   = "secret"
// )
// ```
// Anda mendefinisikan tiga konstanta:
// - `USER` dan `PASSWORD` adalah nama pengguna (username) dan kata sandi (password) yang digunakan untuk autentikasi. Dalam contoh ini, nama pengguna dan kata sandi tetap (hardcoded).
// - `SECRET` adalah kunci rahasia yang digunakan untuk menandatangani token JWT.

// ### Fungsi `LoginHandler()`
// ```go
// func LoginHandler(c *gin.Context) {
// 	var user models.Credential

// 	err := c.Bind(&user)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "bad request",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// ```
// - Fungsi `LoginHandler()` adalah handler yang akan dipanggil ketika permintaan HTTP POST masuk ke endpoint login. Fungsi ini bertanggung jawab untuk mengautentikasi pengguna dan mengeluarkan token JWT.

// - Anda menggunakan `c.Bind(&user)` untuk mengikat data dari permintaan HTTP ke variabel `user`. Data ini diharapkan berisi informasi pengguna seperti nama pengguna (username) dan kata sandi (password).

// - Jika terjadi kesalahan dalam pengikatan data, misalnya jika data tidak valid, Anda mengirimkan respons JSON dengan status kode 400 (Bad Request) dan pesan kesalahan.

// ```go
// 	if user.Username != USER { //jika user yang didapat dari req http tidak sama dengan user yang kita telah tetapkan di const
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "User Invalid",
// 		})
// 		return
// 	} else {
// 		if user.Password != PASSWORD {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "password invalid",
// 			})
// 			return
// 		}
// 	}
// ```
// - Anda memeriksa apakah nama pengguna dan kata sandi yang diberikan oleh pengguna sesuai dengan nilai konstan `USER` dan `PASSWORD`. Jika tidak sesuai, Anda mengirimkan respons JSON dengan status kode 401 (Unauthorized) dan pesan kesalahan yang sesuai.

// ```go
// 	//token
// 	claim := jwt.StandardClaims{
// 		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(), //token ini bertahan selama 1 menit setelah

// Maaf atas ketidaknyamanannya. Mari kita lanjutkan penjelasan dari kode tersebut:

// ```go
// 		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(), // Token ini bertahan selama 1 menit setelah 1 menit akan di-drop
// 		Issuer:    "test",
// 		IssuedAt:  time.Now().Unix(),
// 	}

// 	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 	token, err := sign.SignedString([]byte(SECRET))

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		c.Abort()
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "success",
// 		"token":   token,
// 	})
// }
// ```

// - Pada bagian ini, Anda membuat token JWT dengan mengisi informasi klaim (claim) seperti waktu kedaluwarsa (expiresAt), penerbit (issuer), dan waktu diterbitkan (issuedAt). Token ini akan berlaku selama 1 menit (`time.Minute * 1`) setelah itu akan di-drop.

// - Anda menggunakan `jwt.NewWithClaims()` untuk membuat token dengan metode penandatanganan HMAC-SHA256 (`jwt.SigningMethodHS256`) dan klaim yang telah Anda siapkan sebelumnya.

// - Kemudian, Anda menandatangani token menggunakan `sign.SignedString([]byte(SECRET))`, di mana `SECRET` adalah kunci rahasia yang telah Anda tetapkan. Jika terjadi kesalahan dalam pembuatan token, Anda mengirimkan respons JSON dengan status kode 500 (Internal Server Error) dan pesan kesalahan.

// - Jika berhasil, Anda mengirimkan respons JSON dengan status kode 200 (OK) yang berisi pesan "success" dan token JWT yang telah dibuat.

// Ini adalah contoh sederhana dari sebuah handler dalam sebuah aplikasi web yang melakukan autentikasi pengguna dengan menggunakan JWT. Setelah pengguna berhasil diotentikasi, aplikasi menghasilkan token JWT yang dapat digunakan untuk otentikasi pada permintaan-permintaan selanjutnya. Dalam pengembangan nyata, ini hanya merupakan langkah awal dan biasanya akan ada lebih banyak logika dan lapisan keamanan yang diterapkan.
