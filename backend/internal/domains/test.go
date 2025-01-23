package domains

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type Test struct {
	ID        uuid.UUID
	ChapterID uuid.UUID
	Input     string
	Output    string
}

func NewTest(test *repo.TTest) *Test {
	return &Test{
		ID:        test.ID,
		ChapterID: test.ChapterID,
		Input:     test.InputValue,
		Output:    test.OutputValue,
	}
}

func NewTests(tests []repo.TTest) []Test {
	var domainTests []Test
	for _, test := range tests {
		domainTests = append(domainTests, *NewTest(&test))
	}

	return domainTests
}
