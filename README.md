# Flash Card Problem Generator

This code defines two types of problems:

1. Fill in the blank problem
2. Multiple choice problem

The code generates these problems from a slice of `FlashCard`, where each card is an instance of the type `FlashCard`.

## Type Definitions

### `FlashCard`

The struct `FlashCard` has three fields:

- `Word`: A string indicating the word to be learned.
- `Sentences`: a slice of strings indicating sentences containing the word.
- `Definitions`: a slice of strings indicating the definitions of the given word.

### `FillInTheBlankProblem`

The struct `FillInTheBlankProblem` has two fields:

- `Question`: A string with some words replaced by blanks.
- `Answer`: A string indicating the correct answer.

### `MultiChoiceProblem`

The struct `MultiChoiceProblem` has three fields:

- `Question`: A string indicating the definition of the word.
- `Choices`: a slice of strings representing multiple choices for the question.
- `Answer`: A string indicating the correct answer.

## Functions

### `GenerateFillInTheBlankProblems(cards []FlashCard) []FillInTheBlankProblem`

This function takes a slice of `FlashCard`s, and returns a slice of `FillInTheBlankProblem`s generated from the sentences in the `FlashCard`s. It maps over each `FlashCard` and generates problems for each sentence in `Sentences`.

### `GenerateFillInTheBlankProblem(card FlashCard) []FillInTheBlankProblem`

This function generates fill in the blank problem for a single `FlashCard`.

### `generateChoices(words []string, ans string) []string`

This recursive function selects 4 random choices from the slice of words and return them as a slice. The correct answer is also included among the choices. If the correct answer is not among randomly selected words, the function recursively calls itself until it gets 4 choices including the correct answer.

### `GenerateMultiChoiceProblems(cards []FlashCard) []MultiChoiceProblem`

This function takes a slice of `FlashCard`s, and returns  a slice of `MultiChoiceProblem`s generated from the definitions of the `FlashCards`. It maps over each `FlashCard` and generates a `MultiChoiceProblem` using a random definition as the `Question`, four randomly selected words including the correct word as multiple choices for the question, and the correct word itself as the `Answer`.
