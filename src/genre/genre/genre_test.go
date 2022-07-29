package genre

import (
	"context"
	"github.com/go-kit/kit/log"
	"reflect"
	"testing"
)

func Test_repository_GetGenre(t *testing.T) {
	type fields struct {
		db     []Genre
		logger log.Logger
	}
	type args struct {
		ctx  context.Context
		slug string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantGenre Genre
		wantErr   bool
	}{
		{
			name: "slug incorrect",
			args: args{slug: "alt-"},
			wantGenre: Genre{
				BpmRange:         "70-160",
				Description:      "A more experimental take on rap, often characterised by genre-blending, unique production and alternative arrangements.",
				Id:               "25f94646-d0b7-41fe-89fe-ec62c0463e5a",
				IsParent:         false,
				KeyInstruments:   "",
				KeyLocation:      "",
				ParentId:         "55bce497-7c9a-47c6-a600-c0cdc5db687e",
				ReleaseDate:      "2020-02-04T20:28:33.693Z",
				ShortDescription: "Alternative, experimental take on rap",
				Slug:             "alt-rap",
				Title:            "Alt. Rap",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			gotGenre, err := r.GetGenre(tt.args.ctx, tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGenre() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGenre, tt.wantGenre) {
				t.Errorf("GetGenre() gotGenre = %v, want %v", gotGenre, tt.wantGenre)
			}
		})
	}
}
