// fs/interface.go
package fs

import (
  "context"
  "io"
  "time"
)

// file's metadata
type FileInfo struct {
  Path string
  IsDir bool
  Size int64
  ModTime time.Time
}

type FileSystem interface {
  WriteFile(ctx context.Context, path string, data []byte) (*FileInfo, error)
  ReadFile(ctx context.Context, path string) (io.ReadCloser, *FileInfo, error)
  Delete(ctx context.Context, path string, recursive bool) error
  MakeDir(ctx context.Context, path string) (*FileInfo, error)
  ListDir(ctx context.Context, path string) ([]*FileInfo, error)
  Stat(ctx context.Context, path string) (*FileInfo, error)
  Copy(ctx context.Context, srcPath, destPath string) (*FileInfo, error)

  // Find files matching a pattern
  // pattern can be prefix / suffix / regex or even prefix and suffix
  Find(ctx context.Context, path string, pattern string) ([]*FileInfo, error)
}
