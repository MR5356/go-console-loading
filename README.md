
[![go-report](https://goreportcard.com/badge/github.com/MR5356/go-console-loading)](https://goreportcard.com/report/github.com/MR5356/go-console-loading)
![contributors](https://img.shields.io/github/contributors/MR5356/go-console-loading)
![used-by](https://img.shields.io/sourcegraph/rrc/github.com/MR5356/go-console-loading)
[![license](https://img.shields.io/github/license/MR5356/go-console-loading)](https://www.apache.org/licenses/LICENSE-2.0.html)


## Usage
```go
package main

import (
	"fmt"
	"github.com/MR5356/go-console-loading/loading"
	"os"
	"time"
)

func main() {
	ld := loading.New(os.Args[1])

	ld.Start("Processing...")
	time.Sleep(time.Second * 10)
	fmt.Println()
	ld.Stop()
}

```
![loading-runner](img/loading-runner.gif)

## Supported spinners
| name                | spinner                                                |
|---------------------|--------------------------------------------------------|
| aesthetic           | ![loading-runner](img/loading-aesthetic.gif)           |
| arc                 | ![loading-runner](img/loading-arc.gif)                 |
| arrow               | ![loading-runner](img/loading-arrow.gif)               |
| arrow2              | ![loading-runner](img/loading-arrow2.gif)              |
| arrow3              | ![loading-runner](img/loading-arrow3.gif)              |
| balloon             | ![loading-runner](img/loading-balloon.gif)             |
| balloon2            | ![loading-runner](img/loading-balloon2.gif)            |
| betaWave            | ![loading-runner](img/loading-betaWave.gif)            |
| binary              | ![loading-runner](img/loading-binary.gif)              |
| bluePulse           | ![loading-runner](img/loading-bluePulse.gif)           |
| bounce              | ![loading-runner](img/loading-bounce.gif)              |
| bouncingBall        | ![loading-runner](img/loading-bouncingBall.gif)        |
| bouncingBar         | ![loading-runner](img/loading-bouncingBar.gif)         |
| boxBounce           | ![loading-runner](img/loading-boxBounce.gif)           |
| boxBounce2          | ![loading-runner](img/loading-boxBounce2.gif)          |
| christmas           | ![loading-runner](img/loading-christmas.gif)           |
| circle              | ![loading-runner](img/loading-circle.gif)              |
| circleHalves        | ![loading-runner](img/loading-circleHalves.gif)        |
| circleQuarters      | ![loading-runner](img/loading-circleQuarters.gif)      |
| clock               | ![loading-runner](img/loading-clock.gif)               |
| dots                | ![loading-runner](img/loading-dots.gif)                |
| dots10              | ![loading-runner](img/loading-dots10.gif)              |
| dots11              | ![loading-runner](img/loading-dots11.gif)              |
| dots12              | ![loading-runner](img/loading-dots12.gif)              |
| dots13              | ![loading-runner](img/loading-dots13.gif)              |
| dots2               | ![loading-runner](img/loading-dots2.gif)               |
| dots3               | ![loading-runner](img/loading-dots3.gif)               |
| dots4               | ![loading-runner](img/loading-dots4.gif)               |
| dots5               | ![loading-runner](img/loading-dots5.gif)               |
| dots6               | ![loading-runner](img/loading-dots6.gif)               |
| dots7               | ![loading-runner](img/loading-dots7.gif)               |
| dots8               | ![loading-runner](img/loading-dots8.gif)               |
| dots8Bit            | ![loading-runner](img/loading-dots8Bit.gif)            |
| dots9               | ![loading-runner](img/loading-dots9.gif)               |
| dqpb                | ![loading-runner](img/loading-dqpb.gif)                |
| dwarfFortress       | ![loading-runner](img/loading-dwarfFortress.gif)       |
| earth               | ![loading-runner](img/loading-earth.gif)               |
| fingerDance         | ![loading-runner](img/loading-fingerDance.gif)         |
| fistBump            | ![loading-runner](img/loading-fistBump.gif)            |
| flip                | ![loading-runner](img/loading-flip.gif)                |
| grenade             | ![loading-runner](img/loading-grenade.gif)             |
| growHorizontal      | ![loading-runner](img/loading-growHorizontal.gif)      |
| growVertical        | ![loading-runner](img/loading-growVertical.gif)        |
| hamburger           | ![loading-runner](img/loading-hamburger.gif)           |
| hearts              | ![loading-runner](img/loading-hearts.gif)              |
| layer               | ![loading-runner](img/loading-layer.gif)               |
| line                | ![loading-runner](img/loading-line.gif)                |
| line2               | ![loading-runner](img/loading-line2.gif)               |
| material            | ![loading-runner](img/loading-material.gif)            |
| mindblown           | ![loading-runner](img/loading-mindblown.gif)           |
| monkey              | ![loading-runner](img/loading-monkey.gif)              |
| moon                | ![loading-runner](img/loading-moon.gif)                |
| noise               | ![loading-runner](img/loading-noise.gif)               |
| orangeBluePulse     | ![loading-runner](img/loading-orangeBluePulse.gif)     |
| orangePulse         | ![loading-runner](img/loading-orangePulse.gif)         |
| pipe                | ![loading-runner](img/loading-pipe.gif)                |
| point               | ![loading-runner](img/loading-point.gif)               |
| pong                | ![loading-runner](img/loading-pong.gif)                |
| runner              | ![loading-runner](img/loading-runner.gif)              |
| sand                | ![loading-runner](img/loading-sand.gif)                |
| shark               | ![loading-runner](img/loading-shark.gif)               |▌|
| simpleDots          | ![loading-runner](img/loading-simpleDots.gif)          |
| simpleDotsScrolling | ![loading-runner](img/loading-simpleDotsScrolling.gif) |
| smiley              | ![loading-runner](img/loading-smiley.gif)              |
| soccerHeader        | ![loading-runner](img/loading-soccerHeader.gif)        |
| speaker             | ![loading-runner](img/loading-speaker.gif)             |
| squareCorners       | ![loading-runner](img/loading-squareCorners.gif)       |
| squish              | ![loading-runner](img/loading-squish.gif)              |
| star                | ![loading-runner](img/loading-star.gif)                |
| star2               | ![loading-runner](img/loading-star2.gif)               |
| timeTravel          | ![loading-runner](img/loading-timeTravel.gif)          |
| toggle              | ![loading-runner](img/loading-toggle.gif)              |
| toggle10            | ![loading-runner](img/loading-toggle10.gif)            |
| toggle11            | ![loading-runner](img/loading-toggle11.gif)            |
| toggle12            | ![loading-runner](img/loading-toggle12.gif)            |
| toggle13            | ![loading-runner](img/loading-toggle13.gif)            |
| toggle2             | ![loading-runner](img/loading-toggle2.gif)             |
| toggle3             | ![loading-runner](img/loading-toggle3.gif)             |
| toggle4             | ![loading-runner](img/loading-toggle4.gif)             |
| toggle5             | ![loading-runner](img/loading-toggle5.gif)             |
| toggle6             | ![loading-runner](img/loading-toggle6.gif)             |
| toggle7             | ![loading-runner](img/loading-toggle7.gif)             |
| toggle8             | ![loading-runner](img/loading-toggle8.gif)             |
| toggle9             | ![loading-runner](img/loading-toggle9.gif)             |
| triangle            | ![loading-runner](img/loading-triangle.gif)            |
| weather             | ![loading-runner](img/loading-weather.gif)             |

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=MR5356/go-console-loading&type=Date)](https://star-history.com/#MR5356/go-console-loading&Date)