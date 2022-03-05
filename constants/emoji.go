package constants

import "errors"

const X = ":x:"
const CHECK_MARK = ":white_check_mark:"
const GERMAN_FLAG = ":flag_de:"
const MOON = ":full_moon_with_face:"
const HELMET = ":military_helmet:"
const ROBOT = ":robot:"
const ALIEN = ":space_invader:"
const BRANCH = ":herb:"
const SNAKE = ":snake:"
const MOVIE_CAMERA = ":movie_camera:"
const ROCKET = ":rocket:"
const TOILET = ":toilet:"

func ZMapToEmoji(mapName string) (string, error) {
	if !MapExists(mapName) {
		return "", errors.New(mapName + "doesn't exist as a map.")
	}

	zMapEmoji := map[string]string{
		"der riese":         GERMAN_FLAG,
		"moon":              MOON,
		"nacht der untoten": HELMET,
		"origins":           ROBOT,
		"shadows of evil":   ALIEN,
		"shang ri la":       BRANCH,
		"shi no numa":       SNAKE,
		"kino der toten":    MOVIE_CAMERA,
		"ascension":         ROCKET,
		"verruckt":          TOILET,
	}

	return zMapEmoji[mapName], nil 
}