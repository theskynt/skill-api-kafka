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

	if err := h.producer.SendMessageWithAction("Insert", skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"data":   newSkill,
	// })

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
	if err := h.producer.SendMessageWithAction("Update", skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
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

	if err := h.producer.SendMessageWithKey("UpdateName", key, skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
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

	if err := h.producer.SendMessageWithKey("UpdateDescription", key, skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
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

	if err := h.producer.SendMessageWithKey("UpdateLogo", key, skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
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

	if err := h.producer.SendMessageWithKey("UpdateTags", key, skill); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
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

	if err := h.producer.SendMessageKey("DeleteSkill", key); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Status:  "error",
			Message: "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "Skill deleted",
	})
}
