package main

import (
	"io"
	"log/slog"
	"strings"
	"testing"
)

func TestBstFromPreOrder(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	testCases := []struct {
		numList   []string
		expectStr string
		expectErr error
	}{
		{[]string{"1", "2", "3"}, "ListNode(1,ListNode(2,ListNode(3)))", nil},
		{[]string{"1"}, "ListNode(1)", nil},
		{[]string{"5", "9", "1", "3", "6", "7"},
			"ListNode(5,ListNode(9,ListNode(1,ListNode(3,ListNode(6,ListNode(7))))))", nil},
	}
	sb := new(strings.Builder)
	for _, testCase := range testCases {
		sb.Reset()
		head, err := listNodeFromArray(logger, testCase.numList)
		if testCase.expectErr != err {
			t.Logf("expect error %v but with %v", testCase.expectErr, err)
			t.Fail()
			return
		}
		iterate(sb, head)
		if sb.String() != testCase.expectStr {
			t.Logf("expect output %s but with %s", testCase.expectStr, sb.String())
			t.Fail()
			return
		}
	}
}
