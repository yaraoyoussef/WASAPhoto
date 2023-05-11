package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users/:id", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:id", rt.wrap(rt.setMyUsername))
	rt.router.GET("/users/:id/home/", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:id/follow/:otherUserId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:id/follow/:otherUserId", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:id/ban/:otherUserId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:id/ban/:otherUserId", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:id/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:id/photos/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.PUT("/users/:id/photos/:photoId/likes/:otherUserId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:id/photos/:photoId/likes/:otherUserId", rt.wrap(rt.unlikePhoto))
	rt.router.POST("/users/:id/photos/:photoId/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:id/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
