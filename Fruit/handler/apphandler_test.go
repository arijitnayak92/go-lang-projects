package handler

import (
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/domain"
	"gitlab.com/affordmed/affmed/util"
	"gitlab.com/affordmed/affmed/validation"
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	cases := map[string]struct {
		want       *Handler
		appContext *appcontext.AppContext
		domain     domain.AppDomain
		validation validation.AppValidation
		util       util.AppUtil
	}{
		"when everything is passed as nil": {
			want:       NewHandler(nil, nil, nil, nil),
			appContext: nil,
			domain:     nil,
			validation: nil,
			util:       nil,
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
