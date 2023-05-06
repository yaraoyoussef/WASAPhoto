package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))                                      // tested
	rt.router.GET("/users/:username", rt.wrap(rt.getUserProfile))                        // tested
	rt.router.PUT("/users/:username", rt.wrap(rt.setMyUsername))                         // tested
	rt.router.GET("/users/:username/home/", rt.wrap(rt.getMyStream))                     // tested
	rt.router.PUT("/users/:username/follow/:otherUsername", rt.wrap(rt.followUser))      // tested
	rt.router.DELETE("/users/:username/follow/:otherUsername", rt.wrap(rt.unfollowUser)) // tested
	rt.router.PUT("/users/:username/ban/:otherUsername", rt.wrap(rt.banUser))            // tested
	rt.router.DELETE("/users/:username/ban/:otherUsername", rt.wrap(rt.unbanUser))       // tested
	rt.router.POST("/users/:username/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/photos/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.PUT("/users/:username/photos/:photoId/likes/:otherUsername", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/photos/:photoId/likes/:otherUsername", rt.wrap(rt.unlikePhoto))
	rt.router.POST("/users/:username/photos/:photoId/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
