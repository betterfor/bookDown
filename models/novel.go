package models

type SearchNovelSlice []Novel

type Novel struct {
	NovelName    string    // 小说名
	Author       string    // 作者
	NovelURL     string    // 小说连接
	UpdatedTime  string    // 更新时间
	Introduction string    // 小说简介
	Chapters     []Chapter // 小说章节
	Image        string    // 小说封面
}

type Chapter struct {
	ChapterIndex int      // 章节序号
	ChapterName  string   // 章节名称
	LinkURL      string   // 章节链接
	Content      []string // 章节内容
}
