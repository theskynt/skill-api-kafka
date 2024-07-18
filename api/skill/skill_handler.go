package skill

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	st       storager
	producer *Producer
}

func NewHandler(st storager, producer *Producer) *handler {
	return &handler{st: st, producer: producer}
}


type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func (h handler) GetAllSkill(c *gin.Context) {
	skills, err := h.st.FindAllSkill()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "skill not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skills,
	})
}

func (h handler) GetSkillByKey(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}

	getSkill, err := h.st.FindSkillByKey(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "skill not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   getSkill,
	})
}

func (h handler) CreateSkill(c *gin.Context) {
	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "Skill deatail is required",
		})
		return
	}

	// newSkill, err := h.st.PostSkill(skill)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, ResponseError{
	// 		Status:  "error",
	// 		Message: "Skill already exist",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"data":   newSkill,
	// })

	if err := h.producer.SendMessageWithAction("Insert", skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}
}

func (h handler) UpdateSkill(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}

	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "Request payload is invalid",
		})
		return
	}
	skill.Key = key
	newSkill, err := h.st.EditSkill(skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to update skill",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSkill,
	})
}

func (h handler) DeleteSkill(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return

	}

	res := h.st.DeleteSkill(key)
	if res != "success" {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to delete skill",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "Skill deleted",
	})
}

func (h handler) UpdateSkillName(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}

	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Request payload is invalid",
		})
		return
	}

	newSkill, err := h.st.EditSkillName(key, skill.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to update skill name",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSkill,
	})
}

func (h handler) UpdateSkillDescription(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}

	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "Request payload is invalid",
		})
		return
	}

	newSkill, err := h.st.EditSkillDescription(key, skill.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to update skill description",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSkill,
	})
}

func (h handler) UpdateSkillLogo(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}

	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "Request payload is invalid",
		})
		return
	}

	newSkill, err := h.st.EditSkillLogo(key, skill.Logo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to update skill logo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSkill,
	})
}

func (h handler) UpdateSkillTag(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "key is required",
		})
		return
	}
	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Status:  "error",
			Message: "Request payload is invalid",
		})
		return
	}
	newSkill, err := h.st.EditSkillTags(key, skill.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "not be able to update skill Tags",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSkill,
	})
}
