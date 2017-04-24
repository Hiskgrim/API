package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
