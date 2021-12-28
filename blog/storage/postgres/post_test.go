package postgres

import (
	"context"
	"myGo/Blogs/blog/storage"
	"sort"

	// "log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Create(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Post
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Post{
				Title:       "This is title",
				Description: "This is description",
				CategoryId:  1,
			},
			want: 1,
		},
		{
			name: "CREATE_BLOG_SUCCESS 2",
			in: storage.Post{
				Title:       "This is titl2",
				Description: "This is descriptio2",
				CategoryId:  1,
			},
			want: 2,
		},
		{
			name: "IF_NOT_UNIQUE",
			in: storage.Post{
				Title:       "This is title",
				Description: "This is description",
				CategoryId:  1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Post
		want    int64
		wantErr bool
	}{
		{
			name: "UPDATE_BLOG_SUCCESS",
			in: storage.Post{
				ID:          1,
				Title:       "This is Updated",
				Description: "This is Updted",
				CategoryId:  1,

			},
			want: 1,
		},
		{
			name: "UPDATE_BLOG_SUCCESS",
			in: storage.Post{
				ID:          2,
				Title:       "This is Updated 2",
				Description: "This is Updted 2",
				CategoryId:  1,
			},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Update(context.Background(), tt.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListPost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "GET_LIST_POST_SUCCESS",
			want: []storage.Post{
				{
					ID:          1,
				Title:       "This is Updated",
				Description: "This is Updted",
				CategoryName: "Category",

				CategoryId:  1,
				},
				{	ID:          2,
					Title:       "This is Updated 2",
					Description: "This is Updted 2",
				CategoryName: "Category",

					CategoryId:  1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			gotList, err := s.List(context.Background())

			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantList := tt.want

			sort.Slice(wantList, func(i, j int) bool {
				return wantList[i].ID < wantList[j].ID
			})

			sort.Slice(gotList, func(i, j int) bool {
				return gotList[i].ID < gotList[j].ID
			})

			for i, got := range gotList {

				if !cmp.Equal(got, wantList[i]) {
					t.Errorf("Diff: got -, want += %v", cmp.Diff(got, wantList[i]))
				}

			}
		})
	}
}

// list
func Test_Delete(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name:    "DELETE_BLOG_SUCCESS",
			in:      23,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(context.Background(), tt.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
