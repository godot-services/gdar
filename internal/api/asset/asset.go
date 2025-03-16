package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// QueryParameter for list of assets
// ?type=(any|addon|project)
// &category=(category id)
// &support=(official|featured|community|testing)
// &filter=(search text)
// &user=(submitter username)
// &cost=(license)
// &godot_version=(major).(minor).(patch)
// &max_results=(number 1…500)
// &page=(number, pages to skip) OR &offset=(number, rows to skip)
// &sort=(rating|cost|name|updated)
// &reverse
type ListAssetsRequest struct {
	Type         string `query:"type"`
	Category     string `query:"category"`
	Support      string `query:"support"`
	Filter       string `query:"filter"`
	User         string `query:"user"`
	Cost         string `query:"cost"`
	GodotVersion string `query:"godot_version"`
	MaxResults   int64  `query:"max_results"`
	Page         int64  `query:"page"`
	Offset       int64  `query:"offset"`
	Sort         string `query:"sort"`
	Reverse      string `query:"reverse"`
}

type ListAssetsResponse struct {
	Result     []ListAssetsResponseResult `json:"result"`
	Page       int64                      `json:"page"`
	Pages      int64                      `json:"pages"`
	PageLength int64                      `json:"page_length"`
	TotalItems int64                      `json:"total_items"`
}

type ListAssetsResponseResult struct {
	AssetID       string `json:"asset_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	AuthorID      string `json:"author_id"`
	Category      string `json:"category"`
	CategoryID    string `json:"category_id"`
	GodotVersion  string `json:"godot_version"`
	Rating        string `json:"rating"`
	Cost          string `json:"cost"`
	SupportLevel  string `json:"support_level"`
	IconURL       string `json:"icon_url"`
	Version       string `json:"version"`
	VersionString string `json:"version_string"`
	ModifyDate    string `json:"modify_date"`
}

// Get a list of assets.
//
// Some notes:
//   - Leading and trailing whitespace in filter is trimmed on the server side.
//   - For legacy purposes, not supplying godot version would list only 2.1 assets, while not supplying type would list only addons.
//   - To specify multiple support levels, join them with +, e.g. support=featured+community.
//   - Godot version can be specified as you see fit, for example, godot_version=3.1 or godot_version=3.1.5. Currently, the patch version is disregarded, but may be honored in the future.
func ListAssetsHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type GetAssetResponse struct {
	AssetID          string                    `json:"asset_id"`
	Type             string                    `json:"type"`
	Title            string                    `json:"title"`
	Author           string                    `json:"author"`
	AuthorID         string                    `json:"author_id"`
	Version          string                    `json:"version"`
	VersionString    string                    `json:"version_string"`
	Category         string                    `json:"category"`
	CategoryID       string                    `json:"category_id"`
	GodotVersion     string                    `json:"godot_version"`
	Rating           string                    `json:"rating"`
	Cost             string                    `json:"cost"`
	Description      string                    `json:"description"`
	SupportLevel     string                    `json:"support_level"`
	DownloadProvider string                    `json:"download_provider"`
	DownloadCommit   string                    `json:"download_commit"`
	DownloadHash     string                    `json:"download_hash"`
	BrowseURL        string                    `json:"browse_url"`
	IssuesURL        string                    `json:"issues_url"`
	IconURL          string                    `json:"icon_url"`
	Searchable       string                    `json:"searchable"`
	ModifyDate       string                    `json:"modify_date"`
	DownloadURL      string                    `json:"download_url"`
	Previews         []GetAssetResponsePreview `json:"previews"`
}

type GetAssetResponsePreview struct {
	PreviewID string `json:"preview_id"`
	Type      string `json:"type"`
	Link      string `json:"link"`
	Thumbnail string `json:"thumbnail"`
}

// Notes:
//   - The cost field is the license. Other asset libraries may put the price there and supply a download URL which requires authentication.
//   - In the official asset library, the download_hash field is always empty and is kept for compatibility only. The editor will skip hash checks if download_hash is an empty string. Third-party asset libraries may specify a SHA-256 hash to be used by the editor to verify the download integrity.
//   - The download URL is generated based on the download commit and the browse URL.
func GetAssetHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type CreateAssetRequest struct {
	Token            string                      `json:"token"`
	Title            string                      `json:"title"`
	Description      string                      `json:"description"`
	CategoryID       string                      `json:"category_id"`
	GodotVersion     string                      `json:"godot_version"`
	VersionString    string                      `json:"version_string"`
	Cost             string                      `json:"cost"`
	DownloadProvider string                      `json:"download_provider"`
	DownloadCommit   string                      `json:"download_commit"`
	BrowseURL        string                      `json:"browse_url"`
	IssuesURL        string                      `json:"issues_url"`
	IconURL          string                      `json:"icon_url"`
	DownloadURL      string                      `json:"download_url"`
	Previews         []CreateAssetRequestPreview `json:"previews"`
}

type CreateAssetRequestPreview struct {
	Enabled       bool    `json:"enabled"`
	Operation     string  `json:"operation"`
	Type          *string `json:"type,omitempty"`
	Link          *string `json:"link,omitempty"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
	EditPreviewID *string `json:"edit_preview_id,omitempty"`
}

type CreateAssetResponse struct {
	ID string `json:"id"`
}

// Create a new asset. Fields are required when creating a new asset, and are optional otherwise. Same for previews -- required when creating a new preview, may be missing if updating one.
//
// Notes:
//   - Not passing "enabled": true for previews will result in them not being included in the edit. This may be fixed in the future.
//   - version_string is free-form text, but major.minor or major.minor.patch format is best.
//   - Available download providers can be seen on the asset library fronted.
func CreateAssetHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type UpdateAssetRequest struct {
	Token            string                      `json:"token"`
	Title            string                      `json:"title"`
	Description      string                      `json:"description"`
	CategoryID       string                      `json:"category_id"`
	GodotVersion     string                      `json:"godot_version"`
	VersionString    string                      `json:"version_string"`
	Cost             string                      `json:"cost"`
	DownloadProvider string                      `json:"download_provider"`
	DownloadCommit   string                      `json:"download_commit"`
	BrowseURL        string                      `json:"browse_url"`
	IssuesURL        string                      `json:"issues_url"`
	IconURL          string                      `json:"icon_url"`
	DownloadURL      string                      `json:"download_url"`
	Previews         []UpdateAssetRequestPreview `json:"previews"`
}

type UpdateAssetRequestPreview struct {
	Enabled       bool    `json:"enabled"`
	Operation     string  `json:"operation"`
	Type          *string `json:"type,omitempty"`
	Link          *string `json:"link,omitempty"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
	EditPreviewID *string `json:"edit_preview_id,omitempty"`
}

type UpdateAssetResponse struct {
	ID string `json:"id"`
}

// Update an existing one. Fields are required when creating a new asset, and are optional otherwise. Same for previews -- required when creating a new preview, may be missing if updating one.
//
// Notes:
//   - Not passing "enabled": true for previews will result in them not being included in the edit. This may be fixed in the future.
//   - version_string is free-form text, but major.minor or major.minor.patch format is best.
//   - Available download providers can be seen on the asset library fronted.
func UpdateAssetHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type DeleteAssetRequest struct {
	Token string `json:"token"`
}

type DeleteAssetResponse struct {
	Changed bool `json:"changed"`
}

// Soft-delete an asset. Useable by moderators and the owner of the asset.
func DeleteAssetHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type UndeleteAssetRequest struct {
	Token string `json:"token"`
}

type UndeleteAssetResponse struct {
	Changed bool `json:"changed"`
}

// Revert a deletion of an asset. Useable by moderators and the owner of the asset.
func UndeleteAssetHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

type UpdateSupportLevelRequest struct {
	SupportLevel string `json:"support_level"`
	Token        string `json:"token"`
}

type UpdateSupportLevelResponse struct {
	Changed bool `json:"changed"`
}

// API used by moderators to change the support level of an asset.
func UpdateSupportLevelHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}
