package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/handlers/transactions"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, registerUC *user.RegisterUseCase) {
	registerHandler := transactions.NewRegisterHandler(registerUC)
	r.POST("/ownerschool/register", registerHandler.Register)
}
