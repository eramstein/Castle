package ui

func toggleInfoDetails(ui *State, entityType, entity, data1, data2 int) {
	if ui.EntityDetails.ID == entity && ui.EntityDetails.Type == entityType {
		ui.EntityDetails.ID = 0
		ui.EntityDetails.Type = 0
		ui.EntityDetails.Data1 = 0
		ui.EntityDetails.Data2 = 0
	} else {
		setInfoDetails(ui, entityType, entity, data1, data2)
	}
}

func setInfoDetails(ui *State, entityType, entity, data1, data2 int) {
	ui.EntityDetails.ID = entity
	ui.EntityDetails.Type = entityType
	ui.EntityDetails.Data1 = data1
	ui.EntityDetails.Data2 = data2
}
