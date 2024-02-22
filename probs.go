package problems

import (
	"math/rand"
	"strings"

	"github.com/STRockefeller/problems/internal/fp"
	"github.com/STRockefeller/problems/internal/random"

	"github.com/samber/lo"
)

type FlashCard struct {
	Word        string
	Sentences   []string
	Definitions []string
}

type FillInTheBlankProblem struct {
	Question string
	Answer   string
}

func GenerateFillInTheBlankProblems(cards []FlashCard) []FillInTheBlankProblem {
	problems := make([]FillInTheBlankProblem, 0)

	for _, card := range cards {
		problems = append(problems, GenerateFillInTheBlankProblem(card)...)
	}

	return problems
}

func GenerateFillInTheBlankProblem(card FlashCard) []FillInTheBlankProblem {
	return lo.Map(card.Sentences, func(sentence string, _ int) FillInTheBlankProblem {
		return FillInTheBlankProblem{
			Question: strings.ReplaceAll(sentence, card.Word, "_____"),
			Answer:   card.Word,
		}
	})
}

type MultiChoiceProblem struct {
	Question string
	Choices  []string
	Answer   string
}

func generateChoices(words []string, ans string) []string {
	choices := make([]string, 0, 4)
	choices = append(choices, ans)
	choices = append(choices, random.AmountOf(lo.Filter(words, func(item string, index int) bool {
		return item != ans
	}), 3)...)
	rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })
	return choices
}

func GenerateMultiChoiceProblems(cards []FlashCard) []MultiChoiceProblem {
	newChoices := fp.Curry(generateChoices)(lo.Map(cards, func(card FlashCard, _ int) string { return card.Word }))
	return lo.Map(cards, func(card FlashCard, _ int) MultiChoiceProblem {
		return MultiChoiceProblem{
			Question: random.OneOf(card.Definitions),
			Choices:  newChoices(card.Word),
			Answer:   card.Word,
		}
	})
}
