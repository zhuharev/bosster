package bosster

import "github.com/zhuharev/vkutil"

func PostToVk(credentials Credentials, post Post) (id int, err error) {
	// TODO: make *vkutil.Api pool
	u := vkutil.NewWithToken(credentials.VkAccessToken)
	u.SetDebug(true)
	postOptions := vkutil.OptsWallPost{
		OwnerId:   credentials.VkOwnerID,
		FromGroup: true,
		Message:   post.Body,
		ImageURLs: post.ImageURLs,
	}
	return u.WallPost(postOptions)
}
