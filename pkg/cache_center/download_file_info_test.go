package cache_center

import (
	"testing"

	"github.com/An-Yan-d/ChineseSubFinder/pkg/types/language"
	"github.com/An-Yan-d/ChineseSubFinder/pkg/types/supplier"

	"github.com/An-Yan-d/ChineseSubFinder/pkg/log_helper"
)

func TestCacheCenter_DownloadFileAdd(t *testing.T) {
	cc := NewCacheCenter("testFile", log_helper.GetLogger4Tester())

	subInfo := supplier.NewSubInfo(
		"test",
		1,
		"name",
		language.ChineseSimple,
		"url123123",
		0,
		0,
		"ext",
		[]byte{1, 2, 3, 4, 5},
	)
	err := cc.DownloadFileAdd(subInfo)
	if err != nil {
		t.Fatal(err)
	}

	bok, getSubInfo, err := cc.DownloadFileGet(subInfo.GetUID())
	if err != nil {
		t.Fatal(err)
	}
	if bok == false {
		t.Fatal("bok == false")
	}

	if subInfo.FileUrl != getSubInfo.FileUrl {
		t.Fatal("subInfo.FileUrl != getSubInfo.FileUrl")
	}
}
