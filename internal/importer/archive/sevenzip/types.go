package sevenzip

import (
	"context"

	"github.com/WhispersOfJ/bearmount/internal/importer/archive"
	"github.com/WhispersOfJ/bearmount/internal/importer/parser"
	metapb "github.com/WhispersOfJ/bearmount/internal/metadata/proto"
	"github.com/WhispersOfJ/bearmount/internal/progress"
)

// Content is an alias for archive.Content
type Content = archive.Content

// NestedSource is an alias for archive.NestedSource
type NestedSource = archive.NestedSource

// Processor interface for analyzing 7zip content from NZB data
type Processor interface {
	// AnalyzeSevenZipContentFromNzb analyzes a 7zip archive directly from NZB data
	// without downloading. Returns an array of Content with file metadata and segments.
	// password parameter is used to unlock password-protected 7zip archives.
	// progressTracker is used to report progress during analysis.
	AnalyzeSevenZipContentFromNzb(ctx context.Context, sevenZipFiles []parser.ParsedFile, password string, progressTracker *progress.Tracker) ([]Content, error)
	// CreateFileMetadataFromSevenZipContent creates FileMetadata from Content for the metadata
	// system. This is used to convert Content into the protobuf format used by the metadata system.
	CreateFileMetadataFromSevenZipContent(content Content, sourceNzbPath string, releaseDate int64, nzbdavId string) *metapb.FileMetadata
}
