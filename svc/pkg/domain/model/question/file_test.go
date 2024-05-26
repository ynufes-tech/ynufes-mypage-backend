package question

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
)

func TestImportFileQuestion(t *testing.T) {
	f1, f2, f3 := float32(0.5), float32(1.5), float32(3.0)
	v1, v2, v3 := 100, 200, 300
	SampleID1, sampleID2 := identity.IssueID(), identity.IssueID()
	tests := []struct {
		name string
		from StandardQuestion
		want *FileQuestion
	}{
		{
			name: "ImageQuestion - Simple",
			from: StandardQuestion{
				ID:     SampleID1,
				Text:   "Image Question",
				FormID: sampleID2,
				Type:   TypeFile,
				Customs: map[string]interface{}{
					FileQuestionFileTypeField: []bool{false, true, false},
					FileImageConstraintField: map[string]interface{}{
						FileImageConstraintRatio:  map[string]interface{}{},
						FileImageConstraintWidth:  map[string]interface{}{},
						FileImageConstraintHeight: map[string]interface{}{},
					},
				},
			},
			want: NewFileQuestion(SampleID1,
				"Image Question",
				FileTypes{
					AcceptAny:   false,
					AcceptImage: true,
					AcceptPDF:   false,
				}, ImageFileConstraint{}, sampleID2),
		},
		{
			name: "ImageQuestion - With Constraints",
			from: StandardQuestion{
				ID:     SampleID1,
				Text:   "Image Question",
				FormID: sampleID2,
				Type:   TypeFile,
				Customs: map[string]interface{}{
					FileQuestionFileTypeField: []bool{false, true, false},
					FileImageConstraintField: map[string]interface{}{
						FileImageConstraintRatio: map[string]interface{}{
							FileImageConstraintRatioEq: f1,
						},
						FileImageConstraintWidth: map[string]interface{}{
							FileImageConstraintDimensionMin: v1,
							FileImageConstraintDimensionMax: v2,
						},
						FileImageConstraintHeight: map[string]interface{}{
							FileImageConstraintDimensionMin: v2,
							FileImageConstraintDimensionMax: v3,
						},
					},
				},
			},
			want: NewFileQuestion(SampleID1,
				"Image Question",
				FileTypes{
					AcceptAny:   false,
					AcceptImage: true,
					AcceptPDF:   false,
				}, ImageFileConstraint{
					Ratio:     NewRatioSpec(nil, nil, &f1),
					MinNumber: nil,
					MaxNumber: nil,
					Width:     NewDimensionSpec(&v1, &v2, nil),
					Height:    NewDimensionSpec(&v2, &v3, nil),
				}, sampleID2),
		},
		{
			name: "ImageQuestion - With Constraints - MinMax",
			from: StandardQuestion{
				ID:     SampleID1,
				Text:   "Image Question",
				FormID: sampleID2,
				Type:   TypeFile,
				Customs: map[string]interface{}{
					FileQuestionFileTypeField: []bool{false, true, false},
					FileImageConstraintField: map[string]interface{}{
						FileImageConstraintRatio: map[string]interface{}{
							FileImageConstraintRatioMin: f2,
							FileImageConstraintRatioMax: f3,
						},
						FileImageConstraintWidth: map[string]interface{}{
							FileImageConstraintRatioEq: v1,
						},
						FileImageConstraintHeight: map[string]interface{}{
							FileImageConstraintRatioEq: v2,
						},
					},
				},
			},
			want: NewFileQuestion(SampleID1,
				"Image Question",
				FileTypes{
					AcceptAny:   false,
					AcceptImage: true,
					AcceptPDF:   false,
				}, ImageFileConstraint{
					Ratio:     NewRatioSpec(&f2, &f3, nil),
					MinNumber: nil,
					MaxNumber: nil,
					Width:     NewDimensionSpec(nil, nil, &v1),
					Height:    NewDimensionSpec(nil, nil, &v2),
				}, sampleID2),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ImportFileQuestion(tc.from)
			assert.NoError(t, err)
			assert.Equal(t, tc.want.ID, got.ID)
			assert.Equal(t, tc.want.Text, got.Text)
			assert.Equal(t, tc.want.MinNumber, got.MinNumber)
			assert.Equal(t, tc.want.MaxNumber, got.MaxNumber)
			assert.Equal(t, tc.want.Width, got.Width)
			assert.Equal(t, tc.want.Height, got.Height)
			assert.Equal(t, tc.want.Ratio, got.Ratio)
			export, err := got.Export()
			if assert.NoError(t, err) {
				assert.Equal(t, tc.from, *export)
			}
		})
	}
}
