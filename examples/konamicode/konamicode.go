/*

Simple game whose only goal is to type the Konami Code (https://en.wikipedia.org/wiki/Konami_Code)

The game runs forever until the code is entered in the right order, and a player does "well" by typing it sooner

*/
package konamicode

type KonamiCodeEnvironment struct{}

func New() KonamiCodeEnvironment {
	return KonamiCodeEnvironment{}
}
func (k KonamiCodeEnvironment) Reset() {

}
