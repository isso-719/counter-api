package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	test_helper "github.com/isso-719/counter-api/lib/test"
	"github.com/isso-719/counter-api/src/domain"
	"github.com/isso-719/counter-api/src/repository"
	"reflect"
	"testing"
)

func Test_Increment(t *testing.T) {
	type Fields struct {
		repo repository.IFCounterRepository
	}
	type Args struct {
		ctx context.Context
		url string
	}
	type Returns struct {
		want *domain.Counter
		err  error
	}
	type testContext struct {
		fields  Fields
		args    Args
		returns Returns
	}

	tests := []struct {
		name        string
		testContext func(ctrl *gomock.Controller) *testContext
	}{
		{
			name: "正常: 何もないURLを指定した場合",
			testContext: func(ctrl *gomock.Controller) *testContext {
				counterRepository := repository.NewMockIFCounterRepository(ctrl)
				counterRepository.EXPECT().BeginTx(gomock.Any()).Return(nil)
				counterRepository.EXPECT().CommitTx().Return(nil)
				counterRepository.EXPECT().TxRead("http://example.com").Return(nil, nil)
				counterRepository.EXPECT().TxWrite("http://example.com", &domain.Counter{URL: "http://example.com", Count: 1}).Return(&domain.Counter{URL: "http://example.com", Count: 1}, nil)

				return &testContext{
					fields: Fields{
						repo: counterRepository,
					},
					args: Args{
						ctx: context.Background(),
						url: "http://example.com",
					},
					returns: Returns{
						want: &domain.Counter{URL: "http://example.com", Count: 1},
						err:  nil,
					},
				}
			},
		},
		{
			name: "正常: すでに登録したURLを指定した場合",
			testContext: func(ctrl *gomock.Controller) *testContext {
				counterRepository := repository.NewMockIFCounterRepository(ctrl)
				counterRepository.EXPECT().BeginTx(gomock.Any()).Return(nil)
				counterRepository.EXPECT().CommitTx().Return(nil)
				counterRepository.EXPECT().TxRead("http://example.com/abc/d-e/f_g").Return(&domain.Counter{URL: "http://example.com/abc/d-e/f_g", Count: 5}, nil)
				counterRepository.EXPECT().TxWrite("http://example.com/abc/d-e/f_g", &domain.Counter{URL: "http://example.com/abc/d-e/f_g", Count: 6}).Return(&domain.Counter{URL: "http://example.com/abc/d-e/f_g", Count: 6}, nil)

				return &testContext{
					fields: Fields{
						repo: counterRepository,
					},
					args: Args{
						ctx: context.Background(),
						url: "http://example.com/abc/d-e/f_g",
					},
					returns: Returns{
						want: &domain.Counter{URL: "http://example.com/abc/d-e/f_g", Count: 6},
						err:  nil,
					},
				}
			},
		},
		{
			name: "異常: URLが空文字の場合",
			testContext: func(ctrl *gomock.Controller) *testContext {
				counterRepository := repository.NewMockIFCounterRepository(ctrl)

				return &testContext{
					fields: Fields{
						repo: counterRepository,
					},
					args: Args{
						ctx: context.Background(),
						url: "",
					},
					returns: Returns{
						want: nil,
						err:  errors.New("url is required"),
					},
				}
			},
		},
		{
			name: "異常: URLでない場合",
			testContext: func(ctrl *gomock.Controller) *testContext {
				counterRepository := repository.NewMockIFCounterRepository(ctrl)

				return &testContext{
					fields: Fields{
						repo: counterRepository,
					},
					args: Args{
						ctx: context.Background(),
						url: "test",
					},
					returns: Returns{
						want: nil,
						err:  errors.New("url is invalid"),
					},
				}
			},
		},
		{
			name: "異常: URLがhttp://で始まらない場合",
			testContext: func(ctrl *gomock.Controller) *testContext {
				counterRepository := repository.NewMockIFCounterRepository(ctrl)

				return &testContext{
					fields: Fields{
						repo: counterRepository,
					},
					args: Args{
						ctx: context.Background(),
						url: "example.com",
					},
					returns: Returns{
						want: nil,
						err:  errors.New("url is invalid"),
					},
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test_helper.RunTestWithGoMock(t, func(ctrl *gomock.Controller) {
				tc := tt.testContext(ctrl)
				u := &counterService{
					counterRepository: tc.fields.repo,
				}
				got, err := u.Increment(tc.args.ctx, tc.args.url)
				if !reflect.DeepEqual(got, tc.returns.want) {
					t.Errorf("got: %v, want: %v", got, tc.returns.want)
				}
				if !reflect.DeepEqual(err, tc.returns.err) {
					t.Errorf("got: %v, want: %v", err, tc.returns.err)
				}
			})
		})
	}
}
