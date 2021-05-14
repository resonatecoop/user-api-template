package iamserver

import (
	"context"
	"time"

	"github.com/resonatecoop/user-api/internal/model"

	iam "github.com/resonatecoop/user-api/proto/iam"
)

var svcName = "iam"

// NewLoggingService creates new Template logging service
func NewLoggingService(svc iam.ResonateIAMServer, logger model.Logger) *LoggingService {
	return &LoggingService{
		IAM:    svc,
		logger: logger,
	}
}

// LoggingService represents iam logging service
type LoggingService struct {
	IAM    iam.ResonateIAMServer
	logger model.Logger
}

// Auth logging
func (ls *LoggingService) Auth(ctx context.Context, req *iam.AuthReq) (resp *iam.AuthResp, err error) {
	defer func(begin time.Time) {
		req.Password = "xxx-redacted-xxx"
		ls.logger.Log(
			ctx,
			svcName, "Auth request", err,
			map[string]interface{}{
				"took": time.Since(begin),
				"req":  req,
				"resp": resp,
			},
		)
	}(time.Now())
	return ls.IAM.Auth(ctx, req)
}

// Refresh token logging
func (ls *LoggingService) Refresh(ctx context.Context, req *iam.RefreshReq) (resp *iam.RefreshResp, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			ctx,
			svcName, "Refresh request", err,
			map[string]interface{}{
				"took": time.Since(begin),
				"req":  req,
				"resp": resp,
			},
		)
	}(time.Now())
	return ls.IAM.Refresh(ctx, req)
}
