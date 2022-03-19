# spotigo

The library that allows you to access the Spotify API I developed for my other project, Radiofy.

Attention users, by the time you read this repo it is out of date and may have bugs!

https://developer.spotify.com/documentation/web-api/reference/#/


Download:
```
# go get github.com/batuhannoz/spotigo
```

Usage:
```go
package main

import (
	"fmt"
	player "github.com/batuhannoz/spotigo/player"
	user "github.com/batuhannoz/spotigo/user"
)

responsePLayer, err := player.SeekToPosition(player.SeekToPositionInput{
		Token:      "Bearer <<Token>>",
		PositionMS: 100000,
	})
  
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responsePLayer)

	responseUser, err := user.GetCurrentUsersProfile(user.GetCurrentUsersProfileInput{
		Token: "Bearer <<Token>>",
	})
  
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(responseUser)  
}
```
