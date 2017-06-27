# HermitCore Unikernels

UniK uses HermitCore as a platform for compiling Go, Fortran and C/C++ to unikernels.

---

### Golang
Compiling Go  

1. Setup a directory with the go source file and a Makefile describing the build process.

The directory should look similar to this structure:  

    $ tree ./gosrc  
     ./gosrc  
     ├── Makefile  
     └── server.go

In this example server.go is a simple webserver written in go:

    package main
    
    import (
        "fmt"
        "net/http"
    )

    func main() {
        fmt.Printf("HTTP Server written in Go using HermitCore and uniK\n")
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
    }

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from HermitCore")
    }

Makefile  

    all:
        x86_64-hermit-gccgo -O3 -pthread -Wall -o server.img server.go -lnetgo

2. Compile the code into an image by the following command:  
`unik build --name hermitcore_example --path ./gosrc --base hermitcore --language go --provider qemu`

3. Run the instance by invoking:  
`unik run --instanceName hermitcore_instance --imageName hermitcore_example`

4. You can validate that the webserver is running by inspecting port 8080. That is in the case your daemon is on your localhost:  
`http://localhost:8080`

Remark:
To allow the network connection to the server, the redirecting port (8080) has to be specified in the daemon-config.yaml

    qemu:
        - name: unikqemuprov
          redir_port: 8080

5. HermitCore's output:  
   HermitCore's output is written to a file in the tmp directory of the host. While the instance is running you can get these logs by invoking:  
`unik logs --instance hermitcore_instance`

---

### C/C++
Compiling C code works in the same way. A short example is explained in the following:

1. Setup a directory with the C source file and a Makefile describing the build process.

The directory should look similar to this structure:  

    $ tree ./csrc  
     ./csrc  
     ├── Makefile  
     └── hello.c

In this example hello.c prints 'Hello World':

    #include <stdio.h>
    
    int main()
    {
        printf("Hello World\n");
        while(1) {} // do not terminate to keep the instance running
        return 0;
    }
    
Makefile  

    all:
        x86_64-hermit-gcc -O3 -o hello.img hello.c

2. Compile the code into an image by the following command:  
`unik build --name hermitcore_example --path ./csrc --base hermitcore --language c --provider qemu`

3. Run the instance by invoking:  
`unik run --instanceName hermitcore_instance --imageName hermitcore_example`

4. HermitCore's output:  
   HermitCore's output is written to a file in the tmp directory of the host. While the instance is running you can get these logs by invoking:  
`unik logs --instance hermitcore_instance`

For compiling C++ code edit the Makefile such that `x86_64-hermit-g++` is used. Then build it with:  
`unik build --name hermitcore_example --path ./cppsrc --base hermitcore --language cpp --provider qemu`

### Fortran
The following example shows how to compile Fortran into a HermitCore unikernel:

1. Setup a directory with the Fortran source file and a Makefile describing the build process.

The directory should look similar to this structure:  

    $ tree ./fsrc  
     ./fsrc  
     ├── Makefile  
     └── hello.f90

In this example hello.f90 prints 'Hello World':

    program hello
        print *, "Hello World"
	    do
	    end do
    end program hello

Makefile  

    all:
	    x86_64-hermit-gfortran -O3 -o hello.img hello.f90
    
2. Compile the code into an image by the following command:  
`unik build --name hermitcore_example --path ./fsrc --base hermitcore --language fortran --provider qemu`

3. Run the instance by invoking:  
`unik run --instanceName hermitcore_instance --imageName hermitcore_example`

4. HermitCore's output:  
   HermitCore's output is written to a file in the tmp directory of the host. While the instance is running you can get these logs by invoking:  
`unik logs --instance hermitcore_instance`
