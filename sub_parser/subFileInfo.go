package sub_parser

import "github.com/allanpk716/ChineseSubFinder/common"

type SubFileInfo struct {
	Name	string			// 字幕的名称，注意，这里需要额外的赋值，不会自动检测
	Ext		string			// 字幕的后缀名
	Lang common.Language	// 识别出来的语言
	Dialogues []OneDialogue	// 整个字幕文件的所有对话
}

// OneDialogue 一句对话
type OneDialogue struct {
	StartTime string		// 开始时间
	EndTime string			// 结束时间
	Lines	[]string		// 台词
}