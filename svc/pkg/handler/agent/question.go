package agent

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	"ynufes-mypage-backend/svc/pkg/registry"
	schema "ynufes-mypage-backend/svc/pkg/schema/agent"
	schemaQ "ynufes-mypage-backend/svc/pkg/schema/question"
	uc "ynufes-mypage-backend/svc/pkg/uc/question"
)

type Question struct {
	createUC uc.CreateUseCase
}

func NewQuestion(rgst registry.Registry) Question {
	return Question{
		createUC: uc.NewCreate(rgst),
	}
}

func (q Question) CreateHandler() gin.HandlerFunc {
	return util.Handler(func(c *gin.Context, tUser user.User) {
		var req schema.CreateQuestionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		secID, err := identity.ImportID(req.SectionID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var qIDAfter id.QuestionID
		if req.AfterID != "" {
			qIDAfter, err = identity.ImportID(req.AfterID)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}

		qType, err := question.NewType(req.Type)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var targetQ question.Question
		switch qType {
		case question.TypeCheckBox:
			targetQ, err = q.loadCheckboxQuestion(req)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		case question.TypeRadio:
			targetQ, err = q.loadRadioQuestion(req)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		//case question.TypeFile:
		default:
			c.JSON(400, gin.H{"error": errors.New("NOT IMPLEMENTED YET").Error()})
		}

		ipt := uc.CreateInput{
			Ctx:       c,
			UserID:    tUser.ID,
			SectionID: secID,
			After:     qIDAfter,
			Order:     req.PosAt,
			Question:  targetQ,
		}
		res, err := q.createUC.Do(ipt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var radio *schemaQ.RadioQuestionInfo
		var checkbox *schemaQ.CheckboxQuestionInfo
		switch qType {
		case question.TypeCheckBox:
			checkQ, err := question.ImportCheckBoxQuestion(res.Question.Export())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			opts := make([]schemaQ.CheckboxOptionInfo, len(checkQ.Options))
			for i := range checkQ.Options {
				opts[i] = schemaQ.CheckboxOptionInfo{
					ID:   checkQ.Options[i].ID.ExportID(),
					Text: checkQ.Options[i].Text,
				}
			}
			checkbox = &schemaQ.CheckboxQuestionInfo{
				Options: opts,
			}
		case question.TypeRadio:
			radioQ, err := question.ImportRadioButtonsQuestion(res.Question.Export())
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			opts := make([]schemaQ.RadioOptionInfo, len(radioQ.Options))
			for i := range radioQ.Options {
				opts[i] = schemaQ.RadioOptionInfo{
					ID:   radioQ.Options[i].ID.ExportID(),
					Text: radioQ.Options[i].Text,
				}
			}
			radio = &schemaQ.RadioQuestionInfo{
				Options: opts,
			}
		}
		resp := schema.CreateQuestionResponse{
			QuestionID: res.Question.GetID().ExportID(),
			Radio:      radio,
			Checkbox:   checkbox,
		}
		c.AbortWithStatusJSON(200, resp)
	}).GinHandler()
}

func (q Question) loadCheckboxQuestion(req schema.CreateQuestionRequest) (*question.CheckBoxQuestion, error) {
	if req.Checkbox == nil {
		return nil, errors.New("checkbox is nil")
	}

	fid, err := identity.ImportID(req.FormID)
	if err != nil {
		return nil, err
	}

	options := make([]question.CheckBoxOption, len(req.Checkbox.Options))
	for i := range options {
		options[i] = question.CheckBoxOption{
			ID:   identity.IssueID(),
			Text: req.Checkbox.Options[i],
		}
	}
	optionsOrder := make(map[question.CheckBoxOptionID]float64, len(req.Checkbox.Options))
	for i := range req.Checkbox.Options {
		// default order value will be 0.0, 1.0, 2.0, 3.0, 4.0...
		optionsOrder[options[i].ID] = float64(i)
	}
	newQ := question.NewCheckBoxQuestion(
		nil, req.Text, options, optionsOrder, fid,
	)
	return newQ, nil
}

func (q Question) loadRadioQuestion(req schema.CreateQuestionRequest) (*question.RadioButtonsQuestion, error) {
	if req.Radio == nil {
		return nil, errors.New("radio button is nil")
	}

	fid, err := identity.ImportID(req.FormID)
	if err != nil {
		return nil, err
	}

	options := make([]question.RadioButtonOption, len(req.Radio.Options))
	for i := range options {
		options[i] = question.RadioButtonOption{
			ID:   identity.IssueID(),
			Text: req.Radio.Options[i],
		}
	}
	optionsOrder := make(map[question.RadioButtonOptionID]float64, len(req.Radio.Options))
	for i := range req.Radio.Options {
		// default order value will be 0.0, 1.0, 2.0, 3.0, 4.0...
		optionsOrder[options[i].ID] = float64(i)
	}
	newQ := question.NewRadioButtonsQuestion(
		nil, req.Text, options, optionsOrder, fid,
	)
	return newQ, nil
}
