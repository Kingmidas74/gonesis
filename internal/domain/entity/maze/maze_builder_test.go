package maze

import "testing"

func TestGenerateGrid_FirstCell(t *testing.T) {
	width, height := 5, 5
	type fields struct {
		FirstFilled bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "first filled",
			fields: fields{
				FirstFilled: true,
			},
			wantErr: false,
		},
		{
			name: "first unfilled",
			fields: fields{
				FirstFilled: false,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze, err := NewMazeBuilder().generateGrid(width, height, tt.fields.FirstFilled)
			if (err != nil) != tt.wantErr {
				t.Errorf("given error %v, wantErr %v", err, tt.wantErr)
			}
			if maze[0] != tt.fields.FirstFilled {
				t.Errorf("first cell expected %t, but got %t", tt.fields.FirstFilled, maze[0])
			}
		})
	}
}

func TestGenerateGrid_Size(t *testing.T) {
	firstFilled := true

	type fields struct {
		Width  int
		Height int
		Size   int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal size",
			fields: fields{
				Width:  5,
				Height: 5,
				Size:   25,
			},
			wantErr: false,
		},
		{
			name: "zero size",
			fields: fields{
				Width:  0,
				Height: 0,
			},
			wantErr: true,
		},
		{
			name: "negative width",
			fields: fields{
				Width:  -5,
				Height: 5,
			},
			wantErr: true,
		},
		{
			name: "negative height",
			fields: fields{
				Width:  5,
				Height: -5,
			},
			wantErr: true,
		},
		{
			name: "negative width and height",
			fields: fields{
				Width:  -5,
				Height: -5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze, err := NewMazeBuilder().generateGrid(tt.fields.Width, tt.fields.Height, firstFilled)
			if (err != nil) != tt.wantErr {
				t.Errorf("given error %v, wantErr %v", err, tt.wantErr)
			}
			if len(maze) != tt.fields.Size {
				t.Errorf("maze size expected %d, but got %d", tt.fields.Size, len(maze))
			}
		})
	}
}
