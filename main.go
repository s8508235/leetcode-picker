package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	flag "github.com/spf13/pflag"
)

const (
	leetcodeAPIURL        = "https://leetcode.com/api/problems/all/"
	leetcodeProblemPrefix = "https://leetcode.com/problems"
	leetcodeGraphQL       = "https://leetcode.com/graphql"
)

var (
	data           API
	problems       []Problem
	easyProblems   []Problem
	mediumProblems []Problem
	hardProblems   []Problem
	level          string
	credit         int
)

func init() {
	flag.StringVarP(&level, "level", "l", "all", "problem level for all,easy,medium,hard")
	flag.IntVarP(&credit, "credit", "c", 5, "problem credit with ")
}

func setup() error {
	flag.Parse()
	resp, err := http.Get(leetcodeAPIURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func main() {
	if err := setup(); err != nil {
		fmt.Println("Setup error")
	}

	for _, stat := range data.StatPair {
		if stat.Stat.PaidToWin {
			continue
		}
		problems = append(problems, stat.Stat)

		switch Level(stat.Difficulty.Level) {
		case Easy:
			easyProblems = append(easyProblems, stat.Stat)
		case Medium:
			mediumProblems = append(mediumProblems, stat.Stat)
		case Hard:
			hardProblems = append(hardProblems, stat.Stat)
		}
	}

	switch level {
	case "a", "all":
		problem, err := getProblem(problems)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(problem)
		fmt.Printf("%s/%s", leetcodeProblemPrefix, problem.TitleSlug)
	case "e", "easy":
		problem, err := getProblem(easyProblems)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(problem)
		fmt.Printf("%s/%s", leetcodeProblemPrefix, problem.TitleSlug)
	case "m", "medium":
		problem, err := getProblem(mediumProblems)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(problem)
		fmt.Printf("%s/%s", leetcodeProblemPrefix, problem.TitleSlug)
	case "h", "hard":
		problem, err := getProblem(hardProblems)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(problem)
		fmt.Printf("%s/%s", leetcodeProblemPrefix, problem.TitleSlug)
	default:
		fmt.Println("Wrong flag! Please input one of all/easy/medium/hard")
	}
}

func getProblem(problems []Problem) (Problem, error) {
	if len(problems) <= 0 {
		return Problem{}, errors.New("No enough problem")
	}
	problemNo, err := getRandomNumber(int64(len(problems)))

	if err != nil {
		fmt.Println(err)
		return Problem{}, err
	}
	return problems[problemNo], nil
}

// fetch("https://leetcode.com/graphql", {
//   "headers": {
//     "accept": "*/*",
//     "accept-language": "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6",
//     "content-type": "application/json",
//     "sec-fetch-dest": "empty",
//     "sec-fetch-mode": "cors",
//     "sec-fetch-site": "same-origin",
//     "x-csrftoken": "cGm3bVZCQIST5y5J4qfpNA8SxABgbrELiFkD3MG9GjgVyDa3tPTB97avFXiJnyDs",
//     "x-newrelic-id": "UAQDVFVRGwEAXVlbBAg="
//   },
//   "referrer": "https://leetcode.com/problems/single-number-ii/",
//   "referrerPolicy": "strict-origin-when-cross-origin",
//   "body": "{\"operationName\":\"questionData\",\"variables\":{\"titleSlug\":\"single-number-ii\"},\"query\":\"query questionData($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    boundTopicId\\n    title\\n    titleSlug\\n    content\\n    translatedTitle\\n    translatedContent\\n    isPaidOnly\\n    difficulty\\n    likes\\n    dislikes\\n    isLiked\\n    similarQuestions\\n    contributors {\\n      username\\n      profileUrl\\n      avatarUrl\\n      __typename\\n    }\\n    topicTags {\\n      name\\n      slug\\n      translatedName\\n      __typename\\n    }\\n    companyTagStats\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n      __typename\\n    }\\n    stats\\n    hints\\n    solution {\\n      id\\n      canSeeDetail\\n      paidOnly\\n      __typename\\n    }\\n    status\\n    sampleTestCase\\n    metaData\\n    judgerAvailable\\n    judgeType\\n    mysqlSchemas\\n    enableRunCode\\n    enableTestMode\\n    enableDebugger\\n    envInfo\\n    libraryUrl\\n    adminUrl\\n    __typename\\n  }\\n}\\n\"}",
//   "method": "POST",
//   "mode": "cors",
//   "credentials": "include"
// });
