package main

import (
	"common"
	pb "common/api/user_service/proto"
	"gateway/gateway"
	"net/http"
	"service_one/types"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	gateway gateway.OMSGateway
}

func NewHandler(gateway gateway.OMSGateway) *handler {
	return &handler{gateway}
}

func (h *handler) registerRoutes(router *gin.Engine) {
	router.POST("/create_user", h.CreateUser)
}

func (h *handler) CreateUser(ctx *gin.Context){
	var req types.CreateUserRequest

	if err := common.ReadJSON(ctx, &req); err != nil {
		common.WriteError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tr := otel.Tracer("gin-server")
	_, span := tr.Start(ctx.Request.Context(), "CreateUser", trace.WithAttributes(attribute.String("user_name", req.UserName)))
	defer span.End()

	o, err := h.gateway.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		UserName: req.UserName,
		Password: req.Password,
	})
	rStatus := status.Convert(err)
	if rStatus != nil {
		span.SetStatus(otelCodes.Error, err.Error())

		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(ctx, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(ctx, http.StatusOK, o)

}

