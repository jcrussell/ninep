// Copyright 2009 The Ninep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"bytes"

	"github.com/Harvey-OS/ninep/protocol"
)

type Server struct {
	protocol.NineServer

	trace protocol.Tracer
}

type ServerOpt func(*Server) error

func NewServer(s protocol.NineServer, opts ...ServerOpt) (*Server, error) {
	s2 := &Server{NineServer: s}

	for _, opt := range opts {
		if err := opt(s2); err != nil {
			return nil, err
		}
	}

	return s2, nil
}

func Trace(tracer protocol.Tracer) ServerOpt {
	return func(s *Server) error {
		s.trace = tracer
		return nil
	}
}

func (s *Server) logf(format string, args ...interface{}) {
	if s.trace != nil {
		s.trace(format, args...)
	}
}

func (s *Server) Rversion(msize protocol.MaxSize, version string) (protocol.MaxSize, string, error) {
	s.logf(">>> Tversion %v %v\n", msize, version)
	msize, version, err := s.NineServer.Rversion(msize, version)
	if err == nil {
		s.logf("<<< Rversion %v %v\n", msize, version)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return msize, version, err
}

func (s *Server) Rattach(fid protocol.FID, afid protocol.FID, uname string, aname string) (protocol.QID, error) {
	s.logf(">>> Tattach fid %v,  afid %v, uname %v, aname %v\n", fid, afid,
		uname, aname)
	qid, err := s.NineServer.Rattach(fid, afid, uname, aname)
	if err == nil {
		s.logf("<<< Rattach %v\n", qid)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return qid, err
}

func (s *Server) Rflush(o protocol.Tag) error {
	s.logf(">>> Tflush tag %v\n", o)
	err := s.NineServer.Rflush(o)
	if err == nil {
		s.logf("<<< Rflush\n")
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return err
}

func (s *Server) Rwalk(fid protocol.FID, newfid protocol.FID, paths []string) ([]protocol.QID, error) {
	s.logf(">>> Twalk fid %v, newfid %v, paths %v\n", fid, newfid, paths)
	qid, err := s.NineServer.Rwalk(fid, newfid, paths)
	if err == nil {
		s.logf("<<< Rwalk %v\n", qid)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return qid, err
}

func (s *Server) Ropen(fid protocol.FID, mode protocol.Mode) (protocol.QID, protocol.MaxSize, error) {
	s.logf(">>> Topen fid %v, mode %v\n", fid, mode)
	qid, iounit, err := s.NineServer.Ropen(fid, mode)
	if err == nil {
		s.logf("<<< Ropen %v %v\n", qid, iounit)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return qid, iounit, err
}

func (s *Server) Rcreate(fid protocol.FID, name string, perm protocol.Perm, mode protocol.Mode) (protocol.QID, protocol.MaxSize, error) {
	s.logf(">>> Tcreate fid %v, name %v, perm %v, mode %v\n", fid, name,
		perm, mode)
	qid, iounit, err := s.NineServer.Rcreate(fid, name, perm, mode)
	if err == nil {
		s.logf("<<< Rcreate %v %v\n", qid, iounit)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return qid, iounit, err
}

func (s *Server) Rclunk(fid protocol.FID) error {
	s.logf(">>> Tclunk fid %v\n", fid)
	err := s.NineServer.Rclunk(fid)
	if err == nil {
		s.logf("<<< Rclunk\n")
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return err
}

func (s *Server) Rstat(fid protocol.FID) ([]byte, error) {
	s.logf(">>> Tstat fid %v\n", fid)
	b, err := s.NineServer.Rstat(fid)
	if err == nil {
		dir, _ := protocol.Unmarshaldir(bytes.NewBuffer(b))
		s.logf("<<< Rstat %v\n", dir)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return b, err
}

func (s *Server) Rwstat(fid protocol.FID, b []byte) error {
	dir, _ := protocol.Unmarshaldir(bytes.NewBuffer(b))
	s.logf(">>> Twstat fid %v, %v\n", fid, dir)
	err := s.NineServer.Rwstat(fid, b)
	if err == nil {
		s.logf("<<< Rwstat\n")
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return err
}

func (s *Server) Rremove(fid protocol.FID) error {
	s.logf(">>> Tremove fid %v\n", fid)
	err := s.NineServer.Rremove(fid)
	if err == nil {
		s.logf("<<< Rremove\n")
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return err
}

func (s *Server) Rread(fid protocol.FID, o protocol.Offset, c protocol.Count) ([]byte, error) {
	s.logf(">>> Tread fid %v, off %v, count %v\n", fid, o, c)
	b, err := s.NineServer.Rread(fid, o, c)
	if err == nil {
		s.logf("<<< Rread %v\n", len(b))
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return b, err
}

func (s *Server) Rwrite(fid protocol.FID, o protocol.Offset, b []byte) (protocol.Count, error) {
	s.logf(">>> Twrite fid %v, off %v, count %v\n", fid, o, len(b))
	c, err := s.NineServer.Rwrite(fid, o, b)
	if err == nil {
		s.logf("<<< Rwrite %v\n", c)
	} else {
		s.logf("<<< Error %v\n", err)
	}
	return c, err
}
