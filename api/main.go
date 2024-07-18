package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/narunart-atise/skill-api-kafka/api/database"
	"github.com/narunart-atise/skill-api-kafka/api/skill"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db, closeDB := database.NewPostgres()
	s := skill.NewStorage(db)
	h := skill.NewHandler(s)

	r := gin.Default()

	skillRoute := r.Group("/api/v1/skills")
	skillRoute.GET("", h.GetAllSkill)
	skillRoute.GET(":key", h.GetSkillByKey)
	skillRoute.POST("", h.CreateSkill)
	skillRoute.PUT(":key", h.UpdateSkill)
	skillRoute.PATCH(":key/actions/name", h.UpdateSkillName)
	skillRoute.PATCH(":key/actions/description", h.UpdateSkillDescription)
	skillRoute.PATCH(":key/actions/logo", h.UpdateSkillLogo)
	skillRoute.PATCH(":key/actions/tags", h.UpdateSkillTag)
	skillRoute.DELETE(":key", h.DeleteSkill)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("shuttign down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		defer closeDB()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err)
			}
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
