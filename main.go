package main

import (
	"fmt"
	spotify "github.com/batuhannoz/spotigo/player"
)

func main() {
	res, err := spotify.GetAvailableDevices(spotify.GetAvailableDevicesInput{Token: "Bearer BQBxQQvi6B4xu3YXbaE2hqwl8oMrVK4DzbueIYvk1eml_CM-fMx_fs4seq92MisdCURRSXShZUk_D10yFSZMzRMGtem_UHHEFyiQooG0nDkUkOLwEKANAl3etQa30BR8JImVim6aGZeYJkACYjwuLC1mGktyVDa32LaWEpka50hqvrOS0kIePQkDf5PEWtoFiIikNssGn1BnJto"})
	fmt.Println("-------------------------------------------------------")
	fmt.Println(err)
	fmt.Println("-------------------------------------------------------")
	fmt.Println(res)
	fmt.Println("-------------------------------------------------------")

}
