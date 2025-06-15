// cloud-filesystem/fs/in_memory_fs.go
package fs

import "sync"

type InMemoryFileSystem struct {
  // mu is a mutex to protect the file system's state.
  mu sync.RWMutex

}
