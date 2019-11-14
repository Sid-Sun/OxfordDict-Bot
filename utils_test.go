package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetFormattedMessage(t *testing.T) {
	testCases := []struct {
		name           string
		exec           func() (string, bool)
		expectedResult bool
		expectedOutput string
	}{
		{
			name: "Successfully find the definition and form the initial message for turn",
			exec: func() (string, bool) {
				query := "turn"
				index := 0
				response, err := getDefinition(query)
				require.NoError(t, err)
				return getFormattedMessage(response, index, query)
			},
			expectedOutput: `turn by Oxford University Press

Language: 
EN-GB

Definition: 
move in a circular direction wholly or partly round an axis or point

Lexical Category: 
Verb

Examples: 
the big wheel was turning`,
			expectedResult: true,
		},
		{
			name: "Successfully find the definition and form the fourth message for turn",
			exec: func() (string, bool) {
				query := "turn"
				index := 3
				response, err := getDefinition(query)
				require.NoError(t, err)
				return getFormattedMessage(response, index, query)
			},
			expectedOutput: `turn by Oxford University Press

Language: 
EN-GB

Definition: 
start doing or becoming involved with

Lexical Category: 
Verb

Examples: 
in 1939 he turned to films in earnest`,
			expectedResult: true,
		},
		{
			name: "Fail to find the initial definition for vps",
			exec: func() (string, bool) {
				query := "VPS"
				index := 0
				response, err := getDefinition(query)
				require.NoError(t, err)
				return getFormattedMessage(response, index, query)
			},
			expectedOutput: ``,
			expectedResult: false,
		},
		{
			name: "Fail to find the sixth definition for vps",
			exec: func() (string, bool) {
				query := "VPS"
				index := 5
				response, err := getDefinition(query)
				require.NoError(t, err)
				return getFormattedMessage(response, index, query)
			},
			expectedOutput: ``,
			expectedResult: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output, success := testCase.exec()
			assert.Equal(t, testCase.expectedResult, success)
			assert.Equal(t, testCase.expectedOutput, output)
		})
	}
}

