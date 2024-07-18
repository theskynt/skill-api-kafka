package skill

import (
	"database/sql"
	"log"
	"reflect"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite", "file:TestCreateSkillHandlerIT?mode=memory&cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS skill (
	key TEXT PRIMARY KEY,
	name TEXT NOT NULL DEFAULT '',
	description TEXT NOT NULL DEFAULT '',
	logo TEXT NOT NULL DEFAULT '',
	tags TEXT [] NOT NULL DEFAULT '{}'
);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	return db
}

func TestSkillStorage(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	storage := NewStorage(db)

	testSkill := Skill{
		Key:         "test-skill",
		Name:        "Test Skill",
		Description: "This is a test skill",
		Logo:        "test-logo-url",
		Tags:        []string{"tag1", "tag2"},
	}

	t.Run("PostSkill", func(t *testing.T) {
		createdSkill, err := storage.PostSkill(testSkill)
		if err != nil {
			t.Fatalf("PostSkill error: %v", err)
		}
		if createdSkill.Key != testSkill.Key {
			t.Errorf("Expected key %s, got %s", testSkill.Key, createdSkill.Key)
		}
	})

	t.Run("FindAllSkill", func(t *testing.T) {
		skills, err := storage.FindAllSkill()
		if err != nil {
			t.Fatalf("FindAllSkill error: %v", err)
		}
		if len(skills) == 0 {
			t.Error("Expected at least one skill, got none")
		}
	})

	t.Run("FindSkillByKey", func(t *testing.T) {
		foundSkill, err := storage.FindSkillByKey(testSkill.Key)
		if err != nil {
			t.Fatalf("FindSkillByKey error: %v", err)
		}
		if foundSkill.Key != testSkill.Key {
			t.Errorf("Expected key %s, got %s", testSkill.Key, foundSkill.Key)
		}
	})

	t.Run("EditSkill", func(t *testing.T) {
		testEditSkill := Skill{
			Key:         "test-skill",
			Name:        "Edit Skill",
			Description: "This is an Edit skill",
			Logo:        "Edit-logo-url",
			Tags:        []string{"Edittag1", "Edittag2"},
		}
		updatedSkill, err := storage.EditSkill(testEditSkill)
		if err != nil {
			t.Fatalf("EditSkill error: %v", err)
		}
		if !reflect.DeepEqual(updatedSkill, testEditSkill) {
			t.Errorf("Expected skill %v, got %v", testEditSkill, updatedSkill)
		}
	})

	t.Run("EditSkillName", func(t *testing.T) {
		newName := "Updated Test Skill"
		updatedSkill, err := storage.EditSkillName(testSkill.Key, newName)
		if err != nil {
			t.Fatalf("EditSkill error: %v", err)
		}
		if updatedSkill.Name != newName {
			t.Errorf("Expected name %s, got %s", newName, updatedSkill.Name)
		}
	})

	t.Run("EditSkillDescription", func(t *testing.T) {
		newDescription := "Updated Description"
		updatedSkill, err := storage.EditSkillDescription(testSkill.Key, newDescription)
		if err != nil {
			t.Fatalf("EditSkill error: %v", err)
		}
		if updatedSkill.Description != newDescription {
			t.Errorf("Expected name %s, got %s", newDescription, updatedSkill.Description)
		}
	})

	t.Run("EditSkillLogo", func(t *testing.T) {
		newLogo := "Updated Logo"
		updatedSkill, err := storage.EditSkillLogo(testSkill.Key, newLogo)
		if err != nil {
			t.Fatalf("EditSkill error: %v", err)
		}
		if updatedSkill.Logo != newLogo {
			t.Errorf("Expected name %s, got %s", newLogo, updatedSkill.Logo)
		}
	})

	t.Run("EditSkillTags", func(t *testing.T) {
		newTags := []string{"Updatedtag1", "Updatedtag2"}
		updatedSkill, err := storage.EditSkillTags(testSkill.Key, newTags)
		if err != nil {
			t.Fatalf("EditSkill error: %v", err)
		}
		if !reflect.DeepEqual(updatedSkill.Tags, newTags) {
			t.Errorf("Expected tags %v, got %v", newTags, updatedSkill.Tags)
		}
	})
	t.Run("DeleteSkill", func(t *testing.T) {
		result := storage.DeleteSkill(testSkill.Key)
		if result != "success" {
			t.Errorf("DeleteSkill failed, expected 'success', got '%s'", result)
		}
	})
}
