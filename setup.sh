go run cmd/jwg/main.go -type Sample -output misc/fixture/a/model_json.go misc/fixture/a
go run cmd/jwg/main.go -type Sample -output misc/fixture/b/model_json.go misc/fixture/b
go run cmd/jwg/main.go -type Sample -output misc/fixture/c/model_json.go misc/fixture/c
go run cmd/jwg/main.go -output misc/fixture/d/model_json.go misc/fixture/d
go run cmd/jwg/main.go -type Sample -output misc/fixture/e/model_json.go misc/fixture/e
go run cmd/jwg/main.go -type SampleF -output misc/fixture/f/model_json.go misc/fixture/f
go run cmd/jwg/main.go -type Sample,Inner -output misc/fixture/g/model_json.go misc/fixture/g
go run cmd/jwg/main.go -type Sample -output misc/fixture/h/model_json.go misc/fixture/h
go run cmd/jwg/main.go -output misc/fixture/i/model_json.go misc/fixture/i
go run cmd/jwg/main.go -output misc/fixture/j/model_json.go misc/fixture/j
go run cmd/jwg/main.go -output misc/fixture/k/model_json.go misc/fixture/k
go run cmd/jwg/main.go -type Sample -output misc/fixture/l/model_json.go -transcripttag swagger,includes misc/fixture/l
