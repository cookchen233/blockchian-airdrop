package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"strings"
	"blockchain-deal-hunter/api/utility"
	"errors"
	"gopkg.in/go-playground/validator.v9"
)

type Controller struct {

}

func (ctrl *Controller) DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return utility.ParamError{err}
	}
	//获取验证器
	valid, err := ctrl.GetValidator(c)
	if err != nil {
		return err
	}
	//获取翻译器
	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return utility.ParamError{errors.New(strings.Join(sliceErrs, ","))}
	}
	return nil
}

func (ctrl *Controller) GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get("ValidatorKey")
	if !ok {
		return nil, errors.New("未设置验证器")
	}
	validator, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validator, nil
}

func GetTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get("TranslatorKey")
	if !ok {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}