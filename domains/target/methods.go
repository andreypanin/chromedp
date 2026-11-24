package target

import "github.com/andreypanin/chromedp/domains/browser"

type GetBrowserContextsRequest struct{}

func (GetBrowserContextsRequest) Method() string {
	return "Target.getBrowserContexts"
}

type GetBrowserContextsResponse struct {
	BrowserContextIDs       []browser.BrowserContextID `json:"browserContextIds"`
	DefaultBrowserContextID *browser.BrowserContextID  `json:"defaultBrowserContextId"`
}

type CreateTargetRequest struct {
	URL                     string                    `json:"url"`
	Left                    *int                      `json:"left,omitempty"`
	Top                     *int                      `json:"top,omitempty"`
	Width                   *int                      `json:"width,omitempty"`
	Height                  *int                      `json:"height,omitempty"`
	WindowState             *WindowState              `json:"windowState,omitempty"`
	BrowserContextID        *browser.BrowserContextID `json:"browserContextId,omitempty"`
	EnableBeginFrameControl *bool                     `json:"enableBeginFrameControl,omitempty"`
	NewWindow               *bool                     `json:"newWindow,omitempty"`
	Background              *bool                     `json:"background,omitempty"`
	Focus                   *bool                     `json:"focus,omitempty"`
	ForTab                  *bool                     `json:"forTab,omitempty"`
	Hidden                  *bool                     `json:"hidden,omitempty"`
}

func (CreateTargetRequest) Method() string {
	return "Target.createTarget"
}

type CreateTargetResponse struct {
	TargetID TargetID `json:"targetId"`
}

type AttachToTargetRequest struct {
	TargetID TargetID `json:"targetId"`
	Flatten  *bool    `json:"flatten"`
}

func (AttachToTargetRequest) Method() string {
	return "Target.attachToTarget"
}

type AttachToTargetResponse struct {
	SessionID SessionID `json:"sessionId"`
}
