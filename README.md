# Session Initiation Protocol (SIP) server in golang

## Building
$ cd $GOPATH
$ mkdir -p src/github.com/uname # uname(username) will be changed later
$ git clone https://kalpacg@bitbucket.org/teamsip/sip-server.git src/github.com/uname
# It is necessary to build siputility library before using sip-server
$ go install github.com/uname/sip-server
$ ./bin/sip-server
 
