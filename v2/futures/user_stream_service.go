package futures

import (
	"context"
	"net/http"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	m := map[string]interface{}{
		"url":    "/fapi/v3/listenKey",
		"method": http.MethodPost,
		"params": map[string]interface{}{},
	}
	data, err := s.c.call(m, true)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	m := map[string]interface{}{
		"url":    "/fapi/v3/listenKey",
		"method": http.MethodPost,
		"params": map[string]interface{}{
			"listenKey": s.listenKey,
		},
	}
	_, err = s.c.call(m, true)
	return err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	m := map[string]interface{}{
		"url":    "/fapi/v3/listenKey",
		"method": http.MethodDelete,
		"params": map[string]interface{}{
			"listenKey": s.listenKey,
		},
	}
	_, err = s.c.call(m, true)
	return err
}
