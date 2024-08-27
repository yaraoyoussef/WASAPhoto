# WASAPhoto
This repository contains the WASAPhoto project, developed as part of the Web and Software Architecture (WASA) course. WASAPhoto is a web application designed to help users stay connected by sharing photos of special moments. It allows users to upload photos, follow friends, and interact with shared content.

## Project Overview
WASAPhoto was developed with the following key technologies:
- Backend: Go (Golang)
- Frontend: JavaScript
- API Definition: OpenAPI standard
- Deployment: Docker container

## Key Features
- Photo Stream: Users see a stream of photos from people they follow, ordered by the most recent uploads. Each photo displays the upload date, time, likes, and comments.
- User Interactions: Users can like, comment, and remove their comments on photos. They can also ban or unban other users to control visibility.
- User Profiles: Each user has a profile page showing their uploaded photos, follower/following count, and profile management options.
- Simplified Login: Users can log in with just a username. New users are automatically registered upon their first login.

## Simplified Login
In this project, a simplified login mechanism is implemented. Users only need to provide a username to log in or register. This approach bypasses traditional authentication complexities, streamlining the development process.

## API Specification
The project defines a set of RESTful APIs following the OpenAPI standard. These include:
- User Authentication: Simple login API without password.
- User Management: Endpoints for setting usernames, following/unfollowing users, and banning/unbanning users.
- Photo Management: APIs for uploading, deleting, liking, unliking, commenting, and uncommenting photos.
- Profile and Stream: Endpoints to retrieve user profiles and photo streams.

## Deployment
WASAPhoto can be easily deployed using Docker. A Docker container image is provided, ensuring consistent deployment across different environments.
