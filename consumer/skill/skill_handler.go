package skill

import (
	"log"
)

type ActionHandler struct {
	storage storager
}

func NewActionHandler(storage storager) *ActionHandler {
	return &ActionHandler{storage: storage}
}

func (a *ActionHandler) HandleAction(message message) {
	switch message.Action {
	case "Insert":
		if _, err := a.storage.PostSkill(message.Data); err != nil {
			log.Printf("Failed to insert skill: %v", err)
		}
	case "Update":
		if _, err := a.storage.EditSkill(message.Data); err != nil {
			log.Printf("Failed to update skill: %v", err)
		}
	case "UpdateName":
		if _, err := a.storage.EditSkillName(message.Key, message.Data.Name); err != nil {
			log.Printf("Failed to update skill name: %v", err)
		}
	case "UpdateDescription":
		if _, err := a.storage.EditSkillDescription(message.Key, message.Data.Description); err != nil {
			log.Printf("Failed to update skill description: %v", err)
		}
	case "UpdateLogo":
		if _, err := a.storage.EditSkillLogo(message.Key, message.Data.Logo); err != nil {
			log.Printf("Failed to update skill logo: %v", err)
		}
	case "UpdateTags":
		if _, err := a.storage.EditSkillTags(message.Key, message.Data.Tags); err != nil {
			log.Printf("Failed to update skill tags: %v", err)
		}
	case "DeleteSkill":
		if res := a.storage.DeleteSkill(message.Key); res != "success" {
			log.Printf("Failed to delete skill: %v")
		}
	default:
		log.Printf("Unknown action: %s", message.Action)
	}
}
