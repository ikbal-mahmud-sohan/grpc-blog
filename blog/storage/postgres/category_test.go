package postgres

import (
	"myGo/Blogs/blog/storage"
	"context"
	"log"
	"testing"
)

func TestCreateCat(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Category
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				Title: "Category",
			},
			want: 1,
		},
		{
			name: "IF_NOT_UNIQUE",
			in: storage.Category{
				Title:       "Category",
			},
			wantErr:true,
		},
		{
			name: "IF_EMPTY",
			in: storage.Category{
				Title: "",
			},
			want:3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCat(context.Background(), tt.in)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateCat() = %v, want %v", got, tt.want)
			}
		})
	}
}
