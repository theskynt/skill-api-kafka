package skill

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type storage struct {
	db *sql.DB
}

type storager interface {
	FindAllSkill() ([]Skill, error)
	FindSkillByKey(key string) (Skill, error)
	PostSkill(skill Skill) (Skill, error)
	EditSkill(skill Skill) (Skill, error)
	EditSkillName(key string, name string) (Skill, error)
	EditSkillDescription(key, description string) (Skill, error)
	EditSkillLogo(key, logo string) (Skill, error)
	EditSkillTags(key string, Tags []string) (Skill, error)
	DeleteSkill(rowKey string) string
}

func NewStorage(db *sql.DB) *storage {
	return &storage{db}
}

func (s storage) FindAllSkill() ([]Skill, error) {
	rows, err := s.db.Query("SELECT key, name, description,logo,tags FROM skill")
	if err != nil {
		return []Skill{}, nil
	}

	var Skills []Skill
	for rows.Next() {
		var skill Skill
		err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}

		Skills = append(Skills, Skill{
			Key:         skill.Key,
			Name:        skill.Name,
			Description: skill.Description,
			Logo:        skill.Logo,
			Tags:        skill.Tags,
		})
	}

	return Skills, nil
}

func (s storage) FindSkillByKey(key string) (Skill, error) {
	q := "SELECT key, name, description,logo,tags FROM skill WHERE key=$1"
	row := s.db.QueryRow(q, key)

	var skill Skill
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	if err != nil {
		return Skill{}, err
	}

	return Skill{
		Key:         skill.Key,
		Name:        skill.Name,
		Description: skill.Description,
		Logo:        skill.Logo,
		Tags:        skill.Tags,
	}, nil
}

func (s storage) PostSkill(skill Skill) (Skill, error) {
	q := "INSERT INTO skill (key,name, description,logo,tags) values ($1, $2,$3,$4,$5) RETURNING key"
	row := s.db.QueryRow(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))

	var keyid string
	err := row.Scan(&keyid)
	if err != nil {
		return Skill{}, err
	}
	return s.FindSkillByKey(keyid)
}

func (s storage) EditSkill(skill Skill) (Skill, error) {
	q := "UPDATE skill SET name=$2, description=$3, logo=$4, tags=$5 WHERE key=$1;"
	if _, err := s.db.Exec(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags)); err != nil {
		return Skill{}, err
	}

	return s.FindSkillByKey(skill.Key)
}

func (s storage) EditSkillName(key string, name string) (Skill, error) {
	q := "UPDATE skill SET name=$2 WHERE key=$1;"
	if _, err := s.db.Exec(q, key, name); err != nil {
		return Skill{}, err
	}
	return s.FindSkillByKey(key)
}

func (s storage) EditSkillDescription(key, description string) (Skill, error) {
	q := "UPDATE skill SET description=$2 WHERE key=$1;"
	if _, err := s.db.Exec(q, key, description); err != nil {
		return Skill{}, err
	}
	return s.FindSkillByKey(key)
}

func (s storage) EditSkillLogo(key, logo string) (Skill, error) {
	q := "UPDATE skill SET logo=$2 WHERE key=$1;"
	if _, err := s.db.Exec(q, key, logo); err != nil {
		return Skill{}, err
	}
	return s.FindSkillByKey(key)
}

func (s storage) EditSkillTags(key string, Tags []string) (Skill, error) {
	q := "UPDATE skill SET tags=$2 WHERE key=$1;"
	if _, err := s.db.Exec(q, key, pq.Array(Tags)); err != nil {
		return Skill{}, err
	}
	return s.FindSkillByKey(key)
}

func (s storage) DeleteSkill(rowKey string) string {
	q := "DELETE FROM skill WHERE key=$1;"
	if _, err := s.db.Exec(q, rowKey); err != nil {
		return "fail"
	}

	return "success"
}
