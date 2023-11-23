package brain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type randomCommandSequenceGenerator struct {
}

func (g randomCommandSequenceGenerator) Generate(volume int) ([]int, error) {
	if volume < 0 {
		return nil, errors.New("volume is incorrect")
	}

	commands := make([]int, volume)
	for i := 0; i < volume; i++ {
		commands[i] = i
	}
	return commands, nil
}

func TestBrain_IncreaseAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		volume          int
		delta           int
		expectedAddress int
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name:            "empty brain with delta 0",
			volume:          0,
			delta:           0,
			expectedAddress: 0,
			wantErr:         assert.Error,
		},
		{
			name:            "empty brain with positive delta",
			volume:          0,
			delta:           1,
			expectedAddress: 0,
			wantErr:         assert.Error,
		},
		{
			name:            "empty brain with negative delta",
			volume:          0,
			delta:           -1,
			expectedAddress: 0,
			wantErr:         assert.Error,
		},
		{
			name:            "brain with positive delta",
			volume:          4,
			delta:           1,
			expectedAddress: 1,
		},
		{
			name:            "brain with negative delta",
			volume:          4,
			delta:           -1,
			expectedAddress: 3,
		},
		{
			name:            "brain with positive delta greater than volume",
			volume:          4,
			delta:           5,
			expectedAddress: 1,
		},
		{
			name:            "brain with negative delta greater than volume",
			volume:          4,
			delta:           -5,
			expectedAddress: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := &randomCommandSequenceGenerator{}
			brainBuilder, err := NewBuilder().WithVolume(tt.volume, generator)
			if (err != nil) == tt.wantErr(t, err) {
				return
			}
			brain := brainBuilder.Build()

			brain.IncreaseAddress(tt.delta)

			assert.EqualValuesf(t, tt.expectedAddress, brain.Address(), "address should be %d", tt.expectedAddress)
		})
	}
}

func TestBrain_Commands(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		volume   int
		expected int
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:    "empty brain",
			volume:  0,
			wantErr: assert.Error,
		},
		{
			name:     "brain with volume greater than 0",
			volume:   3,
			expected: 3,
			wantErr:  assert.NoError,
		},
		{
			name:    "brain with volume less than 0",
			volume:  -1,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := &randomCommandSequenceGenerator{}
			brainBuilder, err := NewBuilder().WithVolume(tt.volume, generator)
			if (err != nil) == tt.wantErr(t, err) {
				return
			}
			brain := brainBuilder.Build()

			assert.EqualValuesf(t, tt.expected, len(brain.Commands()), "commands should be %v", tt.expected)
		})
	}
}

func TestBrain_Command(t *testing.T) {
	t.Parallel()

	var (
		zeroIdentifier = 0

		defaultVolume               = 2
		identifierGreaterThanVolume = 3
		identifierLessThanZero      = -2
	)

	tests := []struct {
		name       string
		volume     int
		identifier *int
		expected   int
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "empty brain with identifier",
			volume:     0,
			identifier: &zeroIdentifier,
			wantErr:    assert.Error,
		},
		{
			name:       "empty brain without identifier",
			volume:     0,
			identifier: nil,
			wantErr:    assert.Error,
		},
		{
			name:       "brain with volume greater than 0 with identifier greater than 0",
			volume:     defaultVolume,
			identifier: &identifierGreaterThanVolume,
			expected:   1,
			wantErr:    assert.NoError,
		},
		{
			name:       "brain with volume greater than 0 with identifier less than 0",
			volume:     defaultVolume,
			identifier: &identifierLessThanZero,
			expected:   0,
			wantErr:    assert.NoError,
		},
		{
			name:     "brain with volume greater than 0 without identifier",
			volume:   defaultVolume,
			expected: 0,
			wantErr:  assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := &randomCommandSequenceGenerator{}
			brainBuilder, err := NewBuilder().WithVolume(tt.volume, generator)
			if (err != nil) == tt.wantErr(t, err) {
				return
			}

			brain := brainBuilder.Build()

			assert.EqualValuesf(t, tt.expected, brain.Command(tt.identifier), "command should be %v", tt.expected)
		})
	}
}
