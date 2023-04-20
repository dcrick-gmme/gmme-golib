/*
//==============================================================================
// GMMEGoLib for Go/Golang
// Copyright (c) 2023, GMM Enterprises, LLC.
// Licensed under the GMM Software License
// All rights reserved
//==============================================================================
//	Author:	David Crickenberger
// -----------------------------------------------------------------------------
// Package:
//		gmme-golib/utils/cmdline
//
// 	Description:
//		Command line processor package.
//
//==============================================================================
*/
package cmdline

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//==============================================================================

// -----------------------------------------------------------------------------
// -- Constants
// -----------------------------------------------------------------------------
/*
const (
	GETDBLOGON = iota + 1
	GETOPTVALUE
	GETOPTVALUEDEF
	GETPATHOPT
	ISOPT
	ISOPT_SETON
	ISOPT_SETOFF
)
*/

// -----------------------------------------------------------------------------
// -- Structs
// -----------------------------------------------------------------------------
type sOptItem struct {
	m_opt     string
	m_val     string
	m_valOrig string
	m_optFile string
	m_tags    []string
}

type sCmdLine struct {
	m_dbgOn  bool
	m_isInit bool
	m_opts   map[string]*sOptItem
}

// -----------------------------------------------------------------------------
// -- Interfaces
// -----------------------------------------------------------------------------
type CmdLine interface {
	AddArgsArray(a_args []string)
	AddArgsFile(a_file string)
	AddArgsLine(a_line string)

	GetOptValue(a_opt string) (string, bool)

	HasOpt(a_opt string) bool
	IsOpt(a_opt string) bool
}

// =============================================================================
// == NewXXX - Constructors
// =============================================================================
func newSOptItem(a_opt string, a_val string, a_optFile string, a_tags []string) *sOptItem {
	l_opt := new(sOptItem)

	l_opt.m_opt = a_opt
	l_opt.m_val = xSubEnv(a_val)
	l_opt.m_valOrig = a_val
	l_opt.m_optFile = a_optFile
	l_opt.m_tags = a_tags

	return l_opt
}

func newSCmdLine() *sCmdLine {
	l_cmdline := new(sCmdLine)

	l_cmdline.m_dbgOn = false
	l_cmdline.m_isInit = true
	l_cmdline.m_opts = make(map[string]*sOptItem)

	return l_cmdline
}
func NewCmdLine() *sCmdLine {
	return newSCmdLine()
}

// =============================================================================
// -- Methods sCmdline and sOptItem
// =============================================================================
// -----------------------------------------------------------------------------
// -- Debug
func (c *sCmdLine) Debug(a_dbg ...bool) bool {
	if len(a_dbg) > 0 {
		c.m_dbgOn = a_dbg[0]
	}

	return c.m_dbgOn
}

func (o *sOptItem) TagExists(a_tag string) bool {
	for _, l_tag := range o.m_tags {
		if l_tag == a_tag {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------
// -- Dump
func (c *sCmdLine) Dump() {
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		fmt.Println("CMDLine: Dump - nothing exists...")
		return
	}

	fmt.Println("CMDLine: Dump - beg")

	//--------------------------------------------------------------------------
	//-- create sorted list of keys of options
	l_keys := make([]string, 0, len(c.m_opts))
	for l_key := range c.m_opts {
		l_keys = append(l_keys, l_key)
	}
	sort.Strings(l_keys)

	//--------------------------------------------------------------------------
	//-- output keys, generating rand object for secret values
	l_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, l_key := range l_keys {
		//----------------------------------------------------------------------
		//-- get opt, format a string, and output
		l_opt := c.m_opts[l_key]

		l_out := strings.Builder{}
		l_out.Grow(128)
	l_val := l_opt.m_val
		if len(l_opt.m_tags)> 0 && l_opt.TagExists("HIDE") {
			l_val = strings.Repeat("*", l_rand.Intn(11)+10)
		}
		lout.WriteString(fmt.Sprintf("   %s == [%s]", l_key, l_val))

	if len(l_opt.m_tags) > 0 {
			l_out.WriteString(fmt.Sprntf("; Tags = %v", l_opt.m_tags))
		}

	fmt.Println(l_out.String())
	}

fmt.Println("CMDLine:: Dump - end")
}
