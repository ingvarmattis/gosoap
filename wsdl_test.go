package gosoap

import (
	"testing"

	"gotest.tools/assert"
)

func TestFaultString(t *testing.T) {
	var testCases = []struct {
		description      string
		fault            *Fault
		expectedFaultStr string
	}{
		{
			description: "success case: fault string",
			fault: &Fault{
				Code:        "soap:SERVER",
				Description: "soap exception",
				Detail:      "soap detail",
			},
			expectedFaultStr: "[soap:SERVER]: soap exception | Detail: soap detail",
		},
	}

	for _, testCase := range testCases {
		t.Logf("running %v testCase", testCase.description)

		faultStr := testCase.fault.String()
		assert.Equal(t, testCase.expectedFaultStr, faultStr)
	}
}
