package chromedp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/andreypanin/chromedp/domains/target"
)

type messageClientToServer struct {
	ID        int               `json:"id"`
	Method    string            `json:"method"`
	Params    any               `json:"params"`
	SessionID *target.SessionID `json:"sessionId,omitempty"`
}

type messageServerToClient struct {
	Params any       `json:"params"` // nil for results of method calls
	Result any       `json:"result"` // nil for events
	Error  *struct { // if not nil, params and result are nil
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

var (
	ErrBadWrite    = errors.New("bad write")
	ErrBadRead     = errors.New("bad read")
	ErrApplication = errors.New("application error")
)

type hasMethod interface {
	Method() string
}

func call[P hasMethod, R any](
	w io.WriteCloser,
	r io.Reader,
	params P,
	sessionID target.SessionID,
) (
	result R,
	err error,
) {

	written := messageClientToServer{
		ID:     1,
		Method: params.Method(),
		Params: params,
	}
	if sessionID != "" {
		written.SessionID = &sessionID
	}

	p, err := json.Marshal(written)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(p)
	if err != nil {
		return result, errors.Join(ErrBadWrite, err)
	}

	p = make([]byte, 0, 1024)

	p, err = io.ReadAll(r)
	if err != nil {
		return result, errors.Join(ErrBadRead, err)
	}

	read := messageServerToClient{}

	err = json.Unmarshal(p, &read)
	if err != nil {
		return result, errors.Join(ErrBadRead, err)
	}

	if read.Error != nil {
		err := fmt.Errorf("(%d) %s", read.Error.Code, read.Error.Message)
		return result, errors.Join(ErrApplication, err)
	}

	if read.Params != nil {
		// The read message is a ChromeDP event. The r's implementation must
		// skip such messages, because they are processed elsewhere.
		panic("unexpected event received")
	}

	return read.Result.(R), nil
}

var (
	TargetCreateTarget       = call[target.CreateTargetRequest, target.CreateTargetResponse]
	TargetGetBrowserContexts = call[target.GetBrowserContextsRequest, target.GetBrowserContextsResponse]
	TargetAttachToTarget     = call[target.AttachToTargetRequest, target.AttachToTargetResponse]
)
