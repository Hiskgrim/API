package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EscalafonPersona struct {
	Estado           bool                       `orm:"column(estado)"`
	Observacion      string                     `orm:"column(observacion);null"`
	FechaRegistro    time.Time                  `orm:"column(fecha_registro);type(date)"`
	IdEscalafon      *Escalafon                 `orm:"column(id_escalafon);rel(fk)"`
	IdPersonaNatural *InformacionPersonaNatural `orm:"column(id_persona_natural);rel(fk)"`
	Id               int                        `orm:"column(id_escalafon_persona);pk"`
}

func (t *EscalafonPersona) TableName() string {
	return "escalafon_persona"
}

func init() {
	orm.RegisterModel(new(EscalafonPersona))
}

// AddEscalafonPersona insert a new EscalafonPersona into database and returns
// last inserted Id on success.
func AddEscalafonPersona(m *EscalafonPersona) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEscalafonPersonaById retrieves EscalafonPersona by Id. Returns error if
// Id doesn't exist
func GetEscalafonPersonaById(id int) (v *EscalafonPersona, err error) {
	o := orm.NewOrm()
	v = &EscalafonPersona{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEscalafonPersona retrieves all EscalafonPersona matches certain condition. Returns empty list if
// no records exist
func GetAllEscalafonPersona(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EscalafonPersona))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EscalafonPersona
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEscalafonPersona updates EscalafonPersona by Id and returns error if
// the record to be updated doesn't exist
func UpdateEscalafonPersonaById(m *EscalafonPersona) (err error) {
	o := orm.NewOrm()
	v := EscalafonPersona{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEscalafonPersona deletes EscalafonPersona by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEscalafonPersona(id int) (err error) {
	o := orm.NewOrm()
	v := EscalafonPersona{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EscalafonPersona{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
