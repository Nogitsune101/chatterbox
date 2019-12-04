# Intro #
So I have determined that the chat app is an evil test.  There were so many places that I could have gone that I eventually ended up having to lock my scope and set a deadline for the project to prevent myself from continuing to keep adding features.  That being said, I was finally able to determine what I would be happy with as a mvc (minimum viable product) and concentrated on getting it done.  I'm so pleased with  the final product that I think I'm going to keep going until it is what I initially planned to be which is a portable web chat that uses end to end AES encryption

# Planning #
During the planning phase I took the initial requirements and determined that the finished product would be a golang based backend that uses websockets and could be compiled into a secure portable chat server.  I should note that while the plan is to ultimately finish this application.  Below is a breakdown on the project milestones and the application you are getting now:

- Milestone 1 - Meet the requirements of the test
- Milestone 2 - Add direct messaging / chat features (currently what I'm turning in) 
- Milestone 3 - Add authentication / security
- Milestone 4 - Add component that takes server assets and converts them to golang resources so they can be compiled directly into the executable making the app portable
- Release - Testing / Bug fixes / refactoring / polish

I decided to use golang / pure javascript (ES6) as my base for several reasons:
- golang is so strict that it is less prone to bugs
- I've been mostly working with nodejs/react for front end design, but wanted to show I could actually understand/write raw javascript and know my html/css p's and q's
- I haven't written a golang server in several months so a good refresher
- I wanted to use as little dependencies as possible, once again, to show I can code

Here are the technologies I've used:
- golang
- gorilla mux (http muxer for golang)
- gorilla websockets (websockets for golang)
- cobra command (cli framework for golang)

IMPORTANT NOTE - I used https://github.com/gorilla/websocket/tree/master/examples/chat as my base for the sockets portion of the application. Most of this code has been touched or replaced to add features

# Installation / Building #
This is pretty much a golang application, all you really need to compile it is golang. I used go version 1.12 on Ubuntu to compile my application.  As for compatibility, I have not tested on OSX or Windows but my experience with go compatibility has been stellar cross platform.  I am not sure however if it will compile with newer versions of go, I have seen backward compatibility issues with v1.13+

Go to https://golang.org/dl/ and grab golang v1.12.12 and install it following the site instructions. (A personal note - make sure you pay attention to the path environmental settings, I've found this to be a sticking point for most people)

Once the correct version of golang is installed, you should be able to clone the repository using git and run of the following to start the server

## Git Cloning ##
- `git clone https://github.com/Nogitsune101/chatterbox.git`

## Development Run ##
 - `go run main.go`

## Build & Run ##
 - `go build`
 - `./chatterbox` (*nix systems)
 - IMPORTANT NOTE - The app is not yer portable, so if you deploy this, you will also need the assets directory

# Enviromentals / Execution #

## Environmentals ##
While the system supports .env and environmentals (typically used in deployments), sadly none were used in this project

## CLI Usage ##
```
Usage:
  chatterbox [flags]

Flags:
  -h, --help            help for chatterbox
  -a, --ipaddr string   IP Address of the webserver (default: 0.0.0.0) (default "0.0.0.0")
  -p, --port string     Port of the webserver (default: 8080) (default "8080")
      --version         version for chatterbox
```

# Testing / Compatibility #
Had I the time, I would have added cypress tests for integration testing on the client side. It supports websockets and think it would be more valuable in this case than unit tests on the server side.  Really the server is a glorified http server and websocket relay system.  I didn't use interface types which is where you run into runtime errors with go.  Since it is comprised mostly of channels, simple types, and standard modeling, the compiler will catch any errors that occur.  Mostly, I relied on manual testing.

## Compile/Run ##
- Ubuntu (works)
- OSX (untested)
- Windows 10 (untested)

## Web Client ##
- Chrome (compatible)
- Firefox (compatible)
- Edge Beta (compatible)
- Edge (nope)
- Safari (unknown)

# Client Usage #
```
/e or /emote {emote text}         (Sends a custom emote)
/h or /help                       (Displays help text)
/r or /reply {message}            (Replies to the last whisper you received)
/s or /shout {message}            (Shouts a message to the others in the room)
/w or /whisper {user} {message}   (Sends a private message to another user)
/fart                             (Fart emote)
/hello                            (Hello emote)
/wave                             (Wave emote)
/do a barrel roll                 (The original!)
/look around                      (The mailbox is a lie!)
```

# Current Features #
- User Sessions
- Chat
- Private Chat
- Custom Emotes
- Static Emotes
- Easter Eggs
- Command Support

# Planned Features # 
- AES end to end message encryption
- Multiroom support
- Room user view
- Chat emoji support
- Chat links support
- Chat file support
- Make the application fully portable

# Known Bugs #
- [Firefox] easter egg 1 is a bit clippy
- [Firefox] easter egg 2, css background-size: cover doesn't appear to be supported
- [Client] people can use the same display name with weird results!
