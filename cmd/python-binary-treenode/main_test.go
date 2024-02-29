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
		{[]string{"1", "null", "2"}, ",TreeNode(1,None,TreeNode(2))", nil},
		{[]string{"1"}, ",TreeNode(1)", nil},
		{[]string{"5", "9", "1", "3", "6", "7"},
			",TreeNode(5,TreeNode(9,TreeNode(3),TreeNode(6)),TreeNode(1,TreeNode(7)))", nil},
		{[]string{"1", "10", "4", "3", "null", "7"},
			",TreeNode(1,TreeNode(10,TreeNode(3)),TreeNode(4,TreeNode(7)))", nil},
	}
	sb := new(strings.Builder)
	for _, testCase := range testCases {
		sb.Reset()
		head, err := bstFromPreOrder(logger, testCase.numList)
		if testCase.expectErr != err {
			t.Logf("expect error %v but with %v", testCase.expectErr, err)
			t.Fail()
			return
		}
		traverse(sb, head)
		if sb.String() != testCase.expectStr {
			t.Logf("expect output %s but with %s", testCase.expectStr, sb.String())
			t.Fail()
			return
		}
	}
}
