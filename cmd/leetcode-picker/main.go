package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/machinebox/graphql"
	"github.com/s8508235/leetcode-picker/pkg/entity"
	"github.com/s8508235/leetcode-picker/pkg/util"
	flag "github.com/spf13/pflag"
)

const (
	leetcodeApiUrl        = "https://leetcode.com/api/problems/all/"
	leetcodeProblemPrefix = "https://leetcode.com/problems"
	leetcodeGraphqlPrefix = "https://leetcode.com/graphql/"
)

var (
	data           entity.API
	problems       []entity.Problem
	easyProblems   []entity.Problem
	mediumProblems []entity.Problem
	hardProblems   []entity.Problem
	level          string
	rating         int
)

func setup() error {
	flag.Parse()
	if _, err := os.Stat("./problems.json"); errors.Is(err, os.ErrNotExist) {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3239.108 Safari/537.36"),
		)
		ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()
		// use chromedp.WithDebugf to log the CDP messages
		// ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
		ctx, cancel = chromedp.NewContext(ctx)
		defer cancel()
		// create a timeout
		ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		// navigate to a page, wait for an element, click
		var value string
		// var str string
		err := chromedp.Run(ctx,
			chromedp.Navigate(leetcodeApiUrl),
			chromedp.Sleep(500*time.Millisecond),
			// chromedp.ActionFunc(func(ctx context.Context) error {
			// 	node, err := dom.GetDocument().Do(ctx)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			// 	return err
			// }),
			// wait for body element is visible (ie, page is loaded)
			chromedp.Text(`body`, &value, chromedp.NodeVisible),
		)
		// os.WriteFile("chromedp.html", []byte(str), 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = json.Unmarshal([]byte(value), &data)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	// work around forbidden issue
	problemFile, err := os.OpenFile("./problems.json", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.NewDecoder(problemFile).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	flag.StringVarP(&level, "level", "l", "all", "problem level for all, easy, medium, normal, hard")
	flag.IntVarP(&rating, "rating", "r", 0, "problem rating")
	// st := time.Now()
	if err := setup(); err != nil {
		fmt.Println("Setup error", err)
	}
	// fmt.Println("time elapsed setup", float64(time.Since(st).Milliseconds())/1000.0, "secs")
	for _, stat := range data.StatPair {
		if stat.PaidToWin {
			continue
		}
		problems = append(problems, stat.Stat)

		switch entity.Level(stat.Difficulty.Level) {
		case entity.Easy:
			easyProblems = append(easyProblems, stat.Stat)
		case entity.Medium:
			mediumProblems = append(mediumProblems, stat.Stat)
		case entity.Hard:
			hardProblems = append(hardProblems, stat.Stat)
		}
	}
	// fmt.Println("time elapsed problem classification", float64(time.Since(st).Milliseconds())/1000.0, "secs")
	var problem entity.Problem
	var err error
pick:
	for {
		switch level {
		case "a", "all":
			problem, err = getProblem(problems)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "e", "easy":
			problem, err = getProblem(easyProblems)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "m", "medium":
			problem, err = getProblem(mediumProblems)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "h", "hard":
			problem, err = getProblem(hardProblems)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "n", "normal":
			problem, err = getProblem(append(mediumProblems, hardProblems...))
			if err != nil {
				fmt.Println(err)
				return
			}
		default:
			fmt.Println("Wrong flag! Please input one of all/easy/medium/hard/normal")
			return
		}
		// fmt.Println("time elapsed pick a problem", float64(time.Since(st).Milliseconds())/1000.0, "secs")
		if entity.Rating(rating) == entity.Negative {
			break
		}
		likes, dislikes, err := getProblemLike(problem.TitleSlug)
		if err != nil {
			fmt.Println(err)
			return
		}
		rate := float64(likes) / float64(dislikes+likes)
		// fmt.Println(likes, dislikes, rate)
		// fmt.Println("time elapsed fetch likes of a problem", float64(time.Since(st).Milliseconds())/1000.0, "secs")
		switch entity.Rating(rating) {
		case entity.MostlyNegative:
			if rate > 0.2 {
				break pick
			}
			continue pick
		case entity.Mixed:
			if rate > 0.4 {
				break pick
			}
			continue pick
		case entity.MostlyPositive:
			if rate > 0.7 {
				break pick
			}
			continue pick
		case entity.Positive:
			if rate > 0.8 {
				break pick
			}
			continue pick
		case entity.OverwhelminglyPositive:
			if rate > 0.95 {
				break pick
			}
			continue pick
		}
	}
	fmt.Printf("%s/%s\n", leetcodeProblemPrefix, problem.TitleSlug)
}

func getProblem(problems []entity.Problem) (entity.Problem, error) {
	if len(problems) <= 0 {
		return entity.Problem{}, errors.New("not enough problem")
	}

	problemNo, err := util.GetRandomNumber(int64(len(problems)))
	if err != nil {
		fmt.Println(err)
		return entity.Problem{}, err
	}
	return problems[problemNo], nil
}

func getProblemLike(problemSlug string) (int, int, error) {

	client := graphql.NewClient(leetcodeGraphqlPrefix)

	// make a request
	req := graphql.NewRequest(`
	query questionTitle($titleSlug: String!) {
		question(titleSlug: $titleSlug) {
			titleSlug
			difficulty
			likes
			dislikes
		}
	      }
	`)
	req.Var("titleSlug", problemSlug)
	var respData entity.ProblemGraphql
	if err := client.Run(context.Background(), req, &respData); err != nil {
		fmt.Println("graphql error", err)
		return -1, -1, err
	}

	if respData.Question.TitleSlug == problemSlug {
		return respData.Question.Likes, respData.Question.DisLikes, nil
	}
	return -1, -1, errors.New("wrong problem")
}
