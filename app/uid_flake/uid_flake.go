package uid_flake

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type req struct {
	Uid     int     `json:"uid"`
	UidList []int64 `json:"uid_list"`
	//UinList []int64 `json:"uinList"`
}

type UidFlake struct {
}

func (u *UidFlake) Router() string {
	return "/api/uid_flake"
}

func (u *UidFlake) Handler(c *gin.Context) {
	fmt.Println(c.Params.Get("uid"))
	var request req
	err := c.Bind(&request)
	if err != nil {
		fmt.Printf("bind failed! err:%+v", err)
		return
	}
	fmt.Println(request)
	fmt.Println(c.Request.Body)
}
