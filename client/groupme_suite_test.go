package groupme_test

import (
	"testing"
)

func TestGroupme(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Groupme Suite")
}
