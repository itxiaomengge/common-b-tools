package result

type Res struct {
	Code    int         `json:"code" desc:"返回码"`
	Data    interface{} `json:"data" desc:"返回数据"`
	Message string      `json:"message" desc:"返回消息（比如Error信息）"`
}

func (r *Res) Ok(data interface{}) *Res {
	r.Code = CODE_SUCCESS
	r.Data = data
	r.Message = MESSAGE_SUCCESS
	return r
}

func (r *Res) Error(err error) *Res {
	r.Code = CODE_FAIL
	r.Data = nil
	r.Message = MESSAGE_FAIL + ": " + err.Error()
	return r
}
