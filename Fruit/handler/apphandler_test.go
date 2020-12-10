package handler

import (
	"reflect"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
)

func TestNewHandler(t *testing.T) {
	cases := map[string]struct {
		want       *Handler
		appContext *appcontext.AppContext
		domain     domain.AppDomain
	}{
		"when everything is passed as nil": {
			want:       NewHandler(nil, nil),
			appContext: nil,
			domain:     nil,
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			h := NewHandler(v.appContext, v.domain)
			if !reflect.DeepEqual(v.want, h) {
				t.Errorf("handler mismatched\nwant: %v\ngot:%v\n", v.want, h)
			}
		})
	}

}
