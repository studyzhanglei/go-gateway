package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/studyzhanglei/grpc-proto/pb/search"
	"go-gateway/global"
	"go-gateway/model/request"
	"go-gateway/model/response"
	"go-gateway/utils"
)


func Grpc2(c *gin.Context) {
	defer utils.ErrorHandle(c)

	var G  request.Grpc
	_ = c.ShouldBindJSON(&G)
	grpc, err := global.GRPC_POOL.Get(c)
	if err != nil {
		return
	}

	defer grpc.Close()

	client := search.NewSearchServiceClient(grpc.ClientConn)
	stream, err := client.Search(c)
	if err != nil {
		return
	}

	utils.GetTraceLog(c).Info(G.Username)
	stream.Send(&search.SearchRequest{
		Header: utils.GetGrpcHeader(c),
		Request: G.Username,
	})

	result, err := stream.Recv()
	fmt.Println(err, 8888888)
	if err != nil {
		return
	}

	fmt.Println(result.Status)

	status := result.GetStatus()

	if status.Code != 0 {
		return
	}

	type Data struct {
		Name string `json:"name"`
	}

	response.OkWithDetailed(&Data{Name:result.GetResponse()}, "success", c)

}

func Grpc3(c *gin.Context) {
	var G request.Grpc
	_ = c.ShouldBindJSON(&G)
}


func Login(c *gin.Context) {
	if claims, exists := c.Get("claims"); exists {
		fmt.Println(claims.(*request.MyClaims))
	}

	utils.GetTraceLog(c)
}

func Index(c *gin.Context) {
	if claims, exists := c.Get("claims"); exists {
		fmt.Println(claims.(*request.MyClaims))
	}

	response.OkWithMessage("hello", c)
}


func UserInfo(c *gin.Context) {
	//defer utils.ErrorHandle(c)
	var G  request.Grpc
	_ = c.ShouldBindJSON(&G)
	grpc, err := global.GRPC_POOL.Get(c)
	if err != nil {
		return
	}

	defer grpc.Close()
	client := search.NewSearchServiceClient(grpc.ClientConn)
	stream, err := client.GetUserInfo(c)
	if err != nil {
		return
	}


	utils.GetTraceLog(c).Info(G.Username)
	stream.Send(&search.UserInfoRequest{
		Header: utils.GetGrpcHeader(c),
		Uid: uint64(G.Uid),
	})

	result,err := stream.Recv()

	if err != nil {
		fmt.Println(err)
		return
	}

	status := result.GetStatus()

	if status.Code != 0 {
		response.Result(int(status.Code), map[string]interface{}{}, status.Msg, c)
		return
	}

	type res struct{
		Name string
		Uid int
	}

	response.OkWithData(&res{
		Name: result.Username,
		Uid: int(result.Ud),
	}, c)
}