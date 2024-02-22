package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFillInTheBlankProblems(t *testing.T) {
	cards := []FlashCard{
		{Word: "apple", Sentences: []string{"I have an apple", "Do you like apples?"}, Definitions: []string{"A fruit"}},
		{Word: "table", Sentences: []string{"Put it on the table", "The table is made of wood"}, Definitions: []string{"A piece of furniture"}},
	}
	expected := []FillInTheBlankProblem{
		{Question: "I have an _____", Answer: "apple"},
		{Question: "Do you like _____s?", Answer: "apple"},
		{Question: "Put it on the _____", Answer: "table"},
		{Question: "The _____ is made of wood", Answer: "table"},
	}

	result := GenerateFillInTheBlankProblems(cards)

	if len(result) != len(expected) {
		t.Errorf("Expected %d problems, but got %d", len(expected), len(result))
	}

	for i, problem := range result {
		if problem != expected[i] {
			t.Errorf("Problem %d: Expected %+v, but got %+v", i+1, expected[i], problem)
		}
	}
}

func TestGenerateMultiChoiceProblems(t *testing.T) {
	cards := []FlashCard{
		{
			Word:        "apple",
			Definitions: []string{"A fruit"},
		},
		{
			Word:        "banana",
			Definitions: []string{"A curved yellow fruit"},
		},
		{
			Word:        "dog",
			Definitions: []string{"A domesticated mammal and common household pet", "Known for their loyalty and friendly demeanor"},
		},
		{
			Word:        "computer",
			Definitions: []string{"An electronic device used for processing and storing data", "Used for work, entertainment, and communication"},
		},
		{
			Word:        "chemistry",
			Definitions: []string{"The scientific study of matter and its properties", "Involves the analysis of atoms, molecules, and chemical reactions"},
		},
		{
			Word:        "history",
			Definitions: []string{"The study of past events and how they influence the present", "Can include topics like politics, wars, and culture"},
		},
	}
	res := GenerateMultiChoiceProblems(cards)
	assert.Len(t, res, len(cards))
	for i, prob := range res {
		assert.Equal(t, cards[i].Word, prob.Answer)
		assert.Contains(t, cards[i].Definitions, prob.Question)
		assert.Contains(t, prob.Choices, prob.Answer)
	}
}
