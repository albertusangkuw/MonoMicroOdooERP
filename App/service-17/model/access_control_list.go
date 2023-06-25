package model

//Setara dengan ir_model_access
type AccessControlList struct {
	GroupID int  // PK
	IrModelAccessID int
	ModelID int
	Read bool
	Write bool
	Create bool
	Unlink bool
}
