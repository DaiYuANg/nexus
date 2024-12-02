package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/lo"
	"nexus/cmd"
)

func main() {
	lo.Must0(cmd.Execute())
}
