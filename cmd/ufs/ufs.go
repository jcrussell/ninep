// UFS is a userspace server which exports a filesystem over 9p2000.
//
// By default, it will export / over a TCP on port 5640 under the username
// of "harvey".
package main

import (
	"flag"
	"log"
	"net"

	"github.com/Harvey-OS/ninep/filesystem"
	"github.com/Harvey-OS/ninep/protocol"
)

var (
	ntype = flag.String("ntype", "tcp4", "Default network type")
	naddr = flag.String("addr", ":5640", "Network address")
	root  = flag.String("root", "/", "filesystem root")
	debug = flag.Bool("debug", false, "enable debug messages")
)

func main() {
	flag.Parse()

	ln, err := net.Listen(*ntype, *naddr)
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	fs, err := ufs.NewServer(ufs.Root(*root), func(fs *ufs.FileServer) error {
		if *debug {
			return ufs.Trace(log.Printf)(fs)
		}

		return nil
	})

	var ninefs protocol.NineServer = fs
	if *debug {
		ninefs = fs.Debug()
	}

	s, err := protocol.NewServer(ninefs, func(s *protocol.Server) error {
		if *debug {
			s.Trace = log.Printf
		}

		return nil
	})

	if err := s.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
