package timeparser_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimeparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timeparser Suite")
}
