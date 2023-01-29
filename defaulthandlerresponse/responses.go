package response

import (
	"context"
	"sync"
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	errorHandler    func(error) (int, interface{})
	errorHandlerCtx func(context.Context, error) (int, interface{})
	lock            sync.RWMutex
)

// @Title HandlerResponseResult
// @Description writes v as json string into w with code.
// @Author Xiaomeng.Ge
// @Date 2023-01-29 10:16:11
//
// @Param r *http.Request
// @Param w http.ResponseWriter
// @Param resp interface{}
// @Param err error
func HandlerResponseResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	// 得到统一返回结果
	response := NewDefaultHandlerResponse(resp, err)

	if err != nil {
		// 错误返回
		code := GetErrorCtxCode(r.Context(), w, err)
		httpx.WriteJsonCtx(r.Context(), w, code, response)
	} else {
		// 成功返回
		httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, response)
	}
}

// Error 得到错误状态码
func GetErrorCode(w http.ResponseWriter, err error, fns ...func(w http.ResponseWriter, err error)) int {
	lock.RLock()
	handler := errorHandler
	lock.RUnlock()

	return doHandleError(w, err, handler)
}

// ErrorCtx 得到错误状态码
func GetErrorCtxCode(ctx context.Context, w http.ResponseWriter, err error,
	fns ...func(w http.ResponseWriter, err error)) int {
	lock.RLock()
	handlerCtx := errorHandlerCtx
	lock.RUnlock()

	var handler func(error) (int, interface{})
	if handlerCtx != nil {
		handler = func(err error) (int, interface{}) {
			return handlerCtx(ctx, err)
		}
	}

	return doHandleError(w, err, handler)
}

// 得到错误状态码 code
func doHandleError(w http.ResponseWriter, err error, handler func(error) (int, interface{})) int {
	if handler == nil {
		if IsGrpcError(err) {
			// don't unwrap error and get status.Message(),
			// it hides the rpc error headers.
			http.Error(w, err.Error(), CodeFromGrpcError(err))
			return CodeFromGrpcError(err)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return http.StatusBadRequest
		}
	}

	code, body := handler(err)
	if body == nil {
		w.WriteHeader(code)
		return code
	}

	e, ok := body.(error)
	if ok {
		http.Error(w, e.Error(), code)
	}

	return code
}
