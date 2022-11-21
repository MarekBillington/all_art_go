module allart

go 1.17

replace filemanager => ./filemanager

replace imagebuilder => ./imagebuilder

require (
	filemanager v0.0.0-00010101000000-000000000000
	imagebuilder v0.0.0-00010101000000-000000000000
)

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
)
