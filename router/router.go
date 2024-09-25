package router

import (
	"database/sql"
	"online-pathsaala/controller/user"
	"online-pathsaala/pkg/db"

	"github.com/gin-gonic/gin"
)

func AddRoutes(dbCon *sql.DB, e *gin.Engine) {
	dbMan := db.Database{
		Db: dbCon,
	}

	au := user.UserAcc{
		DdManager: &dbMan,
	}

	public := e.Group("/api")
	{
		public.POST("/register", au.Register)
		public.POST("/login", au.Login)
	}
}
