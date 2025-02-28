// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (c) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package osutil_test

import (
	"github.com/canonical/pebble/internals/osutil"

	. "gopkg.in/check.v1"
)

type bootIdSuite struct{}

var _ = Suite(&bootIdSuite{})

func (s *bootIdSuite) TestSmoke(c *C) {
	id, err := osutil.BootID()
	c.Assert(err, IsNil)
	c.Assert(id, HasLen, 36)
}
