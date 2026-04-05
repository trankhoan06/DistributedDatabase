package main

import (
	"esim/component/middleware"
	"esim/component/tokenprovider/jwt"
	"esim/config"
	emailSend "esim/email"
	"esim/modules/email/biz"
	storageEmail "esim/modules/email/storage"
	ginMail "esim/modules/email/transport/gin"
	storageUser "esim/modules/user/storage"
	ginUser "esim/modules/user/transport/gin"
	ProviderMysql "esim/provider/mysql"
	"esim/worker"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"log"
	"sync"
)

func main() {
	cfg := config.GetConfig()
	db, err := ProviderMysql.NewMysql(cfg)
	if err != nil {
		log.Fatal(err)
	}
	prefix := jwt.NewJwtProvider(cfg.SecretApp, cfg.PrefixApp)
	auth := storageUser.NewSql(db)
	middle := middleware.NewModelMiddleware(auth, prefix)
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.HostRedis + ":" + cfg.PortRedis,
		Password: cfg.PasswordRedis,
	}
	taskDistributor := worker.NewRedisTaskDistributor(&redisOpt)
	r := gin.Default()
	configCORS := setupCors()
	r.Use(cors.New(configCORS))
	r.Use(middle.Recover())
	user := r.Group("/user")
	{
		user.POST("/create_user", middle.RequestAuthorize(), ginUser.Register(db))
		user.POST("/forgot", ginUser.ForgotPassword(db, taskDistributor))
		user.POST("/login", ginUser.Login(db, cfg, prefix, taskDistributor))
		user.PATCH("/change_pass", middle.RequestAuthorize(), ginUser.ChangPassword(db))
		user.POST("/forgot_password", middle.RequestAuthorize(), ginUser.ForgotPassword(db, taskDistributor))
		user.PATCH("/change_password_forgot", middle.RequestAuthorize(), ginUser.ChangeForgotPassword(db))
		user.POST("/get_access_token", ginUser.GetAccessToken(db, prefix))
		user.PATCH("/update", middle.RequestAuthorize(), ginUser.UpdateUser(db))
		user.GET("/profile", middle.RequestAuthorize(), ginUser.GetProfile(db))
	}
	mail := r.Group("/mail")
	{
		mail.PATCH("/verify_code_email", ginMail.VerifyCodeEmail(db, prefix))
		mail.PATCH("/verify_forgot_password", ginMail.VerifyForgotPassword(db))
		mail.POST("/create_verify_code_email", ginMail.CreateVerifyCodeEmail(db))
	}
	accountSto := storageUser.NewSql(db)
	emailSto := storageEmail.NewSqlModel(db)
	emailCase := biz.NewSendEmailBiz(emailSto, accountSto)
	NewEmail := emailSend.NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		processor := worker.NewRedisTaskProcessor(&redisOpt,
			accountSto,
			NewEmail,
			emailCase,
		)
		err1 := processor.Start()
		if err1 != nil {
			log.Fatal(err1)

		}
	}()
	r.Run(":5000")

}
func setupCors() cors.Config {
	configCORS := cors.DefaultConfig()
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"}
	configCORS.AllowCredentials = true
	//configCORS.AllowOrigins = []string{"http://localhost:3000"}
	configCORS.AllowAllOrigins = true

	return configCORS
}
