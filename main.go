package main

import (
	"fmt"
)

/*
deps = "npm", "nodejs"
workspace = "https://github.com/scirats/scirats.git"
dots {
	repo = "https://github.com/heaveless/dotfiles.git"
	ext {
		git {
			config {
				user {
					email = "me@heaveless.com"
					name = "heaveless"
				}
				credential {
					helper = "store"
				}
			}
			credentials = "xyz", "jkl"
		}
	}
}*/
func main() {
	cfg, _ := ParseFile("data.xo")

	fmt.Println(cfg.StringList("deps"))
	fmt.Println(cfg.String("workspace"))
	fmt.Println(cfg.Block("dots").String("repo"))
	fmt.Println(cfg.Block("dots").Block("ext").Block("git").Block("config").Block("user").String("email"))
}
