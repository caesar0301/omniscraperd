package page

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// FrameResource information about the Resource on the page.
type FrameResource struct {
	URL          string               `json:"url"`                    // Resource URL.
	Type         network.ResourceType `json:"type"`                   // Type of this resource.
	MimeType     string               `json:"mimeType"`               // Resource mimeType as determined by the browser.
	LastModified *cdp.TimeSinceEpoch  `json:"lastModified,omitempty"` // last-modified timestamp as reported by server.
	ContentSize  float64              `json:"contentSize,omitempty"`  // Resource content size.
	Failed       bool                 `json:"failed,omitempty"`       // True if the resource failed to load.
	Canceled     bool                 `json:"canceled,omitempty"`     // True if the resource was canceled during loading.
}

// FrameResourceTree information about the Frame hierarchy along with their
// cached resources.
type FrameResourceTree struct {
	Frame       *cdp.Frame           `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []*FrameResource     `json:"resources"`             // Information about frame resources.
}

// FrameTree information about the Frame hierarchy.
type FrameTree struct {
	Frame       *cdp.Frame   `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*FrameTree `json:"childFrames,omitempty"` // Child frames.
}

// ScriptIdentifier unique script identifier.
type ScriptIdentifier string

// String returns the ScriptIdentifier as string value.
func (t ScriptIdentifier) String() string {
	return string(t)
}

// TransitionType transition type.
type TransitionType string

// String returns the TransitionType as string value.
func (t TransitionType) String() string {
	return string(t)
}

// TransitionType values.
const (
	TransitionTypeLink             TransitionType = "link"
	TransitionTypeTyped            TransitionType = "typed"
	TransitionTypeAddressBar       TransitionType = "address_bar"
	TransitionTypeAutoBookmark     TransitionType = "auto_bookmark"
	TransitionTypeAutoSubframe     TransitionType = "auto_subframe"
	TransitionTypeManualSubframe   TransitionType = "manual_subframe"
	TransitionTypeGenerated        TransitionType = "generated"
	TransitionTypeAutoToplevel     TransitionType = "auto_toplevel"
	TransitionTypeFormSubmit       TransitionType = "form_submit"
	TransitionTypeReload           TransitionType = "reload"
	TransitionTypeKeyword          TransitionType = "keyword"
	TransitionTypeKeywordGenerated TransitionType = "keyword_generated"
	TransitionTypeOther            TransitionType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TransitionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t TransitionType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TransitionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch TransitionType(in.String()) {
	case TransitionTypeLink:
		*t = TransitionTypeLink
	case TransitionTypeTyped:
		*t = TransitionTypeTyped
	case TransitionTypeAddressBar:
		*t = TransitionTypeAddressBar
	case TransitionTypeAutoBookmark:
		*t = TransitionTypeAutoBookmark
	case TransitionTypeAutoSubframe:
		*t = TransitionTypeAutoSubframe
	case TransitionTypeManualSubframe:
		*t = TransitionTypeManualSubframe
	case TransitionTypeGenerated:
		*t = TransitionTypeGenerated
	case TransitionTypeAutoToplevel:
		*t = TransitionTypeAutoToplevel
	case TransitionTypeFormSubmit:
		*t = TransitionTypeFormSubmit
	case TransitionTypeReload:
		*t = TransitionTypeReload
	case TransitionTypeKeyword:
		*t = TransitionTypeKeyword
	case TransitionTypeKeywordGenerated:
		*t = TransitionTypeKeywordGenerated
	case TransitionTypeOther:
		*t = TransitionTypeOther

	default:
		in.AddError(errors.New("unknown TransitionType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TransitionType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// NavigationEntry navigation history entry.
type NavigationEntry struct {
	ID             int64          `json:"id"`             // Unique id of the navigation history entry.
	URL            string         `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string         `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string         `json:"title"`          // Title of the navigation history entry.
	TransitionType TransitionType `json:"transitionType"` // Transition type.
}

// ScreencastFrameMetadata screencast frame metadata.
type ScreencastFrameMetadata struct {
	OffsetTop       float64             `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64             `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64             `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64             `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64             `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64             `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       *cdp.TimeSinceEpoch `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// DialogType javascript dialog type.
type DialogType string

// String returns the DialogType as string value.
func (t DialogType) String() string {
	return string(t)
}

// DialogType values.
const (
	DialogTypeAlert        DialogType = "alert"
	DialogTypeConfirm      DialogType = "confirm"
	DialogTypePrompt       DialogType = "prompt"
	DialogTypeBeforeunload DialogType = "beforeunload"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DialogType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DialogType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DialogType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DialogType(in.String()) {
	case DialogTypeAlert:
		*t = DialogTypeAlert
	case DialogTypeConfirm:
		*t = DialogTypeConfirm
	case DialogTypePrompt:
		*t = DialogTypePrompt
	case DialogTypeBeforeunload:
		*t = DialogTypeBeforeunload

	default:
		in.AddError(errors.New("unknown DialogType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DialogType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AppManifestError error while paring app manifest.
type AppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int64  `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int64  `json:"line"`     // Error line.
	Column   int64  `json:"column"`   // Error column.
}

// LayoutViewport layout viewport position and dimensions.
type LayoutViewport struct {
	PageX        int64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// VisualViewport visual viewport position, dimensions, and scale.
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

// Viewport viewport for capturing screenshot.
type Viewport struct {
	X      float64 `json:"x"`      // X offset in CSS pixels.
	Y      float64 `json:"y"`      // Y offset in CSS pixels
	Width  float64 `json:"width"`  // Rectangle width in CSS pixels
	Height float64 `json:"height"` // Rectangle height in CSS pixels
	Scale  float64 `json:"scale"`  // Page scale factor.
}

// FontFamilies generic font families collection.
type FontFamilies struct {
	Standard   string `json:"standard,omitempty"`   // The standard font-family.
	Fixed      string `json:"fixed,omitempty"`      // The fixed font-family.
	Serif      string `json:"serif,omitempty"`      // The serif font-family.
	SansSerif  string `json:"sansSerif,omitempty"`  // The sansSerif font-family.
	Cursive    string `json:"cursive,omitempty"`    // The cursive font-family.
	Fantasy    string `json:"fantasy,omitempty"`    // The fantasy font-family.
	Pictograph string `json:"pictograph,omitempty"` // The pictograph font-family.
}

// FontSizes default font sizes.
type FontSizes struct {
	Standard int64 `json:"standard,omitempty"` // Default standard font size.
	Fixed    int64 `json:"fixed,omitempty"`    // Default fixed font size.
}

// FrameScheduledNavigationReason the reason for the navigation.
type FrameScheduledNavigationReason string

// String returns the FrameScheduledNavigationReason as string value.
func (t FrameScheduledNavigationReason) String() string {
	return string(t)
}

// FrameScheduledNavigationReason values.
const (
	FrameScheduledNavigationReasonFormSubmissionGet     FrameScheduledNavigationReason = "formSubmissionGet"
	FrameScheduledNavigationReasonFormSubmissionPost    FrameScheduledNavigationReason = "formSubmissionPost"
	FrameScheduledNavigationReasonHTTPHeaderRefresh     FrameScheduledNavigationReason = "httpHeaderRefresh"
	FrameScheduledNavigationReasonScriptInitiated       FrameScheduledNavigationReason = "scriptInitiated"
	FrameScheduledNavigationReasonMetaTagRefresh        FrameScheduledNavigationReason = "metaTagRefresh"
	FrameScheduledNavigationReasonPageBlockInterstitial FrameScheduledNavigationReason = "pageBlockInterstitial"
	FrameScheduledNavigationReasonReload                FrameScheduledNavigationReason = "reload"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t FrameScheduledNavigationReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t FrameScheduledNavigationReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *FrameScheduledNavigationReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch FrameScheduledNavigationReason(in.String()) {
	case FrameScheduledNavigationReasonFormSubmissionGet:
		*t = FrameScheduledNavigationReasonFormSubmissionGet
	case FrameScheduledNavigationReasonFormSubmissionPost:
		*t = FrameScheduledNavigationReasonFormSubmissionPost
	case FrameScheduledNavigationReasonHTTPHeaderRefresh:
		*t = FrameScheduledNavigationReasonHTTPHeaderRefresh
	case FrameScheduledNavigationReasonScriptInitiated:
		*t = FrameScheduledNavigationReasonScriptInitiated
	case FrameScheduledNavigationReasonMetaTagRefresh:
		*t = FrameScheduledNavigationReasonMetaTagRefresh
	case FrameScheduledNavigationReasonPageBlockInterstitial:
		*t = FrameScheduledNavigationReasonPageBlockInterstitial
	case FrameScheduledNavigationReasonReload:
		*t = FrameScheduledNavigationReasonReload

	default:
		in.AddError(errors.New("unknown FrameScheduledNavigationReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *FrameScheduledNavigationReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// CaptureScreenshotFormat image compression format (defaults to png).
type CaptureScreenshotFormat string

// String returns the CaptureScreenshotFormat as string value.
func (t CaptureScreenshotFormat) String() string {
	return string(t)
}

// CaptureScreenshotFormat values.
const (
	CaptureScreenshotFormatJpeg CaptureScreenshotFormat = "jpeg"
	CaptureScreenshotFormatPng  CaptureScreenshotFormat = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CaptureScreenshotFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CaptureScreenshotFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CaptureScreenshotFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CaptureScreenshotFormat(in.String()) {
	case CaptureScreenshotFormatJpeg:
		*t = CaptureScreenshotFormatJpeg
	case CaptureScreenshotFormatPng:
		*t = CaptureScreenshotFormatPng

	default:
		in.AddError(errors.New("unknown CaptureScreenshotFormat value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CaptureScreenshotFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// CaptureSnapshotFormat format (defaults to mhtml).
type CaptureSnapshotFormat string

// String returns the CaptureSnapshotFormat as string value.
func (t CaptureSnapshotFormat) String() string {
	return string(t)
}

// CaptureSnapshotFormat values.
const (
	CaptureSnapshotFormatMhtml CaptureSnapshotFormat = "mhtml"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CaptureSnapshotFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CaptureSnapshotFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CaptureSnapshotFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CaptureSnapshotFormat(in.String()) {
	case CaptureSnapshotFormatMhtml:
		*t = CaptureSnapshotFormatMhtml

	default:
		in.AddError(errors.New("unknown CaptureSnapshotFormat value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CaptureSnapshotFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetDownloadBehaviorBehavior whether to allow all or deny all download
// requests, or use default Chrome behavior if available (otherwise deny).
type SetDownloadBehaviorBehavior string

// String returns the SetDownloadBehaviorBehavior as string value.
func (t SetDownloadBehaviorBehavior) String() string {
	return string(t)
}

// SetDownloadBehaviorBehavior values.
const (
	SetDownloadBehaviorBehaviorDeny    SetDownloadBehaviorBehavior = "deny"
	SetDownloadBehaviorBehaviorAllow   SetDownloadBehaviorBehavior = "allow"
	SetDownloadBehaviorBehaviorDefault SetDownloadBehaviorBehavior = "default"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetDownloadBehaviorBehavior) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetDownloadBehaviorBehavior) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetDownloadBehaviorBehavior) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetDownloadBehaviorBehavior(in.String()) {
	case SetDownloadBehaviorBehaviorDeny:
		*t = SetDownloadBehaviorBehaviorDeny
	case SetDownloadBehaviorBehaviorAllow:
		*t = SetDownloadBehaviorBehaviorAllow
	case SetDownloadBehaviorBehaviorDefault:
		*t = SetDownloadBehaviorBehaviorDefault

	default:
		in.AddError(errors.New("unknown SetDownloadBehaviorBehavior value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetDownloadBehaviorBehavior) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ScreencastFormat image compression format.
type ScreencastFormat string

// String returns the ScreencastFormat as string value.
func (t ScreencastFormat) String() string {
	return string(t)
}

// ScreencastFormat values.
const (
	ScreencastFormatJpeg ScreencastFormat = "jpeg"
	ScreencastFormatPng  ScreencastFormat = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ScreencastFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ScreencastFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ScreencastFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ScreencastFormat(in.String()) {
	case ScreencastFormatJpeg:
		*t = ScreencastFormatJpeg
	case ScreencastFormatPng:
		*t = ScreencastFormatPng

	default:
		in.AddError(errors.New("unknown ScreencastFormat value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ScreencastFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetWebLifecycleStateState target lifecycle state.
type SetWebLifecycleStateState string

// String returns the SetWebLifecycleStateState as string value.
func (t SetWebLifecycleStateState) String() string {
	return string(t)
}

// SetWebLifecycleStateState values.
const (
	SetWebLifecycleStateStateFrozen SetWebLifecycleStateState = "frozen"
	SetWebLifecycleStateStateActive SetWebLifecycleStateState = "active"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetWebLifecycleStateState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetWebLifecycleStateState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetWebLifecycleStateState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SetWebLifecycleStateState(in.String()) {
	case SetWebLifecycleStateStateFrozen:
		*t = SetWebLifecycleStateStateFrozen
	case SetWebLifecycleStateStateActive:
		*t = SetWebLifecycleStateStateActive

	default:
		in.AddError(errors.New("unknown SetWebLifecycleStateState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetWebLifecycleStateState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
