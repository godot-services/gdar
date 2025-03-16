package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetEditResponse struct {
	EditID           string                   `json:"edit_id"`
	AssetID          string                   `json:"asset_id"`
	UserID           string                   `json:"user_id"`
	Title            interface{}              `json:"title"`
	Description      interface{}              `json:"description"`
	CategoryID       interface{}              `json:"category_id"`
	GodotVersion     interface{}              `json:"godot_version"`
	VersionString    interface{}              `json:"version_string"`
	Cost             interface{}              `json:"cost"`
	DownloadProvider interface{}              `json:"download_provider"`
	DownloadCommit   interface{}              `json:"download_commit"`
	BrowseURL        string                   `json:"browse_url"`
	IssuesURL        string                   `json:"issues_url"`
	IconURL          interface{}              `json:"icon_url"`
	DownloadURL      string                   `json:"download_url"`
	Author           string                   `json:"author"`
	Previews         []GetEditResponsePreview `json:"previews"`
	Original         GetAssetResponse         `json:"original"`
	Status           string                   `json:"status"`
	Reason           string                   `json:"reason"`
	Warning          string                   `json:"warning"`
}

type GetEditResponsePreview struct {
	Operation     *string `json:"operation,omitempty"`
	EditPreviewID *string `json:"edit_preview_id,omitempty"`
	PreviewID     *string `json:"preview_id"`
	Type          string  `json:"type"`
	Link          string  `json:"link"`
	Thumbnail     string  `json:"thumbnail"`
}

// Returns a previously-submitted asset edit.
// All fields with null are unchanged, and will stay the same as in the original.
// The previews array is merged from the new and original previews.
func GetEditHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type UpdateEditRequest struct {
	Token            string                     `json:"token"`
	Title            string                     `json:"title"`
	Description      string                     `json:"description"`
	CategoryID       string                     `json:"category_id"`
	GodotVersion     string                     `json:"godot_version"`
	VersionString    string                     `json:"version_string"`
	Cost             string                     `json:"cost"`
	DownloadProvider string                     `json:"download_provider"`
	DownloadCommit   string                     `json:"download_commit"`
	BrowseURL        string                     `json:"browse_url"`
	IssuesURL        string                     `json:"issues_url"`
	IconURL          string                     `json:"icon_url"`
	DownloadURL      string                     `json:"download_url"`
	Previews         []UpdateEditRequestPreview `json:"previews"`
}

type UpdateEditRequestPreview struct {
	Enabled       bool    `json:"enabled"`
	Operation     string  `json:"operation"`
	Type          *string `json:"type,omitempty"`
	Link          *string `json:"link,omitempty"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
	EditPreviewID *string `json:"edit_preview_id,omitempty"`
}

type UpdateEditResponse struct {
	ID string `json:"id"`
}

// Update an existing one. Fields are required when creating a new asset, and are optional otherwise. Same for previews -- required when creating a new preview, may be missing if updating one.
//
// Notes:
//   - Not passing "enabled": true for previews will result in them not being included in the edit. This may be fixed in the future.
//   - version_string is free-form text, but major.minor or major.minor.patch format is best.
//   - Available download providers can be seen on the asset library fronted.
func UpdateEditHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type ReviewEditRequest struct {
	Token string `json:"token"`
}

// Successful result: the asset edit, without the original asset.
//
// Moderator-only. Put an edit in review. It is impossible to change it after this.
func ReviewEditHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type AcceptEditRequest struct {
	Token string `json:"token"`
}

// Successful result: the asset edit, without the original asset.
//
// Moderator-only. Apply an edit previously put in review.
func AcceptEditHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type RejectEditRequest struct {
	Token  string `json:"token"`
	Reason string `json:"reason"`
}

// Successful result: the asset edit, without the original asset.
//
// Moderator-only. Reject an edit previously put in review.
func RejectEditHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}
