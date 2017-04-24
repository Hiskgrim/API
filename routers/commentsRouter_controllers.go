package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DedicacionPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
