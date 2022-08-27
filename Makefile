build:
	go build -o cardgen
gen:
	./cardgen -c example/config.yaml example/blog-post.md
start:
	gen build