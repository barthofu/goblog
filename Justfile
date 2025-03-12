shell:
	nix develop --extra-experimental-features nix-command --extra-experimental-features flakes -c zsh

run:
	go run .

build:
	go build .