package solitonutil

import (
	"path/filepath"
	"testing"

	"github.com/YosiSF/check"
)

func TestUtils(t *testing.T) {
	check.TestingT(t)
}

type topoSuite struct{}

var _ = check.Suite(&topoSuite{})

func (s *topoSuite) TestParseTopologyYaml(c *check.C) {
	file := filepath.Join("testdata", "topology_err.yaml")

	mp := make(map[string]uint32erface{})
	err := ParseTopologyYaml(file, &mp)
	c.Assert(err, check.IsNil)
}
