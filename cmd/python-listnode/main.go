package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

type listNode struct {
	Val  int
	Next *listNode
}

const listNodeStr = "ListNode("

func listNodeFromArray(logger *slog.Logger, numList []string) (*listNode, error) {
	root := new(listNode)
	if numList[0] == "null" {
		return nil, nil
	}
	value, parseErr := strconv.Atoi(strings.TrimSpace(numList[0]))
	if parseErr != nil {
		logger.Error("parse tree node value failed",
			slog.String("value", numList[0]), slog.Any("err", parseErr))
		return nil, parseErr
	}
	current := root
	for _, val := range numList[1:] {
		value, parseErr := strconv.Atoi(strings.TrimSpace(val))
		if parseErr != nil {
			logger.Error("parse tree node value failed",
				slog.String("value", val),
				slog.Any("err", parseErr))
			return nil, parseErr
		}
		current.Next = &listNode{Val: value}
		current = current.Next
	}
	root.Val = value
	return root, nil
}

func iterate(stringBuilder *strings.Builder, head *listNode) {
	if head == nil {
		return
	}
	if stringBuilder.Len() != 0 {
		stringBuilder.WriteRune(',')
	}
	stringBuilder.WriteString(listNodeStr)
	stringBuilder.WriteString(strconv.Itoa(head.Val))
	iterate(stringBuilder, head.Next)
	stringBuilder.WriteRune(')')
}

// make leetcode array to binary tree in python for building test cases
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true}))
	var pythonArrayStr string
	flag.StringVarP(&pythonArrayStr, "array_string", "a", "[]",
		"leetcode array string to convert to python list node e.g. [1,2,3]")
	flag.Parse()
	if len(pythonArrayStr) == 2 && pythonArrayStr == "[]" {
		var readErr error
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Array to tree node: ")
		pythonArrayStr, readErr = reader.ReadString('\n')
		if readErr != nil {
			logger.Error("read reader string failed", slog.Any("err", readErr))
			return
		}
	}
	pythonArrayStr = strings.TrimSpace(pythonArrayStr)
	// sanitize input
	if !strings.HasPrefix(pythonArrayStr, "[") || !strings.HasSuffix(pythonArrayStr, "]") {
		logger.Error("not a string format with starting and ending square brackets [1,2,3]")
		return
	}
	pythonArrayStr = pythonArrayStr[1 : len(pythonArrayStr)-1]
	re, err := regexp.Compile(`^(\s*-?\d+,|\s*null,)*(\s*-?\d+|\s*null)$`)
	if err != nil {
		logger.Error("compile regexp failed", slog.Any("err", err))
		return
	}
	if !re.MatchString(pythonArrayStr) {
		logger.Error("not a string format with starting and ending square brackets [1,null,3]")
		return
	}
	logger.Info("process input", slog.Any("array", pythonArrayStr))
	pythonArray := strings.Split(pythonArrayStr, ",")

	// build tree
	sb := new(strings.Builder)
	if head, err := listNodeFromArray(logger, pythonArray); err != nil {
		logger.Error("compile regexp failed", slog.Any("err", err))
		return
	} else {
		iterate(sb, head)
	}
	fmt.Println(sb.String())
}
