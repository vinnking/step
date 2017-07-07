// 文章
package controllers

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"step/models"
)

type PostController struct {
	beego.Controller
}

// 检查用户是否登陆
func (p *PostController) Prepare() {
	_, user := models.IsLogin(p.Ctx)
	p.Data["nickname"] = user.Nickname
}

// Index 文章列表
func (p *PostController) Index() {
	p.Data["posts"] = models.PostList()
	p.Data["types"] = models.Menus()
	p.Data["status"] = models.PostStatus()
	p.Layout = "base.html"
	p.TplName = "post/index.html"
}

// Create 添加文章
func (p *PostController) Create() {
	if p.Ctx.Request.Method == "POST" {
		typeId, err := strconv.Atoi(strings.TrimSpace(p.Input().Get("type")))
		if err != nil {
			p.Redirect("/post/create", 302)
		}
		title := p.Input().Get("title")
		url := p.Input().Get("url")
		imageUrl := p.Input().Get("imageUrl")
		summary := p.Input().Get("summary")
		remark := p.Input().Get("remark")
		if typeId <= 0 {
			p.Redirect("/post/create", 302)
		}
		if strings.TrimSpace(title) == "" {
			p.Redirect("/post/create", 302)
		}
		if strings.TrimSpace(url) == "" {
			p.Redirect("/post/create", 302)
		}
		if strings.TrimSpace(imageUrl) == "" {
			p.Redirect("/post/create", 302)
		}
		if strings.TrimSpace(summary) == "" {
			p.Redirect("/post/create", 302)
		}
		if strings.TrimSpace(remark) == "" {
			p.Redirect("/post/create", 302)
		}
		var id int64
		if id, err = models.PostSave(&models.Post{
			Title:    title,
			Type:     typeId,
			Url:      url,
			ImageUrl: imageUrl,
			Summary:  summary,
			Remark:   remark,
			UserId:   1,
		}); err != nil {
			p.Redirect("/post/create", 302)
		}
		// 如果存在标签, 则保存标签
		labels := p.GetStrings("labels")
		if len(labels) > 0 {
			for _, labelId := range labels {
				var newLabelId int
				newLabelId, err = strconv.Atoi(labelId)
				if err != nil {
					p.Redirect("/post/create", 302)
				}
				if _, err := models.PostLabelSave(&models.PostLabel{
					LabelId: int64(newLabelId),
					PostId:  id,
					UserId:  1,
				}); err != nil {
					p.Redirect("/post/create", 302)
				}
			}
		}
		p.Redirect("/post/"+strconv.FormatInt(id, 10), 302)
	}

	p.Data["isNewRecord"] = true
	p.Data["post"] = models.Post{}
	p.Data["types"] = models.Menus()
	p.Data["labels"] = models.LabelList()
	p.Data["xsrf"] = template.HTML(p.XSRFFormHTML())
	p.Layout = "base.html"
	p.TplName = "post/create.html"
}

// View 文章详情
func (p *PostController) View() {
	id, err := strconv.Atoi(p.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		p.Redirect("/post", 302)
	}
	post, err := models.PostInfo(int64(id))
	if err != nil {
		p.Redirect("/post ", 302)
	}
	user, err := models.UserInfo(post.UserId)
	if err != nil {
		p.Redirect("/post", 302)
	}
	p.Data["post"] = post
	p.Data["user"] = user
	p.Data["type"] = models.Menus()
	p.Data["status"] = models.PostStatusDesc(post.Status)
	p.Data["labels"] = models.PostLabelDesc(1, post.Id)
	p.Data["ctime"] = time.Unix(post.Ctime, 0).Format("2006-01-02 15:04:05")
	p.Data["utime"] = time.Unix(post.Utime, 0).Format("2006-01-02 15:04:05")
	p.Layout = "base.html"
	p.TplName = "post/view.html"
}

// Update 编辑文章
func (p *PostController) Update() {
	id, err := strconv.Atoi(p.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		p.Redirect("/post", 302)
	}
	post, err := models.PostInfo(int64(id))
	if err != nil {
		p.Redirect("/post ", 302)
	}

	if p.Ctx.Request.Method == "POST" {
		typeId, err := strconv.Atoi(strings.TrimSpace(p.Input().Get("type")))
		if err != nil {
			p.Redirect("/post/update/"+strconv.FormatInt(int64(id), 10), 302)
		}
		post.Type = typeId
		post.Title = strings.TrimSpace(p.Input().Get("title"))
		post.Url = strings.TrimSpace(p.Input().Get("url"))
		post.ImageUrl = strings.TrimSpace(p.Input().Get("imageUrl"))
		post.Summary = strings.TrimSpace(p.Input().Get("summary"))
		post.Remark = strings.TrimSpace(p.Input().Get("remark"))
		var newId int64
		if newId, err = models.PostUpdate(&post); err != nil {
			p.Redirect("/post/update/"+strconv.FormatInt(newId, 10), 302)
		}
		// 如果有新标签, 则删除旧标签, 重新保存新标签, 判断比较麻烦, 故盲目删除旧标签重新保存一次
		labels := p.GetStrings("labels")
		if len(labels) > 0 {
			// 删除旧标签
			if _, err = models.PostLabelDelete(1, post.Id); err != nil {
				p.Redirect("/post/update/"+strconv.FormatInt(int64(id), 10), 302)
			}
			for _, labelId := range labels {
				var newLabelId int
				newLabelId, err = strconv.Atoi(labelId)
				if err != nil {
					p.Redirect("/post/update/"+strconv.FormatInt(int64(id), 10), 302)
				}
				if _, err := models.PostLabelSave(&models.PostLabel{
					LabelId: int64(newLabelId),
					PostId:  post.Id,
					UserId:  1,
				}); err != nil {
					p.Redirect("/post/update/"+strconv.FormatInt(int64(id), 10), 302)
				}
			}
		}
		p.Redirect("/post/"+strconv.FormatInt(newId, 10), 302)
	}

	p.Data["isNewRecord"] = false
	p.Data["post"] = post
	p.Data["types"] = models.Menus()
	p.Data["labels"] = models.LabelList()
	p.Data["hasLabels"] = models.PostLabelList(1, int64(id))
	p.Data["xsrf"] = template.HTML(p.XSRFFormHTML())
	p.Layout = "base.html"
	p.TplName = "post/update.html"
}

// Delete 删除文章
func (p *PostController) Delete() {
	id, err := strconv.Atoi(p.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		p.Redirect("/post", 302)
	}
	post, err := models.PostInfo(int64(id))
	if err != nil {
		p.Redirect("/post", 302)
	}
	post.Status = 2
	if _, err := models.PostUpdate(&post); err != nil {
		fmt.Println(err)
	}
	p.Redirect("/post", 302)
}
