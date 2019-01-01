package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
    "errors"
)

type Sport struct {
    *revel.Controller
}

func (c Sport) Index() revel.Result {
    sports := []models.Sport{}

    result := DB.Order("id desc").Find(&sports);
    err := result.Error
    if err != nil {
        return c.RenderError(errors.New("Record Not Found"))
    }
    return c.Render(sports)
}
func (c Sport) Create() revel.Result {
    sport := models.Sport{
        Sport: c.Params.Form.Get("sport"),
    }
    ret := DB.Create(&sport)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Create failure." + ret.Error.Error()))
    }
    return c.Redirect("/sports")    
}
func (c Sport) Delete() revel.Result {
    id := c.Params.Route.Get("id")
    sports := []models.Sport{}
    ret := DB.Delete(&sports, id)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Delete failure." + ret.Error.Error()))
    }
    return c.Redirect("/sports")    
}

func (c Sport) RedirectToSports() revel.Result {
    return c.Redirect("/sports")    
}