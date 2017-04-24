// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cdve_api_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/contrato_docente",
			beego.NSInclude(
				&controllers.ContratoDocenteController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/pais",
			beego.NSInclude(
				&controllers.PaisController{},
			),
		),

		beego.NSNamespace("/punto_salarial",
			beego.NSInclude(
				&controllers.PuntoSalarialController{},
			),
		),

		beego.NSNamespace("/caja_compensacion",
			beego.NSInclude(
				&controllers.CajaCompensacionController{},
			),
		),

		beego.NSNamespace("/eps",
			beego.NSInclude(
				&controllers.EpsController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/fondo_pension",
			beego.NSInclude(
				&controllers.FondoPensionController{},
			),
		),

		beego.NSNamespace("/escalafon",
			beego.NSInclude(
				&controllers.EscalafonController{},
			),
		),

		beego.NSNamespace("/arl",
			beego.NSInclude(
				&controllers.ArlController{},
			),
		),

		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/contrato_estado",
			beego.NSInclude(
				&controllers.ContratoEstadoController{},
			),
		),

		beego.NSNamespace("/lugar_ejecucion",
			beego.NSInclude(
				&controllers.LugarEjecucionController{},
			),
		),

		beego.NSNamespace("/parametro_entidad",
			beego.NSInclude(
				&controllers.ParametroEntidadController{},
			),
		),

		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),

		beego.NSNamespace("/escalafon_persona",
			beego.NSInclude(
				&controllers.EscalafonPersonaController{},
			),
		),

		beego.NSNamespace("/supervisor_contrato",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),

		beego.NSNamespace("/dedicacion_persona",
			beego.NSInclude(
				&controllers.DedicacionPersonaController{},
			),
		),

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),
		beego.NSNamespace("/contratado",
			beego.NSInclude(
				&controllers.ContratadoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
