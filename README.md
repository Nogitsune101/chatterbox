# Intro #
So I have determined that the chat app is an evil test.  There were so many places that I could have gone that I eventually ended up having to lock my scope and set a deadline for the project to prevent myself from continuing to keep adding features.  That being said, I was finally able to determine what I would be happy with as a mvc (minimum viable product) and concentrated on getting it done.  I'm so pleased with  the final product that I think I'm going to keep going until it is what I initially planned to be which is a portable web chat that uses end to end AES encryption

# Planning #
During the planning phase I took the initial requirements and determined that the finished product would be a golang based backend that uses websockets and could be compiled into a secure portable chat server.  I should note that while the plan is to ultimately finish this application.  Below is a breakdown on the project milestones and the application you are getting now:

- Milestone 1 - Meet the requirements of the test (complete)
- Milestone 2 - Add direct messaging / chat features (complete) 
- Milestone 3 - Add authentication / security (pending)
- Milestone 4 - Add component that takes server assets and converts them to golang resources so they can be compiled directly into the executable making the app portable (complete)
- Release - Testing / Bug fixes / refactoring / polish (pending)

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
This is pretty much a golang application, all you really need to compile it is golang. I used go version 1.12 on dev Ubuntu machine to compile my application. I have at this point confirmed that the application compiles fine with the latest version of golang on OSX (x64), Win10 (x64), and Ubuntu (x64)

If you have issues installing and compiling with golang (the struggle is real), I have provided binaries with the respository.  I'm also including a step-by-step for OSX.

## OSX Installation ##
- Use your favorite browser and navigate to https://golang.org/dl
- Download and install go1.13.5.darwin-amd64.pkg
- Open a new terminal (do not use a previously opened terminal)
- Type `mkdir $HOME/go`
- Type `cd $HOME/go`
- Type `git clone https://github.com/Nogitsune101/chatterbox.git`
- Type `cd chatterbox`
- Type `go run main.go`
- At this point go should download packages and start the chat server
- Use one of the compatible browsers listed below and open a couple windows for testing
- Enter a unique username in each window and have fun

## Git Cloning ##
- `git clone https://github.com/Nogitsune101/chatterbox.git`

## Development Run ##
 - `go run main.go`

## Build & Run ##
 - `go run main.go build`
 - `./chatterbox` (*nix systems)

## Portable Build Note ##
This app now uses a precompiler to convert client assets to go code before compiling the app.  When using the development run,
a build option is available that automatically generates the embeded client assets and compiles it with the application to
make it fully portable. Another quick note, when running the server in development mode (go run main.go) it will use the assets
directory instead of the embeded assets.

## Client Access ##
- Webserver will start by default at http://localhost:8080

# Enviromentals / Execution #

## Environmentals ##
While the system supports .env and environmentals (typically used in deployments), sadly none were used in this project

## CLI Usage ##
```
Usage:
  chatterbox [flags]

Flags:
  -h, --help            help for chatterbox
  -a, --ipaddr string   IP Address of the webserver (default: 0.0.0.0) 
  -p, --port string     Port of the webserver (default: 8080)
      --version         version for chatterbox
```

# Testing / Compatibility #
Had I the time, I would have added cypress tests for integration testing on the client side. It supports websockets and think it would be more valuable in this case than unit tests on the server side.  Really the server is a glorified http server and websocket relay system.  I didn't use interface types which is where you run into runtime errors with go.  Since it is comprised mostly of channels, simple types, and standard modeling, the compiler will catch any errors that occur.  Mostly, I relied on manual testing.

## Compile/Run ##
- Ubuntu (works)
- OSX (works)
- Windows 10 (works)

## Web Client ##
- Chrome (compatible) *recommended
- Edge Beta (compatible)
- Firefox (compatible) *minor css issues
- Safari (nope) *notes in known bugs
- Edge (nope) *notes in known bugs

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
- Fully Portable

# Planned Features # 
- AES end to end message encryption
- Multiroom support
- Room user view
- Chat emoji support
- Chat links support
- Chat file support

# Known Bugs #
- [Firefox] easter egg 1 is a bit clippy
- [Firefox] easter egg 2, css background-size: cover doesn't appear to be supported
- [Client] people can use the same display name with weird results!
- [Edge] issues due to Edge's lack of full ES6 support
- [Safari] issues due to Safari's lack of full ES6 support
