package timeparser_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/timeparser"
	"time"
)

var _ = Describe("Timeparser", func() {

	testYear := 2023
	testMonth := 12

	startDay := 18
	endDay := 20

	startDate := time.Date(testYear, time.Month(testMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(testYear, time.Month(testMonth), endDay, 0, 0, 0, 0, time.UTC)

	startString := "12/18"
	endString := "12/20"
	input := fmt.Sprintf("%s-%s", startString, endString)

	It("can get the start and end date string", func() {
		start, end := timeparser.GetStartAndEndDateString(input)
		Expect(start).To(BeEquivalentTo(startString))
		Expect(end).To(BeEquivalentTo(endString))
	})

	It("can parse a report request with a date", func() {
		actualStart, actualEnd := timeparser.FormatReportRange(startString, endString)
		Expect(actualStart).NotTo(BeNil())
		Expect(*actualStart).To(BeEquivalentTo(startDate))
		Expect(actualEnd).NotTo(BeNil())
		Expect(*actualEnd).To(BeEquivalentTo(endDate))
	})

	It("can parse a report request with a date and year", func() {
		actualStart, actualEnd := timeparser.GetStartAndEndDateFromMessage(input)
		Expect(actualStart).NotTo(BeNil())
		Expect(*actualStart).To(BeEquivalentTo(startDate))
		Expect(actualEnd).NotTo(BeNil())
		Expect(*actualEnd).To(BeEquivalentTo(endDate))
	})

})
