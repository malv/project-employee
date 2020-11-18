package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project-employee/config"
	"strings"

	pb "project-employee/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var g *gorm.DB
var ReqToken string

func SetService(gDB *gorm.DB) {
	g = gDB
}

func catchError(e *error) {
	config.CatchError(e)
}

func MiddlewareCredential(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ReqToken = c.Request().Header.Get("Authorization")
		splitToken := strings.Split(ReqToken, "Bearer ")
		ReqToken = splitToken[1]
		fmt.Println("TOKEN:", ReqToken)
		err := CheckCredentialToken(ReqToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

func CheckCredentialToken(token string) error {
	res, err := config.Client.ValidateToken(config.Ctx,
		&pb.Token{Data: token})

	if err != nil {
		desc := strings.Split(err.Error(), "desc = ")
		err = errors.New(desc[1])
		log.Println("Error validate =>", err)
		return err
	}

	log.Println("Success validate =>", res)
	return nil
}

func GetRequest(uri string) (bodyResp []byte, e error) {
	var bearer = "Bearer " + ReqToken
	defer config.CatchError(&e)
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response =>", err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
