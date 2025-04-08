package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/glamostoffer/arete/auth/internal/domain"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	cl *redis.Client
}

func New(cl *redis.Client) *cache {
	return &cache{
		cl: cl,
	}
}

func (c *cache) getSignUpRequestKey(email string) string {
	return fmt.Sprintf("sign_up_req:email:%s", email)
}

func (c *cache) SetSignUpRequest(
	ctx context.Context,
	req domain.SignUpRequest,
	ttl time.Duration,
) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	set, err := c.cl.SetNX(ctx, c.getSignUpRequestKey(req.Email), data, ttl).Result()
	if err != nil {
		return err
	}
	if !set {
		return errlist.ErrResendCooldown
	}

	return nil
}

func (c *cache) GetSignUpRequest(ctx context.Context, email string) (res domain.SignUpRequest, err error) {
	data, err := c.cl.GetDel(ctx, c.getSignUpRequestKey(email)).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *cache) getUserSessionKey(refToken string) string {
	return fmt.Sprintf("user_session:%s", refToken)
}

func (c *cache) SetUserSession(ctx context.Context, refToken string, userSession domain.UserSession) error {
	sessionData, err := json.Marshal(userSession)
	if err != nil {
		return err
	}

	err = c.cl.Set(ctx, c.getUserSessionKey(refToken), sessionData, userSession.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) GetUserSession(ctx context.Context, refToken string) (res domain.UserSession, err error) {
	data, err := c.cl.Get(ctx, c.getUserSessionKey(refToken)).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
