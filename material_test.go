package weixin

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestUploadTemporaryMaterial(t *testing.T) {
	img, err := os.Open("/Users/mac/Pictures/docker.jpg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	mediaId, createAt, err := UploadTemporaryMaterial(MediaTypeImage, img)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%#v %#v", mediaId, createAt)
}

func TestGetTemporaryMaterial(t *testing.T) {
	mediaId := "Yh_M-L4x4e4BYLXlAqEKfMvAd6mT7AVSrOmNP996xsQJsF4nncxSuICi7WkV25e3"
	filename, content, err := GetTemporaryMaterial(mediaId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("write file to %s", filename)
}

func TestGetNews(t *testing.T) {
	type args struct {
		mediaId string
	}
	//gotRet, err := GetNews(tt.args.mediaId)
}

func TestUploadMaterialUploadnews(t *testing.T) {
	type args struct {
		news MaterialNews
	}
	ars := args{
		MaterialNews{
		Articles: []Articles{
			{
				Title:            "233232",
				ThumbMediaId:     "Yh_M-L4x4e4BYLXlAqEKfMvAd6mT7AVSrOmNP996xsQJsF4nncxSuICi7WkV25e3",
				Author:           "233223",
				Content:          "1212121212",
				ContentSourceURL: "http://www.baidu.com",
			},
		},
	}}
	//gotMediaId, err := AddNews(ars.news.Articles)
	//fmt.Println(gotMediaId, err)

	gotMediaId, gotCreateAt, err := UploadMaterialUploadnews(ars.news)
	fmt.Println(gotMediaId,gotCreateAt,  err)
	//92zIVNrMZxIt5CEhZ0eyqEgk4E1I6yZ_wol52zOJ-yArsps8XlC5irjdHacwJv8u
	//TnidRQxM-SBNu7k7jIoTlD3Mq9X-on_np7CcxxHD2J09i7Z_RlUNNKyxxoGCsrgI
}

func TestUploadMaterialPreviewNews(t *testing.T) {
	type args struct {
		news PreviewNewsMsg
	}
	news := PreviewNewsMsg{
		Touser: "ollXS5kW3dYczxiJHEypRv1kJLcY",
		Mpnews: struct {
			MediaID string `json:"media_id"`
		}{
			"92zIVNrMZxIt5CEhZ0eyqEgk4E1I6yZ_wol52zOJ-yArsps8XlC5irjdHacwJv8u",
		},
		Msgtype: "mpnews",
	}
	UploadMaterialPreviewNews(news)
}