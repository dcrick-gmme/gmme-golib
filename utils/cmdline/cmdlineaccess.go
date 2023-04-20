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
	"strings"
)

// -----------------------------------------------------------------------------
// -- GetOptValue()
// -----------------------------------------------------------------------------
func (c *sCmdLine) GetOptValue(a_opt string) (string, bool) {
	l_func := "DBG-utils.cmdline.GetOptValue:: opt ="

	//--------------------------------------------------------------------------
	//-- if debug on dump a_args, and setup defer
	if c.m_dbgOn {
		fmt.Println(l_func, a_opt, "- beg:")

		defer func() {
			fmt.Println(l_func, a_opt, "- end:")
		}()
	}

	//--------------------------------------------------------------------------
	//-- see if opt exists
	l_val := ""
	l_opt, l_found := c.m_opts[strings.ToUpper(a_opt)]
	if l_found {
		l_val = l_opt.m_val
	}
	if c.m_dbgOn {
		fmt.Println(l_func, a_opt, "- found =", l_found, ", value =", l_val)
	}

	return l_val, l_found
}
