// @Description: 初始化配置文件
// @Author: wucongxing
// @Date:2021/12/22 19:28

package Config

var (
	C Server
)

// Server
// @Description: 配置数据
type Server struct {
	Port  int    `json:"port"`
	Env   string `json:"env"`
	Mysql Mysql
}

func init() {

}
