package types

type Customer struct {
	ID   uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// >>> AUTO_GEN_NEW_TYPE <<< //
