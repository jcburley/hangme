package main

import (
	// _ "archive/tar"
	// _ "archive/zip"
	// _ "bufio"
	// _ "bytes"
	// _ "compress/bzip2"
	// _ "compress/flate"
	// _ "compress/gzip"
	// _ "compress/lzw"
	// _ "compress/zlib"
	// _ "container/heap"
	// _ "container/list"
	// _ "container/ring"
	// _ "context"
	// _ "crypto"
	// _ "crypto/aes"
	// _ "crypto/cipher"
	// _ "crypto/des"
	// _ "crypto/dsa"
	// _ "crypto/ecdsa"
	// _ "crypto/ed25519"
	// _ "crypto/elliptic"
	// _ "crypto/hmac"
	// _ "crypto/md5"
	// _ "crypto/rand"
	// _ "crypto/rc4"
	// _ "crypto/rsa"
	// _ "crypto/sha1"
	// _ "crypto/sha256"
	// _ "crypto/sha512"
	// _ "crypto/subtle"
	// _ "crypto/tls"
	// _ "crypto/x509"
	// _ "crypto/x509/pkix"
	// _ "database/sql"
	// _ "database/sql/driver"
	// _ "debug/dwarf"
	// _ "debug/elf"
	// _ "debug/gosym"
	// _ "debug/macho"
	// _ "debug/pe"
	// _ "debug/plan9obj"
	// _ "encoding"
	// _ "encoding/ascii85"
	// _ "encoding/asn1"
	// _ "encoding/base32"
	// _ "encoding/base64"
	// _ "encoding/binary"
	// _ "encoding/csv"
	// _ "encoding/gob"
	// _ "encoding/hex"
	// _ "encoding/json"
	// _ "encoding/pem"
	// _ "encoding/xml"
	// _ "errors"
	// _ "expvar"
	// _ "flag"
	"fmt"
	// _ "go/ast"
	// _ "go/build"
	// _ "go/constant"
	// _ "go/doc"
	// _ "go/format"
	// _ "go/importer"
	// _ "go/parser"
	// _ "go/printer"
	// _ "go/scanner"
	// _ "go/token"
	// _ "go/types"
	// _ "hash"
	// _ "hash/adler32"
	// _ "hash/crc32"
	// _ "hash/crc64"
	// _ "hash/fnv"
	// _ "hash/maphash"
	// _ "html"
	// _ "html/template"
	// _ "image"
	// _ "image/color"
	// _ "image/color/palette"
	// _ "image/draw"
	// _ "image/gif"
	// _ "image/jpeg"
	// _ "image/png"
	// _ "index/suffixarray"
	// _ "io"
	// _ "io/ioutil"
	// _ "log"
	// _ "log/syslog"
	// _ "math"
	// _ "math/big"
	// _ "math/bits"
	// _ "math/rand"
	// _ "mime"
	// _ "mime/multipart"
	// _ "mime/quotedprintable"
	// _ "net"
	// _ "net/http"
	// _ "net/http/cgi"
	// _ "net/http/cookiejar"
	// _ "net/http/fcgi"
	// _ "net/http/httptest"
	// _ "net/http/httptrace"
	// _ "net/http/httputil"
	// _ "net/http/pprof"
	// _ "net/mail"
	// _ "net/rpc"
	// _ "net/rpc/jsonrpc"
	// _ "net/smtp"
	// _ "net/textproto"
	// _ "net/url"
	"os"
	"os/exec"
	// _ "os/signal"
	// _ "os/user"
	// _ "path"
	// _ "path/filepath"
	_ "plugin"
	// _ "reflect"
	// _ "regexp"
	// _ "regexp/syntax"
	"runtime"
	// _ "runtime/debug"
	// _ "runtime/pprof"
	// _ "runtime/trace"
	// _ "sort"
	"strconv"
	// _ "strings"
	// _ "sync"
	// _ "sync/atomic"
	// _ "syscall"
	// _ "testing"
	// _ "testing/iotest"
	// _ "testing/quick"
	// _ "text/scanner"
	// _ "text/tabwriter"
	// _ "text/template"
	// _ "text/template/parse"
	// _ "time"
	// _ "unicode"
	// _ "unicode/utf16"
	// _ "unicode/utf8"
	// _ "unsafe"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: hangme <iterations>  # 0 means exit immediately\n")
		os.Exit(99)
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse argument as number: %v\n", err)
		os.Exit(1)
	}
	if count == 0 {
		os.Exit(0)
	}
	fmt.Printf("%s\n", runtime.Version())
	for i := 0; i < count; i++ {
		cmd := exec.Command("./hangme", "0")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	}
	return
}
