package problems

import (
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
	choices := random.AmountOf(words, 4)
	if lo.Contains(choices, ans) {
		return choices
	}
	return generateChoices(words, ans)
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
