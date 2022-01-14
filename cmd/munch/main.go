package main

import (
	"github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/scene"
	"github.com/plyr4/munch/logger"
)

func main() {

	// TODO: scene management
	//       game scene:
	//           - init
	//           - update
	//           - draw
	//           - action binds
	//           - repeat
	oak.AddScene("firstScene", scene.Scene{
		Start: func(*scene.Context) {
			// ... draw entities, bind callbacks ...
			// draw menu
			// bind menu inputs
		},
	})

	logger.Logger().Info("running firstScene")
	oak.Init("firstScene")
}
