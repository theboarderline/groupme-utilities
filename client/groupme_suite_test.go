package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGroupme(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Groupme Suite")
}
