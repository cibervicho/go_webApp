# go_webApp
This is a Go web application which gets connected to MongoDB (noSQL database).
The go_webApp runs as a service listening in por 8008.

## Instrucions to install Go and Mongo
First things first... 
1. Install Golang
   - Follow the instructions in the [Golang's web page](https://golang.org/doc/install) to reach this.
     - I also found an exceptional book to learn more about Golang and it's manual installation in the [The Little Go Book](https://www.openmymind.net/The-Little-Go-Book/) web page. This resource is very exciting way to fast learn this wonderfull language.
   - Once you installed Go, set your workspace and tested it is up and running clone this project and it's dependencies:
      ```sh
      ~$ go get github.com/cibervicho/go_webApp
      ~$ go get github.com/globalsign/mgo
      ```
     Usualy, when downloading the main `cibervicho/go_webApp` repository, Go downloads automatically its dependencies, such as `globalsign/mgo`.
2. Install MongoDB CE (MongoDB Community Edition)
   - Follow the instructions in the [MongoDB's installation tutorial](https://docs.mongodb.com/manual/installation/#tutorial-installation).
   - The manual installation instructions from the book [The Little MongoDB Book](https://www.openmymind.net/2011/3/28/The-Little-MongoDB-Book/) are also a good resource.

## Testing Go
To verify that you installed Golang correctly just type the following. You should see the same version of the programing language you installed:
   ```sh
   ~$ go version
   go version go1.12.4 linux/amd64
   ```
If you received a message somewhat like the above then you are good to Go! :)

Now, go ahead and create a `hello.go` file with the following content in it:
```sh
package main

import "fmt"

func main() {
  fmt.Println("Hello Universe!")
}
```
Save the file and try to run it as follows:
```sh
~$ go run hello.go 
Hello Universe!
```
If you see the `Hello Universe!` message printed in the screen, you just double checked that your Go installation is actually compiling and running Go programs correctly. Congratulations!
   
## Configuring and Testing Mongo
If you installed via the package manager (following the mongoDB installation instructions), the data directory `/var/lib/mongodb` and the log directory `/var/log/mongodb` are created during the installation.
To work with MongoDB two entities are required:
1. The Server and
2. The Client
### MongoDB: The Server
The official MongoDB package includes a configuration file `/etc/mongod.conf`. These settings (such as the data directory and log directory specifications) take effect upon startup. That is, if you change the configuration file while the MongoDB instance is running, you must restart the instance for the changes to take effect.

Go ahead and update this `mongod.conf` file and be sure the `dbPath` under `storage:` and the `path` under the `systemLog:` sections are pointing to the correct data and log directories previously mentioned [here](https://github.com/cibervicho/go_webApp/blob/master/README.md#mongodb-the-server).

Another change to be made in this configuration file is under the `net:` section. The `bindIp` parameter should be `0.0.0.0` instead as `127.0.0.1`. This to allow the database to be accesible from different clients, not only from localhost.

The following is an example of the `mongod.conf` file:
   ```sh
   # mongod.conf
   
   # for documentation of all options, see:
   #   http://docs.mongodb.org/manual/reference/configuration-options/
   
   # Where and how to store data.
   storage:
     dbPath: /var/lib/mongodb
     journal:
       enabled: true
   #  engine:
   #  mmapv1:
   #  wiredTiger:
   
   # where to write logging data.
   systemLog:
     destination: file
     logAppend: true
     path: /var/log/mongodb/mongod.log
   
   # network interfaces
   net:
     port: 27017
     bindIp: 0.0.0.0 # Default was 127.0.0.1
   
   
   # how the process runs
   processManagement:
     timeZoneInfo: /usr/share/zoneinfo
   
   #security:
   
   #operationProfiling:
   
   #replication:
   
   #sharding:
   
   ## Enterprise-Only Options:
   
   #auditLog:
   
   #snmp:
   ```
#### Launch MongoDB as a service
We need to create a unit file, which tells systemd how to manage a resource. Most common unit type, service, determine how to start or stop the service, auto-start etc.

Create a configuration file named mongodb.service in /etc/systemd/system to manage the MongoDB service.
```sh
~$ sudo vim /etc/systemd/system/mongodb.service
```
And copy the following contents in this file:
```sh
#Unit contains the dependencies to be satisfied before the service is started.
[Unit]
Description=MongoDB Database
After=network.target
Documentation=https://docs.mongodb.org/manual
# Service tells systemd, how the service should be started.
# Key `User` specifies that the server will run under the mongodb user and
# `ExecStart` defines the startup command for MongoDB server.
[Service]
User=mongodb
Group=mongodb
ExecStart=/usr/bin/mongod --quiet --config /etc/mongod.conf
# Install tells systemd when the service should be automatically started.
# `multi-user.target` means the server will be automatically started during boot.
[Install]
WantedBy=multi-user.target
```
**Note that the `ExecStart` parameter should point to the `mongod.conf` file you created in the above.**

Update the systemd service with the command stated below:
```sh
~$ systemctl daemon-reload
```
Start the service with systemcl:
```sh
~$ sudo systemctl start mongodb
```
Check if mongodb has been started on port `27017` with netstat command:
```sh
~$ netstat -plntu
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.53:53           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      -                   
tcp        0      0 127.0.0.1:631           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:27017           0.0.0.0:*               LISTEN      -                   
tcp6       0      0 :::8080                 :::*                    LISTEN      -                   
tcp6       0      0 :::22                   :::*                    LISTEN      -                   
tcp6       0      0 ::1:631                 :::*                    LISTEN      -                   
udp        0      0 127.0.0.53:53           0.0.0.0:*                           -                   
udp        0      0 0.0.0.0:68              0.0.0.0:*                           -                   
udp        0      0 0.0.0.0:631             0.0.0.0:*                           -                   
udp        0      0 0.0.0.0:52086           0.0.0.0:*                           -                   
udp        0      0 0.0.0.0:5353            0.0.0.0:*                           -                   
udp6       0      0 :::33848                :::*                                -                   
udp6       0      0 :::5353                 :::*                                -                   
udp6       0      0 :::5353                 :::*                                -                   
udp6       0      0 :::46629                :::*                                -                   
udp6       0      0 :::6666                 :::*                                2377/qlipper
```
Check if the service has started properly:
```sh
~$ sudo systemctl status mongodb
● mongodb.service - MongoDB Database
   Loaded: loaded (/etc/systemd/system/mongodb.service; enabled; vendor preset: enabled)
   Active: active (running) since Wed 2019-05-15 21:39:38 CDT; 3min 24s ago
     Docs: https://docs.mongodb.org/manual
 Main PID: 795 (mongod)
    Tasks: 27 (limit: 4915)
   Memory: 180.8M
   CGroup: /system.slice/mongodb.service
           └─795 /usr/bin/mongod --quiet --config /etc/mongod.conf

May 15 21:39:38 server systemd[1]: Started MongoDB Database.
```
The output to the above command will show `active (running)` status with the PID and Memory/CPU it is consuming.

Enable auto start MongoDB when system starts.
```sh
~$ sudo systemctl enable mongodb
```
#### To stop MongoDB
```sh
~$ sudo systemctl stop mongodb
```
#### To restart MongoDB
```sh
~$ sudo systemctl restart mongodb
```

### MongoDB: The Client
Now that MongoDB service is up and running connect to it in a terminal:
```sh
~$ mongo
MongoDB shell version v4.0.9
connecting to: mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("f768f4c0-eab6-4728-82da-6b5ced859698") }
MongoDB server version: 4.0.9
Server has startup warnings: 
2019-05-15T21:39:57.541-0500 I STORAGE  [initandlisten] 
2019-05-15T21:39:57.541-0500 I STORAGE  [initandlisten] ** WARNING: Using the XFS filesystem is strongly recommended with the WiredTiger storage engine
2019-05-15T21:39:57.541-0500 I STORAGE  [initandlisten] **          See http://dochub.mongodb.org/core/prodnotes-filesystem
2019-05-15T21:40:02.403-0500 I CONTROL  [initandlisten] 
2019-05-15T21:40:02.403-0500 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2019-05-15T21:40:02.403-0500 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
2019-05-15T21:40:02.403-0500 I CONTROL  [initandlisten] 
---
Enable MongoDB's free cloud-based monitoring service, which will then receive and display
metrics about your deployment (disk utilization, CPU, operation statistics, etc).

The monitoring data will be available on a MongoDB website with a unique URL accessible to you
and anyone you share the URL with. MongoDB may use this information to make product
improvements and to suggest MongoDB products and deployment options to you.

To enable free monitoring, run the following command: db.enableFreeMonitoring()
To permanently disable this reminder, run the following command: db.disableFreeMonitoring()
---

>
```
Now that we are under the MongoDB shell, lets create our `moviesdb` and populate it with some data we'll be using:
```sh
> use moviesdb
switched to db moviesdb
> 
```
Copy and paste the following instructions, just as they are to insert into the `moviesdb`, using the collection `imdb`:
```sh
db.imdb.insert({title: 'Ant-Man and the Wasp',
year: '2018',
rated: 'PG-13',
genre: ['Action', 'Adventure', 'Comedy'],
plot: 'As Scott Lang balances being both a Super Hero and a father, Hope van Dyne and Dr. Hank Pym present an urgent new mission that finds the Ant-Man fighting alongside The Wasp to uncover secrets from their past.'});

db.imdb.insert({title: 'Independence Day',
year: '1996',
rated: 'PG-13',
genre: ['Action', 'Adventure', 'Sci-Fi'],
plot: "The aliens are coming and their goal is to invade and destroy Earth. Fighting superior technology, mankind's best weapon is the will to survive."});

db.imdb.insert({title: 'How to Be a Latin Lover',
year: '2017',
rated: 'PG-13',
genre: ['Comedy', 'Drama'],
plot: 'Finding himself dumped after 25 years of marriage, a man who made a career of seducing rich older women must move in with his estranged sister, where he begins to learn the value of family.'});

db.imdb.insert({title: 'A Lego Brickumentary',
year: '2014',
rated: 'G',
genre: ['Documentary'],
plot: 'A look at the global culture and appeal of the LEGO building-block toys.'});

db.imdb.insert({title: 'The Imitation Game',
year: '2014',
rated: 'PG-13',
genre: ['Biography', 'Drama', 'Thriller'],
plot: 'During World War II, the English mathematical genius Alan Turing tries to crack the German Enigma code with help from fellow mathematicians.'});

db.imdb.insert({title: 'Y Tu Mama Tambien',
year: '2001',
rated: 'R',
genre: ['Drama'],
plot: 'In Mexico, two teenage boys and an attractive older woman embark on a road trip and learn a thing or two about life, friendship, sex, and each other.'});

db.imdb.insert({title: 'School of Rock',
year: '2003',
rated: 'PG-13',
genre: ['Comedy', 'Music'],
plot: 'After being kicked out of his rock band, Dewey Finn becomes a substitute teacher of an uptight elementary private school, only to try and turn them into a rock band.'});

db.imdb.insert({title: 'The Jungle Book',
year: '2016',
rated: 'PG',
genre: ['Adventure', 'Drama', 'Family'],
plot: 'After a threat from the tiger Shere Khan forces him to flee the jungle, a man-cub named Mowgli embarks on a journey of self discovery with the help of panther Bagheera and free-spirited bear Baloo.'});

db.imdb.insert({title: 'The Girl with All the Gifts',
year: '2016',
rated: 'R',
genre: ['Drama', 'Horror', 'Sci-Fi'],
plot: 'A scientist and a teacher living in a dystopian future embark on a journey of survival with a special young girl named Melanie.'});

db.imdb.insert({title: 'The Peanuts Movie',
year: '2015',
rated: 'G',
genre: ['Animation', 'Comedy', 'Drama'],
plot: 'Snoopy embarks upon his greatest mission as he and his team take to the skies to pursue their archnemesis, while his best pal Charlie Brown begins his own epic quest back home to win the love of his life.'});

db.imdb.insert({title: 'Zootopia',
year: '2016',
rated: 'PG',
genre: ['Animation', 'Adventure', 'Comedy'],
plot: 'In a city of anthropomorphic animals, a rookie bunny cop and a cynical con artist fox must work together to uncover a conspiracy.'});

db.imdb.insert({title: 'La La Land',
year: '2016',
rated: 'PG-13',
genre: ['Comedy', 'Drama', 'Music'],
plot: 'While navigating their careers in Los Angeles, a pianist and an actress fall in love while attempting to reconcile their aspirations for the future.'});

db.imdb.insert({title: 'Memento',
year: '2000',
rated: 'R',
genre: ['Mystery', 'Thriller'],
plot: "A man with short-term memory loss attempts to track down his wife's murderer."});

db.imdb.insert({title: 'The Ottoman Lieutenant',
year: '2017',
rated: 'R',
genre: ['Drama', 'War'],
plot: 'This movie is a love story between an idealistic American nurse and a Turkish officer in World War I.'});

db.imdb.insert({title: 'Quantum Leap',
year: '1989',
rated: 'TV-PG',
genre: ['Action', 'Adventure', 'Drama'],
plot: 'Scientist Sam Beckett finds himself trapped in the past, "leaping" into the bodies of different people on a regular basis.'});

db.imdb.insert({title: 'Dinosaurs',
year: '1991',
rated: 'TV-PG',
genre: ['Comedy', 'Family', 'Fantasy'],
plot: 'Dinosaurs follows the life of a family of dinosaurs, living in a modern world. They have TVs, fridges, etc. The only humans around are cavemen, who are viewed as pets and wild animals.'});

db.imdb.insert({title: 'Ready Player One',
year: '2018',
rated: 'PG-13',
genre: ['Action', 'Adventure', 'Sci-Fi'],
plot: 'When the creator of a virtual reality world called the OASIS dies, he releases a video in which he challenges all OASIS users to find his Easter Egg, which will give the finder his fortune.'});

db.imdb.insert({title: 'Pinocchio',
year: '1940',
rated: 'G',
genre: ['Animation', 'Comedy', 'Family'],
plot: 'A living puppet, with the help of a cricket as his conscience, must prove himself worthy to become a real boy.'});
```
To verify the collection `imdb` has data just execute the following command:
```sh
> db.imdb.find().count()
18

> db.imdb.find()
{ "_id" : ObjectId("5cc54f43d58a696e3bf86eea"), "title" : "Ant-Man and the Wasp", "year" : "2018", "rated" : "PG-13", "genre" : [ "Action", "Adventure", "Comedy" ], "plot" : "As Scott Lang balances being both a Super Hero and a father, Hope van Dyne and Dr. Hank Pym present an urgent new mission that finds the Ant-Man fighting alongside The Wasp to uncover secrets from their past." }
{ "_id" : ObjectId("5cc54f67d58a696e3bf86eeb"), "title" : "Independence Day", "year" : "1996", "rated" : "PG-13", "genre" : [ "Action", "Adventure", "Sci-Fi" ], "plot" : "The aliens are coming and their goal is to invade and destroy Earth. Fighting superior technology, mankind's best weapon is the will to survive." }
{ "_id" : ObjectId("5cc54f71d58a696e3bf86eec"), "title" : "How to Be a Latin Lover", "year" : "2017", "rated" : "PG-13", "genre" : [ "Comedy", "Drama" ], "plot" : "Finding himself dumped after 25 years of marriage, a man who made a career of seducing rich older women must move in with his estranged sister, where he begins to learn the value of family." }
{ "_id" : ObjectId("5cc54f7bd58a696e3bf86eed"), "title" : "A Lego Brickumentary", "year" : "2014", "rated" : "G", "genre" : [ "Documentary" ], "plot" : "A look at the global culture and appeal of the LEGO building-block toys." }
{ "_id" : ObjectId("5cc54f84d58a696e3bf86eee"), "title" : "The Imitation Game", "year" : "2014", "rated" : "PG-13", "genre" : [ "Biography", "Drama", "Thriller" ], "plot" : "During World War II, the English mathematical genius Alan Turing tries to crack the German Enigma code with help from fellow mathematicians." }
{ "_id" : ObjectId("5cc54f8dd58a696e3bf86eef"), "title" : "Y Tu Mama Tambien", "year" : "2001", "rated" : "R", "genre" : [ "Drama" ], "plot" : "In Mexico, two teenage boys and an attractive older woman embark on a road trip and learn a thing or two about life, friendship, sex, and each other." }
{ "_id" : ObjectId("5cc54f94d58a696e3bf86ef0"), "title" : "School of Rock", "year" : "2003", "rated" : "PG-13", "genre" : [ "Comedy", "Music" ], "plot" : "After being kicked out of his rock band, Dewey Finn becomes a substitute teacher of an uptight elementary private school, only to try and turn them into a rock band." }
{ "_id" : ObjectId("5cc54f9dd58a696e3bf86ef1"), "title" : "The Jungle Book", "year" : "2016", "rated" : "PG", "genre" : [ "Adventure", "Drama", "Family" ], "plot" : "After a threat from the tiger Shere Khan forces him to flee the jungle, a man-cub named Mowgli embarks on a journey of self discovery with the help of panther Bagheera and free-spirited bear Baloo." }
{ "_id" : ObjectId("5cc54fa5d58a696e3bf86ef2"), "title" : "The Girl with All the Gifts", "year" : "2016", "rated" : "R", "genre" : [ "Drama", "Horror", "Sci-Fi" ], "plot" : "A scientist and a teacher living in a dystopian future embark on a journey of survival with a special young girl named Melanie." }
{ "_id" : ObjectId("5cc54facd58a696e3bf86ef3"), "title" : "The Peanuts Movie", "year" : "2015", "rated" : "G", "genre" : [ "Animation", "Comedy", "Drama" ], "plot" : "Snoopy embarks upon his greatest mission as he and his team take to the skies to pursue their archnemesis, while his best pal Charlie Brown begins his own epic quest back home to win the love of his life." }
{ "_id" : ObjectId("5cc54fb3d58a696e3bf86ef4"), "title" : "Zootopia", "year" : "2016", "rated" : "PG", "genre" : [ "Animation", "Adventure", "Comedy" ], "plot" : "In a city of anthropomorphic animals, a rookie bunny cop and a cynical con artist fox must work together to uncover a conspiracy." }
{ "_id" : ObjectId("5cc54fbcd58a696e3bf86ef5"), "title" : "La La Land", "year" : "2016", "rated" : "PG-13", "genre" : [ "Comedy", "Drama", "Music" ], "plot" : "While navigating their careers in Los Angeles, a pianist and an actress fall in love while attempting to reconcile their aspirations for the future." }
{ "_id" : ObjectId("5cc54fc5d58a696e3bf86ef6"), "title" : "Memento", "year" : "2000", "rated" : "R", "genre" : [ "Mystery", "Thriller" ], "plot" : "A man with short-term memory loss attempts to track down his wife's murderer." }
{ "_id" : ObjectId("5cc54fcdd58a696e3bf86ef7"), "title" : "The Ottoman Lieutenant", "year" : "2017", "rated" : "R", "genre" : [ "Drama", "War" ], "plot" : "This movie is a love story between an idealistic American nurse and a Turkish officer in World War I." }
{ "_id" : ObjectId("5cc54fd4d58a696e3bf86ef8"), "title" : "Quantum Leap", "year" : "1989", "rated" : "TV-PG", "genre" : [ "Action", "Adventure", "Drama" ], "plot" : "Scientist Sam Beckett finds himself trapped in the past, \"leaping\" into the bodies of different people on a regular basis." }
{ "_id" : ObjectId("5cc54fdad58a696e3bf86ef9"), "title" : "Dinosaurs", "year" : "1991", "rated" : "TV-PG", "genre" : [ "Comedy", "Family", "Fantasy" ], "plot" : "Dinosaurs follows the life of a family of dinosaurs, living in a modern world. They have TVs, fridges, etc. The only humans around are cavemen, who are viewed as pets and wild animals." }
{ "_id" : ObjectId("5cc54fe1d58a696e3bf86efa"), "title" : "Ready Player One", "year" : "2018", "rated" : "PG-13", "genre" : [ "Action", "Adventure", "Sci-Fi" ], "plot" : "When the creator of a virtual reality world called the OASIS dies, he releases a video in which he challenges all OASIS users to find his Easter Egg, which will give the finder his fortune." }
{ "_id" : ObjectId("5cc54feed58a696e3bf86efb"), "title" : "Pinocchio", "year" : "1940", "rated" : "G", "genre" : [ "Animation", "Comedy", "Family" ], "plot" : "A living puppet, with the help of a cricket as his conscience, must prove himself worthy to become a real boy." }
>
```
If you get `18` as the number of entries in the collection `imdb` after typing the command `db.imdb.find().count()`, and then you see the data displayed with the command `db.imdb.find()` you are good to **Go**

## Installing Jenkins and the required plugins
### Explaining the bash scripts in Jenkins
