package smoke_test

import (
	"flag"
	"testing"
	"time"

	"github.com/onsi/gomega/format"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var Builder string

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSmoke(t *testing.T) {
	format.MaxLength = 0
	Expect := NewWithT(t).Expect

	flag.Parse()

	Expect(Builder).NotTo(Equal(""))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Smoke", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("Node.js", testNodejs)
	suite("Java", testJava)
	suite.Run(t)
}
