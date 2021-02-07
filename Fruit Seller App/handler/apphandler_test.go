package handler

import (
	"reflect"
	"testing"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/domain"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"
)

func TestNewHandler(t *testing.T) {
	cases := map[string]struct {
		want       *Handler
		appContext *appcontext.AppContext
		domain     domain.AppDomain
		validation validation.AppValidation
		util       utils.AppUtil
	}{
		"when everything is passed as nil": {
			want:       NewHandler(nil, nil, nil, nil),
			appContext: nil,
			domain:     nil,
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			h := NewHandler(v.appContext, v.domain, v.validation, v.util)
			if !reflect.DeepEqual(v.want, h) {
				t.Errorf("handler mismatched\nwant: %v\ngot:%v\n", v.want, h)
			}
		})
	}

}
