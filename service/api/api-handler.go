package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.GET("/users/:username", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:username", rt.wrap(rt.setMyUsername))
	rt.router.GET("/users/:username/home/", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:username/follow/:otherUsername", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/follow/:otherUsername", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:username/ban/:otherUsername", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:otherUsername", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:username/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
