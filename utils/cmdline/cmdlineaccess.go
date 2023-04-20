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
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		return "", false
	}

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

// -----------------------------------------------------------------------------
// -- HasOpt() and IsOpt()
// -----------------------------------------------------------------------------
func (c *sCmdLine) HasOpt(a_opt string) bool {
	return c.IsOpt(a_opt)
}
func (c *sCmdLine) IsOpt(a_opt string) bool {
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		return false
	}

	//--------------------------------------------------------------------------
	//-- see if opt exists
	_, l_found := c.m_opts[strings.ToUpper(a_opt)]

	return l_found
}

/*
#---------------------------------------------------------------------------
#-- isOpt
#---------------------------------------------------------------------------
def IsOpt(self, a_opt):

	if not self.m_isInit: return None

	if a_opt.upper() not in self.m_opts:
		return False
	return True
*/
