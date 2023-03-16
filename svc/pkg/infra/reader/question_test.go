package reader

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestQuestion_GetByID(t *testing.T) {
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	questionW := writer.NewQuestion(fbt.GetClient())
	op1, op2, op3 := identity.IssueID(), identity.IssueID(), identity.IssueID()
	options := map[question.CheckBoxOptionID]question.CheckBoxOption{
		op1: {Text: "testChoice1", ID: op1},
		op2: {Text: "testChoice2", ID: op2},
		op3: {Text: "testChoice3", ID: op3},
	}
	order := map[question.CheckBoxOptionID]float64{
		op2: 1,
		op1: 2,
		op3: 3,
	}
	var question1 question.Question
	question1 = question.NewCheckBoxQuestion(nil, "testQuestion1",
		options, order, identity.IssueID())
	assert.NoError(t, questionW.Create(context.Background(), &question1))
	tests := []struct {
		name    string
		query   id.QuestionID
		want    question.Question
		wantErr error
	}{
		{
			name:  "Success",
			query: question1.GetID(),
			want:  question1,
		},
		{
			name:    "NotFound",
			query:   identity.IssueID(),
			want:    nil,
			wantErr: exception.ErrNotFound,
		},
	}
	r := NewQuestion(fbt.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := r.GetByID(ctx, tt.query)
			if !assert.ErrorIs(t, tt.wantErr, err) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == nil {
				assert.Equal(t, tt.want.Export(), (*got).Export())
				assert.Equal(t, tt.want.GetID(), (*got).GetID())
				assert.Equal(t, tt.want.GetType(), (*got).GetType())
				assert.Equal(t, tt.want.GetText(), (*got).GetText())
				assert.Equal(t, tt.want.GetFormID(), (*got).GetFormID())
			}
		})
	}
}

func TestQuestion_ListByEventID(t *testing.T) {
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	events := []event.Event{
		{
			Name: "EventExists",
		},
		{
			Name: "EventWOQuestions",
		},
		{
			Name: "EventNotExists",
		},
	}
	eventW := writer.NewEvent(fbt.GetClient())
	assert.NoError(t, eventW.Create(context.Background(), &events[0]))
	assert.NoError(t, eventW.Create(context.Background(), &events[1]))

	qs := make([]question.Question, 2)
	cop1, cop2, cop3 := identity.IssueID(), identity.IssueID(), identity.IssueID()
	copt := map[question.CheckBoxOptionID]question.CheckBoxOption{
		cop1: {Text: "checkChoice1", ID: cop1},
		cop2: {Text: "checkChoice2", ID: cop2},
		cop3: {Text: "checkChoice3", ID: cop3},
	}
	rop1, rop2, rop3 := identity.IssueID(), identity.IssueID(), identity.IssueID()
	ropt := map[question.RadioButtonOptionID]question.RadioButtonOption{
		rop1: {Text: "radioChoice1", ID: rop1},
		rop2: {Text: "radioChoice2", ID: rop2},
		rop3: {Text: "radioChoice3", ID: rop3},
	}
	cOrder := map[question.CheckBoxOptionID]float64{
		cop3: 1,
		cop2: 2,
		cop1: 3,
	}
	rOrder := map[question.RadioButtonOptionID]float64{
		rop2: 1,
		rop1: 2,
		rop3: 3,
	}
	qs[0] = question.NewCheckBoxQuestion(nil, "testQuestion1",
		copt, cOrder, identity.IssueID())
	qs[1] = question.NewRadioButtonsQuestion(nil, "testQuestion2",
		ropt, rOrder, identity.IssueID())
	questionW := writer.NewQuestion(fbt.GetClient())
	assert.NoError(t, questionW.Create(context.Background(), &qs[0]))
	assert.NoError(t, questionW.Create(context.Background(), &qs[1]))

	tests := []struct {
		name     string
		query    id.EventID
		want     []question.Question
		hasError bool
	}{
		{
			name:     "Success",
			query:    events[0].ID,
			want:     qs,
			hasError: false,
		}, {
			name:     "NoHits",
			query:    identity.IssueID(),
			want:     []question.Question{},
			hasError: false,
		}, {
			name:     "EventNotExists",
			query:    events[1].ID,
			want:     []question.Question{},
			hasError: false,
		}, {
			name:     "InvalidQuery",
			query:    nil,
			want:     nil,
			hasError: true,
		},
	}
	r := NewQuestion(fbt.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs, err := r.ListByEventID(context.Background(), tt.query)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.want != nil {
					eq := testutil.CheckWithoutOrder(tt.want, qs,
						func(a, b question.Question) bool {
							var equal bool
							if a.GetType() != b.GetType() {
								fmt.Printf("type not equal: %v, %v\n", a.GetType(), b.GetType())
								equal = false
							}
							if a.GetText() != b.GetText() {
								fmt.Printf("text not equal: %v, %v\n", a.GetText(), b.GetText())
								equal = false
							}
							if a.GetFormID() != b.GetFormID() {
								fmt.Printf("formID not equal: %v, %v\n", a.GetFormID(), b.GetFormID())
								equal = false
							}
							if a.GetID() != b.GetID() {
								fmt.Printf("id not equal: %v, %v\n", a.GetID(), b.GetID())
								equal = false
							}
							if reflect.DeepEqual(a.Export().Customs, b.Export().Customs) {
								fmt.Printf("customs not equal: %v, %v\n", a.Export().Customs, b.Export().Customs)
								equal = false
							}
							return equal
						})
					if !eq {
						fmt.Printf("want: %#v, got: %#v\n", tt.want, qs)
					}
				} else {
					assert.Nil(t, qs)
				}
			}
		})
	}
}
