package maze

import (
	"testing"

	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
)

func TestGenerateGrid_Size(t *testing.T) {
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
			maze, err := NewMazeBuilder[generator.GridGenerator]().SetWidth(tt.fields.Width).SetHeight(tt.fields.Height).Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("given error %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && len(maze.Content) != tt.fields.Size {
				t.Errorf("maze size expected %d, but got %d", tt.fields.Size, len(maze.Content))
			}
		})
	}
}
