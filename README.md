# My-App

- Create a Go app
  - run `go mod init github.com/rajatgupta24/go-app-sbom`
  - 


- Building Hello world project
  - go mod init
  - create a main.go and add a func
  - run go build .
- Create Docker Image
  - Dockerfile
    - FROM (as base image)
    - WORKDIR (pwd for docker)
    - COPY work-repo docker-repo
    - RUN command
    - ENTRYPOINT
  - docker build -t <tag> .
- Makefile
  - a file use to create shortcuts in terminal



- Docker pull -> be used to fetch docker images to local system.