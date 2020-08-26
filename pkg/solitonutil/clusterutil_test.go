package solitonAutomatautil

import (
	"github.com/YosiSF/check"
)

type utilSuite struct{}

var _ = check.Suite(&utilSuite{})

func (s *utilSuite) TestMultiDirAbs(c *check.C) {
	paths := MultiDirAbs("milevadb", "")
	c.Assert(len(paths), check.Equals, 0)

	paths = MultiDirAbs("milevadb", " ")
	c.Assert(len(paths), check.Equals, 0)

	paths = MultiDirAbs("milevadb", "a ")
	c.Assert(len(paths), check.Equals, 1)
	c.Assert(paths[0], check.Equals, "/home/milevadb/a")

	paths = MultiDirAbs("milevadb", "a , /tmp/b")
	c.Assert(len(paths), check.Equals, 2)
	c.Assert(paths[0], check.Equals, "/home/milevadb/a")
	c.Assert(paths[1], check.Equals, "/tmp/b")
}
