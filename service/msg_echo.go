package service

import (
	"io"
	"net/http"

	"github.com/233liyu/wechat-app/logger"
)

type WeChatHandler struct{}

func (*WeChatHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty request body"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read request body"))
		return
	}
	log := logger.GetLogger()
	log.Info("request body: " + string(requestBody))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

}
