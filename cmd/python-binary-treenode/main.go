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

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

const treeNodeStr = "TreeNode("
const noneStr = "None"

func bstFromPreOrder(logger *slog.Logger, numList []string) (*treeNode, error) {
	root := new(treeNode)
	if strings.TrimSpace(numList[0]) == "null" {
		return nil, nil
	}
	value, parseErr := strconv.Atoi(strings.TrimSpace(numList[0]))
	if parseErr != nil {
		logger.Error("parse tree node value failed",
			slog.String("value", numList[0]), slog.Any("err", parseErr))
		return nil, parseErr
	}
	root.Val = value
	stack := []*treeNode{root}
	for index, val := range numList[1:] {
		val := strings.TrimSpace(val)
		if val == "null" {
			// left node
			if index%2 == 0 {
				stack[len(stack)-1].Left = nil
			} else {
				stack = stack[1:]
			}
			continue
		}
		value, parseErr := strconv.Atoi(val)
		if parseErr != nil {
			logger.Error("parse tree node value failed",
				slog.String("value", val),
				slog.Any("err", parseErr))
			return nil, parseErr
		}
		if index%2 == 0 {
			stack[0].Left = &treeNode{Val: value}
			stack = append(stack, stack[0].Left)
		} else {
			last := stack[0]
			last.Right = &treeNode{Val: value}
			stack = stack[1:]
			stack = append(stack, last.Right)
		}

	}
	return root, nil
}

func traverse(stringBuilder *strings.Builder, node *treeNode) {
	if node == nil {
		stringBuilder.WriteRune(',')
		stringBuilder.WriteString(noneStr)
		return
	}
	if stringBuilder.Len() != 0 {
		stringBuilder.WriteRune(',')
	}
	stringBuilder.WriteString(treeNodeStr)
	stringBuilder.WriteString(strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		stringBuilder.WriteRune(')')
		return
	}
	// if node.Left != nil {
	traverse(stringBuilder, node.Left)
	if node.Right != nil {
		// stringBuilder.WriteString("with right" + strconv.Itoa(node.Right.Val))
		// fmt.Println(node.Val, "right", node.Right.Val)
		traverse(stringBuilder, node.Right)
	}
	stringBuilder.WriteRune(')')
}

// make leetcode array to binary tree in python for building test cases
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true}))
	var pythonArrayStr string
	flag.StringVarP(&pythonArrayStr, "array_string", "a", "[]",
		"leetcode array string to convert to python tree node e.g. [1,null,3]")
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
	if head, err := bstFromPreOrder(logger, pythonArray); err != nil {
		logger.Error("compile regexp failed", slog.Any("err", err))
		return
	} else {
		traverse(sb, head)
	}
	fmt.Println(sb.String())
}
