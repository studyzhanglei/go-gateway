package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/pb/common"
	"gin-vue-admin/pb/search"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)



func Grpc2(c *gin.Context) {
	var G  request.Grpc
	_ = c.ShouldBindJSON(&G)
	grpc, err := global.GRPC_POOL.Get(c)
	if err != nil {
		fmt.Println(err, 3333)
	}

	defer grpc.Close()

	client := search.NewSearchServiceClient(grpc.ClientConn)
	stream, err := client.Search(c)
	if err != nil {
		fmt.Println(err, 888888)
	}

	fmt.Println(G.Username, "username")
	utils.GetTraceLog(c).Info(G.Username)

	stream.Send(&search.SearchRequest{
		Header: &common.CommonHeader{
			TraceId: c.Value("trace-id").(string),
		},
		Request: G.Username,
	})

	result, err := stream.Recv()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println(result.Status)

	status := result.GetStatus()

	if status.Code != 0 {
		response.Result(int(status.Code), map[string]interface{}{}, status.Msg, c)
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
