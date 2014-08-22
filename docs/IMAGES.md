# Images

Images are referring to `docker` images used to execute code inside a `docker` container.

They can be found inside the `dockerfiles` directory:

```
dockerfiles
│
├───exec-ruby
│   │   Dockerfile
│   │   run.sh
│
│
└───exec-golang
    │   Dockerfile
    │   run.sh
```

Each image has a directory called `exec-$language`.
Inside this directory there is two files:

- A `Dockerfile` to build the image.
- A shell script `run.sh` that must be copied inside the image.


Images are built has an executable `docker` image. This allow us to do:

    $ docker run grounds/exec-ruby "puts 42"
    42
    
## Create an image

Creating an image for a new language is really trivial.
Take a look at this example for the C language:

Create the directory:

    mkdir dockerfiles/exec-c
    
Add a `Dockerfile` and a shell script inside this directory:

    touch dockerfiles/exec-c/Dockerfile
    touch dockerfiles/exec-c/run.sh

### Inside the Dockerfile:

Base the image on the official ubuntu image:

    FROM ubuntu:14.04

Update ubuntu package manager:

    RUN apt-get update -qq

Install dependencies required to compile C code (e.g `gcc`)

    RUN apt-get -qy install \
        build-essential \
        gcc

Copy the script `run.sh` inside the `/home/dev` directory:

    COPY run.sh /home/dev/run.sh

Add a user and give it access to `/home/dev`

    RUN useradd dev
    RUN chown -R dev: /home/dev

Add:

    WORKDIR /home/dev
    USER dev

    ENTRYPOINT ["/home/dev/run.sh"]

When you run a `docker` container with this image:
- The default `pwd` of the container will be `/home/dev`.
- The user of the container will be `dev`
- The container will run `run.sh` with parameters received from the CLI.

### Inside the shell script:

First make it a shell script:

    #!/bin/sh

Echo first parameter from CLI to a file runnable by the language compiler/interpreter:

    echo "$1" > prog.c
    
Please don't forget to surround `$1` with quotation marks, to avoid unexpected behaviors.
    
Compile and/or run the program:

    gcc -o prog prog.c
    
    if [ -f "prog" ]
    then
      ./prog
    fi
    
### Build the image

Use `docker` to build the image like you usually do:

    $ docker build -t grounds/exec-c dockerfiles/exec-c