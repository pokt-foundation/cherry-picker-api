package cpicker

import "github.com/pokt-foundation/cherry-picker-api/database"

type CherryPicker struct {
	Driver database.CherryPickerStorage
}
