package main

import "godis/tcp"

var banner = `
   ______          ___
  / ____/___  ____/ /______
 / / __/ __ \/ __  / / ___/
/ /_/ / /_/ / /_/ / (__  )
\____/\____/\__,_/_/____/
----------------------------
----------------------------
`

func main() {
	print(banner)
	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address: ":6399",
	}, tcp.MakeEchoHandler())
	if err != nil {
		return
	}
}
