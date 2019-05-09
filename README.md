# go_webApp
This is a Go web application which gets connected to MongoDB (noSQL database).
The go_webApp runs as a service listening in por 8008.

## Instrucions to install Go and Mongo
First things first... 
1. Install Golang
   - Follow the instructions in the [Golang's web page](https://golang.org/doc/install) to reach this.
     - I also found an exceptional book to learn more about Golang and it's manual installation in the [The Little Go Book](https://www.openmymind.net/The-Little-Go-Book/) web page. This resource is very exciting way to fast learn this wonderfull language.
   - Once you installed Go, set your workspace and tested it is up and running clone this project and it's dependencies:
      >```
      >~$ go get github.com/cibervicho/go_webApp
      >~$ go get github.com/globalsign/mgo
      >```
     Usualy, when downloading the main `cibervicho/go_webApp` repository, Go downloads automatically its dependencies, such as `globalsign/mgo`.
2. Install MongoDB CE (MongoDB Community Edition)
   - Follow the instructions in the [MongoDB's installation tutorial](https://docs.mongodb.com/manual/installation/#tutorial-installation).
   - The manual installation instructions from the book [The Little MongoDB Book](https://www.openmymind.net/2011/3/28/The-Little-MongoDB-Book/) are also a good resource.

## Testing Go
To verify that you installed Golang correctly just type the following. You should see the same version of the programing language you installed:
   >```
   >~$ go version
   >go version go1.12.4 linux/amd64
   >```
If you received a message somwhat like the above then you are good to Go! :)

Now, go ahead and create a `hello.go` file with the following content in it:
>```
>package main
>
>import "fmt"
>
>func main() {
>  fmt.Println("Hello Universe!")
>}
>```
Save the file and try to run it as follows:
>```
>$ go run hello.go 
>Hello Universe!
>```
If you see the `Hello Universe!` message printed in the screen, you just double checked that your Go installation is actually compiling and running Go programs correctly. Congratulations!
   
## Configuring and Testing Mongo
If you installed via the package manager (following the mongoDB installation instructions), the data directory `/var/lib/mongodb` and the log directory `/var/log/mongodb` are created during the installation.
To work with MongoDB two entities are required:
1. The Server and
2. The Client
### MongoDB: The Server
The official MongoDB package includes a configuration file `/etc/mongod.conf`. These settings (such as the data directory and log directory specifications) take effect upon startup. That is, if you change the configuration file while the MongoDB instance is running, you must restart the instance for the changes to take effect.
Go ahead and update this `mongod.conf` file and be sure to `dbPath` under `storage:` and the `path` under the `systemLog:` sections are pointing to the correct data and log directories previously mentioned [here]().

## Installing Jenkins and the required plugins
### Explaining the bash scripts in Jenkins
