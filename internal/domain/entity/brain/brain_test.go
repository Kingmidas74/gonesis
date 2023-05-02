package brain

import "testing"

func TestNew_CommandsLength(t *testing.T) {
	commands := make([]int, 64)
	b := New(commands)
	if len(b.commands) != 64 {
		t.Error("length of commands")
	}
}

func TestBrain_Mod(t *testing.T) {
	commands := make([]int, 64)
	b := New(commands)
	if b.currentAddress != 0 {
		t.Error("initial address is wrong")
	}
}

func TestMod_Valid(t *testing.T) {
	commands := make([]int, 3)
	b := New(commands)

	type fields struct {
		rowAddress      int
		expectedAddress int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "address 0",
			fields: fields{
				rowAddress:      0,
				expectedAddress: 0,
			},
		},
		{
			name: "address 1",
			fields: fields{
				rowAddress:      1,
				expectedAddress: 1,
			},
		},
		{
			name: "address 2",
			fields: fields{
				rowAddress:      2,
				expectedAddress: 2,
			},
		},
		{
			name: "address 3",
			fields: fields{
				rowAddress:      3,
				expectedAddress: 0,
			},
		},
		{
			name: "address 4",
			fields: fields{
				rowAddress:      4,
				expectedAddress: 1,
			},
		},
		{
			name: "address 5",
			fields: fields{
				rowAddress:      5,
				expectedAddress: 2,
			},
		},
		{
			name: "address 6",
			fields: fields{
				rowAddress:      6,
				expectedAddress: 0,
			},
		},
		{
			name: "address -11",
			fields: fields{
				rowAddress:      -1,
				expectedAddress: 2,
			},
		},
		{
			name: "address -2",
			fields: fields{
				rowAddress:      -2,
				expectedAddress: 1,
			},
		},
		{
			name: "address -3",
			fields: fields{
				rowAddress:      -3,
				expectedAddress: 0,
			},
		},
		{
			name: "address -4",
			fields: fields{
				rowAddress:      -4,
				expectedAddress: 2,
			},
		},
		{
			name: "address -5",
			fields: fields{
				rowAddress:      -5,
				expectedAddress: 1,
			},
		},
		{
			name: "address -6",
			fields: fields{
				rowAddress:      -6,
				expectedAddress: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.mod(tt.fields.rowAddress); tt.fields.expectedAddress != got {
				t.Errorf("exp addres %d, got address %d", tt.fields.expectedAddress, got)
			}
		})
	}
}

func TestBrain_MoveAddressOn(t *testing.T) {
	commands := make([]int, 3)
	b := New(commands)

	type fields struct {
		currentAddress  int
		delta           int
		expectedAddress int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "address 0",
			fields: fields{
				currentAddress:  0,
				delta:           1,
				expectedAddress: 1,
			},
		},
		{
			name: "address 1",
			fields: fields{
				currentAddress:  0,
				delta:           -1,
				expectedAddress: 2,
			},
		},
		{
			name: "address 2",
			fields: fields{
				currentAddress:  0,
				delta:           0,
				expectedAddress: 0,
			},
		},
		{
			name: "address 3",
			fields: fields{
				currentAddress:  2,
				delta:           2,
				expectedAddress: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.SetCurrentAddress(tt.fields.currentAddress)
			b.MoveAddressOn(tt.fields.delta)
			if got := b.currentAddress; tt.fields.expectedAddress != got {
				t.Errorf("exp addres %d, got address %d", tt.fields.expectedAddress, got)
			}
		})
	}
}
