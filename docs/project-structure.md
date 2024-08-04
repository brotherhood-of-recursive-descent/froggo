We decide to follow the go modules server layout.
The advantage for clarity is that no top-level golang files are reusable
and that reusable files have to be packaged in libraries.

.
├── README.md
├── cmd
│   └── froggo
│       └── main.go
├── docs
│   └── no-sprites.md
├── game
├── go.mod
├── lib
└── resources

