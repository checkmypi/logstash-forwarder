package fs

import (
	"os"
	"bytes"
	"syscall"
	"encoding/binary"
	"github.com/elasticsearch/kriterium/panics"
)

// determine the oid per OS specific FileInfo
// Encodes the tuple (dev,ino) as a 12 byte []byte slice.
func oid(info os.FileInfo) fsoid {
	if info == nil {
		panic("BUG - info is nil")
	}
	var buf bytes.Buffer

	fstat := info.Sys().(*syscall.Stat_t)
	ino, dev := fstat.Ino, fstat.Dev

	e := binary.Write(&buf, binary.BigEndian, dev)
	panics.OnError(e, "binary.Write", "device", dev)

	e = binary.Write(&buf, binary.BigEndian, ino)
	panics.OnError(e, "binary.Write", "inode", ino)

	return buf.Bytes()
}

func ignoredFiles() []string {
	return []string{"."}
}
