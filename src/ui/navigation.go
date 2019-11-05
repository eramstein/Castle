package ui

func toggleInfoDetails(ui *State, entityType, entity int) {
	if ui.EntityDetails.ID == entity && ui.EntityDetails.Type == entityType {
		ui.EntityDetails.ID = 0
		ui.EntityDetails.Type = 0
	} else {
		ui.EntityDetails.ID = entity
		ui.EntityDetails.Type = entityType
	}

}
