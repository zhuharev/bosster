package bosster

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func PostToFacebook(credentials Credentials, post Post) (fbID string, err error) {
	var imageIDs []string
	for _, v := range post.ImageURLs {
		imageFbID, err := facebookUploadUnpablishedImage(credentials.FacebookToken, credentials.FacebookGroupID, v)
		if err != nil {
			return "", err
		}
		imageIDs = append(imageIDs, imageFbID)
	}
	return facebookPost(credentials.FacebookToken, credentials.FacebookGroupID, post.Body, imageIDs)
}

func facebookUploadUnpablishedImage(token string, groupID string, imageURL string) (fbID string, err error) {
	var (
		endpoint = fmt.Sprintf("/%s/photos", groupID)
	)
	res, err := fb.Post(endpoint, fb.Params{
		"access_token": token,
		"name":         "noop",
		"url":          imageURL,
		"published":    "false",
	})
	if err != nil {
		return
	}
	fbID = res.GetField("id").(string)
	return
}

func facebookPost(token, groupID, message string, images []string) (fbID string, err error) {
	var (
		endpoint = fmt.Sprintf("/%s/feed", groupID)
	)
	params := fb.Params{
		"access_token": token,
		"message":      message,
	}
	for i, v := range images {
		key := fmt.Sprintf("attached_media[%d]", i)
		params[key] = fmt.Sprintf(`{"media_fbid":"%s"}`, v)
	}
	res, err := fb.Post(endpoint, params)
	if err != nil {
		return
	}
	fbID = res.GetField("id").(string)
	return
}
