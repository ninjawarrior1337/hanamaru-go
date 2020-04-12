// +build jp

package main

import "hanamaru/commands/jp"

func init() {
	optionals = append(optionals, jp.Romaji)
}