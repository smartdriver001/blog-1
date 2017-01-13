package controllers

import (
	"blog/app/models"
	"blog/app/service"

	"github.com/revel/revel"
)

//Main controller.
type Main struct {
	*revel.Controller
}

//SiteInfo model
type SiteInfo struct {
	Title     string
	SubTitle  string
	Copyright string
}

// Main page.
// 博客首页
func (m *Main) Main() revel.Result {
	set := new(models.Setting)
	set.Key = "site-title"
	title, _ := set.Get()
	set.Key = "site-subtitle"
	subtitle, _ := set.Get()
	set.Key = "site-foot"
	copyr, _ := set.Get()

	// 页面处理
	var p int
	m.Params.Bind(&p, "page")
	if p == 0 {
		p = 1
	}
	blogModel := new(models.Blogger)
	blogs, err := blogModel.GetBlogByPage(p, 0)
	if err != nil {
		revel.ERROR.Println("加载博客列表失败: ", err)
		return m.RenderError(err)
	}
	hotBLogs := blogModel.GetHotBlog(5)

	pageStruct := new(service.BlogPager)

	m.RenderArgs["blogs"] = blogs
	m.RenderArgs["hotblogs"] = hotBLogs
	m.RenderArgs["thisPage"] = p
	m.RenderArgs["pager"] = pageStruct.GetPager(p)
	category := new(models.Category)
	m.RenderArgs["categorys"] = category.FindAll()
	info := &SiteInfo{Title: title, SubTitle: subtitle, Copyright: copyr}
	return m.Render(info)
}

// 某个分类下的博客
func (m *Main) Blog4Category(ca string) revel.Result {
	blog := new(models.Blogger)
	category := new(models.Category)
	id := category.GetByIdent(ca)
	var blogs *[]models.Blogger
	if id > 0 {
		blogs, _ = blog.FindByCategory(id)
	}
	m.RenderArgs["blogs"] = blogs
	return m.RenderTemplate("Main/Blog4Search.html")
}
