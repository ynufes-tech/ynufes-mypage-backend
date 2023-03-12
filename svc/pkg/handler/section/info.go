package section

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/handler/util"
	schemaS "ynufes-mypage-backend/svc/pkg/schema/section"
	"ynufes-mypage-backend/svc/pkg/uc/section"
)

func (h Section) InfoHandler() gin.HandlerFunc {
	return util.Handler(func(c *gin.Context, user user.User) {
		sid, has := c.GetQuery("section_id")
		if !has {
			c.JSON(400, gin.H{"error": "section_id is required"})
			return
		}
		sectionID, err := identity.ImportID(sid)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		opt, err := h.infoUC.Do(
			section.InfoInput{
				Ctx:       c,
				UserID:    user.ID,
				SectionID: sectionID,
			})
		if err != nil {
			if err == exception.ErrUnauthorized {
				c.JSON(401, gin.H{"error": err.Error()})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		qids := opt.Section.QuestionIDs.GetOrderedIDs()
		respQs := make([]schemaS.Question, 0, len(qids))
		for i := range qids {
			target, ok := opt.Section.Questions[qids[i]]
			if !ok {
				fmt.Println("internal error: question not found")
				fmt.Println("section id: ", opt.Section.ID.ExportID())
				fmt.Println("question id: ", qids[i].ExportID())
				for _, qid := range qids {
					fmt.Println("qid: ", qid.ExportID())
				}
				for k := range opt.Section.Questions {
					fmt.Println("key: ", k.ExportID())
				}
				c.JSON(500, gin.H{"error": "server process error: question not found"})
				return
			}
			respQ := schemaS.Question{
				ID:   target.GetID().ExportID(),
				Type: target.GetType().String(),
				Text: target.GetText(),
			}
			switch target.GetType() {
			case question.TypeRadio:
				radioQ, err := question.ImportRadioButtonsQuestion(target.Export())
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}
				options := make([]schemaS.Option, 0, len(radioQ.Options))
				for i := range radioQ.Options {
					options = append(options, schemaS.Option{
						ID:   radioQ.Options[i].ID.ExportID(),
						Text: radioQ.Options[i].Text,
					})
				}
				respQ.Options = &options
			case question.TypeCheckBox:
				checkQ, err := question.ImportCheckBoxQuestion(target.Export())
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}
				options := make([]schemaS.Option, 0, len(checkQ.Options))
				for i := range checkQ.Options {
					options = append(options, schemaS.Option{
						ID:   checkQ.Options[i].ID.ExportID(),
						Text: checkQ.Options[i].Text,
					})
				}
				respQ.Options = &options
			case question.TypeFile:
				fileQ, err := question.ImportFileQuestion(target.Export())
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}
				exts := make([]string, len(fileQ.Constraint.GetExtensions()))
				for i := range fileQ.Constraint.GetExtensions() {
					exts[i] = string(fileQ.Constraint.GetExtensions()[i])
				}
				fConstraint := schemaS.FileConstraint{
					FileType:   fileQ.Constraint.GetFileType().String(),
					Extensions: exts,
				}
				respQ.FileConstraint = &fConstraint
			}
			respQs = append(respQs, respQ)
		}
		resp := schemaS.SectionInfoResponse{
			ID:        opt.Section.ID.ExportID(),
			Questions: respQs,
		}
		c.JSON(200, resp)
	}).GinHandler()
}
