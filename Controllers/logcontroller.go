// controllers/books.go

package controllers

import (
	"fmt"
	"net/http"
	"os"

	models "DT/Models"
	"DT/Util"
	utl "DT/Util"
	b64 "encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Listerror ...
func Listerror(c *gin.Context) {
	// var result []models.Angel_Client
	// models.DB.Find(&logger)
	// models.DB.Raw("Admin_Get_Client_Data").Scan(&logger)
	// result := map[string]interface{}{}
	// models.DB.Raw("SELECT * FROM City_Master").Scan(&logger)
	// models.DB.Raw("Admin_Get_Client_Data @Option = 'Admin_ClientStatus_Cnt_Grid', @ErrorMsg = '@ErrorM OUTPUT'").Scan(&result)
	// rows, err := models.DB.Raw("SELECT * FROM City_Master").Rows()
	// rows, err := models.DB.Raw("Admin_Get_Client_Data @Option = 'Admin_ClientStatus_Cnt_Grid', @ErrorMsg = '@ErrorM OUTPUT'").Rows()
	email := "nilayy1234@gmail.com"
	mobile := "9739693088"
	query := fmt.Sprintf("usp_add_guest_user_session_26082020 @emailID = '%v', @mobileNum =  '%v'", email, mobile)
	fmt.Println("Query:", query)
	rows, err := models.DB.Raw(query).Rows()
	// rows, err := models.DB.Raw("usp_add_guest_user_session_26082020 @emailID = 'nilayy1234@gmail.com', @mobileNum = '9739693088'").Rows()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	data, err := utl.Jsonify(rows)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Webcheck ...
func Webcheck(c *gin.Context) {
	requestdata, _ := c.Get("data")
	requestIv, _ := c.Get("iv")
	ciphertext := fmt.Sprint(requestdata)
	iv := fmt.Sprint(requestIv)
	data, err := Util.Decryptnew(ciphertext, iv)
	// data, err := Util.Decrypt(ciphertext, iv)
	// text := fmt.Sprint(data)
	// response := Util.Base64encode(data, iv)
	empData, err := json.Marshal(data)
	sDec := b64.StdEncoding.EncodeToString(empData)
	// encode := fmt.Sprintf(`%v`, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sDec, "data1": data})
}

//Webcheck2 ...
func Webcheck2(c *gin.Context) {
	SecretKey := os.Getenv("SECRET_KEY")
	key := Util.GetHashSha256(SecretKey)
	s := "mmmI am string"
	i := "kkkkkkkk"
	encode := fmt.Sprintf(`{"Data":"%v","IV":"%v"}`, s, i)
	data, iv := Util.Encryptnew(key, encode)
	response := Util.Base64encode(data, iv)
	c.JSON(http.StatusOK, gin.H{"data": response, "iv": iv, "data1": data})
}
