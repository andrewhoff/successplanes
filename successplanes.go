package successplanes

import (
	"time"

	"github.com/go/src/math/rand"
)

const (
	IDX_SR71 = iota
	IDX_F117
	IDX_A10
	IDX_F18
	IDX_AVRO
)

// SR71 - returns SR71 in ascii art
func SR71() string {
	return successPlanes[IDX_SR71]
}

// F117 - returns F117 in ascii art
func F117() string {
	return successPlanes[IDX_F117]
}

// A10 - returns A10 in ascii art
func A10() string {
	return successPlanes[IDX_A10]
}

// F18 - returns F18 in ascii art
func F18() string {
	return successPlanes[IDX_F18]
}

// AVRO - returns AVRO in ascii art
func AVRO() string {
	return successPlanes[IDX_AVRO]
}

// Rand - returns a random ascii plane
func Rand() string {
	rand.Seed(time.Now().UnixNano())
	return successPlanes[rand.Int()%len(successPlanes)]
}

// I did not make these ASCII planes, I borrowed them from https://asciiart.website//index.php?art=transportation/airplanes
var successPlanes = []string{
	`
                     .                             .
                    //                             \\
                   //                               \\
                  //                                 \\
                 //                _._                \\
              .---.              .//|\\.              .---.
    _________/ .-. \_________..-~ _.-._ ~-..________ / .-. \_________
             \ ~-~ /   /H-     '-=.___.=-'     -H\   \ ~-~ /
               ~~~    / H          [H]          H \    ~~~
                     / _H_         _H_         _H_ \
					   UUU         UUU         UUU`,
	`     ^
         / \
        //V\\
       / \|/ \
      /// v \\\
     /         \
    /           \
   /             \
  /   /|     |\   \
 /   / \     / \   \
 \  /   X   X   \  /
  \/   / \ / \   \/
      /   V   \
	 |         |`,
	`                     ||   ___     ___     ___   ||
                          ||  /   \   /| |\   /   \  ||
                          || |  O  |__|| ||__|  O  | ||
                          ||  \___/--/^^^^^\--\___/  ||
                      __  ||________|       |________||  __
   .-----------------/  \-++--------|   .   |--------++-/  \-----------------.
  /.---------________|  |___________\__(*)__/___________|  |________---------.\
            |    |   '$$'   |                       |   '$$'   |    |
           (o)  (o)        (o)                     (o)        (o)  (o)`,
	`           \     /
                 \ _ /
              ----/_\----
  x--------------( . )--------------x
       x|x   | |_|\_/|_| |   x|x
        x    x           x    x     
`,
	`       /\
           /  \
           |  |
           |/\|
           |00|
          _|oo|_
         | |  | |
         / |  | \
        /  |  |  \
       /   |  |   \
      /    |  |    \
     /     |  |     \
   //      |  |      \\
  /        |  |        \
 /   @     |  |      @  \
/          ||||          \
|          ||||           |
|__________||||___________|
          |_||_|
            ||`,
}
