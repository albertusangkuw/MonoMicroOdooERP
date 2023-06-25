package service

import (
	"fmt"
	"service-17/model"
	"service-17/utils"

	odoo "github.com/ausuwardi/godooj"
)

// Call Mono.
var modelID int

func SearchModelID(name string) {
	c := utils.CallOdoo()
	if !c.IsValid() {
		fmt.Println("Invalid !!!!")
		return
	}
	m, _ := c.Search("ir.model", odoo.List{
		odoo.List{"model", "=", name},
	})
	fmt.Println(m)
	modelID = m[0]
}

// Dipetakan dari groupID
var cacheACL map[int]model.AccessControlList

func UpdateACL() {
	client := utils.CallOdoo()
	if !client.IsValid() {
		fmt.Println("no Odoo Server to Update ACL")
		return
	}
	resRPC, err := client.SearchRead(
		"ir.model.access",
		odoo.List{
			odoo.List{"model_id", "=", modelID},
		}, []string{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	aclData := map[int]model.AccessControlList{}
	for _, rec := range resRPC {
		id, _ := odoo.IntField(rec, "id")
		groupID, _ := odoo.Many2oneField(rec, "group_id")
		modelID, _ := odoo.Many2oneField(rec, "model_id")
		read, _ := BooleanField(rec, "perm_read")
		write, _ := BooleanField(rec, "perm_write")
		create, _ := BooleanField(rec, "perm_create")
		unlink, _ := BooleanField(rec, "perm_unlink")

		aclData[groupID.ID] = model.AccessControlList{
			IrModelAccessID: id,
			GroupID:         groupID.ID,
			ModelID:         modelID.ID,
			Read:            read,
			Write:           write,
			Create:          create,
			Unlink:          unlink,
		}

	}
	// Simpan di cache
	fmt.Println(aclData)
	cacheACL = aclData
}

func IsAllowedByACL(groupID []int, context string) (bool, int) {
	if cacheACL == nil {
		fmt.Println("Cache ACL is empty !!!")
		return true, 0
	}
	for _, g := range groupID {
		acl, ok := cacheACL[g]
		if ok {
			switch context {
			case "c":
				return acl.Create, g
			case "w":
				return acl.Write, g
			case "r":
				return acl.Read, g
			case "u":
				return acl.Unlink, g
			}
			return false, 0
		}
	}
	return false, 0
}

// Val2Boolean converts a field value to boolean
func Val2Boolean(f interface{}) (bool, bool) {
	switch val := f.(type) {
	case bool:
		return bool(val), true
	default:
		return false, false
	}
}

// BooleanField gets boolean value from a record
func BooleanField(rec interface{}, name string) (bool, error) {
	m, ok := rec.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("Not a record")
	}

	iv, ok := m[name]
	if !ok {
		return false, fmt.Errorf("Field not found: %s", name)
	}

	val, ok := Val2Boolean(iv)
	if !ok {
		return false, fmt.Errorf("Error reading float from %s", name)
	}

	return val, nil

}
