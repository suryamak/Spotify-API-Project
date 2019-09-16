# Messing Around With Spotify API

## Concept
  Compare two users and determine how similar their music taste is
 
## Method
  - Playlists
  - Liked Songs
  - Most played artists and songs (I think this requires users' credentials but would be most accurate)

## Setup
  - Register for a Client ID and Client Secret using your Spotify account [here](https://developer.spotify.com/dashboard/login)
  - Install [Go](https://golang.org/dl/)
  - Run ```go get github.com/suryamak/Spotify-API-Project``` to install dependencies
  - Create a file called ```.env``` in the root directory and add the following to it:
    - CLIENT_ID=your_client_id
    - CLIENT_SECRET=your_client_secret
  - Run the application from the root directory with ```go run cmd/main.go```
