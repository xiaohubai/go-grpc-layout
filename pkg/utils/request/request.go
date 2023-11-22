package request

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(c *gin.Context, req any) error {
	err := c.ShouldBindJSON(req)
	if err != nil {
		return err
	}
	refReq := reflect.ValueOf(req)
	res := refReq.MethodByName("ValidateAll").Call(nil)[0]
	if !res.IsNil() {
		return fmt.Errorf("%s", res)
	}

	return nil
}
