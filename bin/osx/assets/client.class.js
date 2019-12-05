    /**
     * ChatClient
     * 
     * Single class that handles communications to chatterbox websocket server
     * 
     * NOTE - I prefer to create classes for complex components, they are easier
     * to decouple and use in other places
     * 
     */
    class ChatterBoxClient {
        userName = null         // This is display name for the user
        roomName = null         // This is the room we are currently in
    
        users = []              // TODO - Users List not yet supported
        rooms = []              // TODO - Rooms List not yet supported

        targetMessageElm = null    // Target element to send messages to

        _connection = null      // In a perfect world, this would be private, underscore is just a visual queue
        _message = new ChatterBoxMessage

        // Initialize a new connection
        init = (username, roomname, targetelm) => {
            this.userName = username
            this.roomName = roomname
            this.targetMessageElm = targetelm

            return this._connect()
        }

        // Sends a new message to the web socket
        send = (message) => {

            // Generate our response to the server
            const payload = this._message.command(message).toString()

            // If there is a resulting payload
            if(payload) {

                // And this it not a SELF or INFO message
                if(payload.charAt(0) !== "0" && payload.charAt(0) !== "4") {

                    // Send it to the websocket
                    this._connection.send(payload)
                
                // Otherwise send it direct to the chat view
                } else {
                    this._handleMessage(payload)
                }
            }
        }
        
        // Used to see if a connection is active - for future use
        isConnected = () => {
            return (this._connection !== null ? true : false)
        }

        // Initializes a new web socket connection to the server
        _connect = () => {
            // TODO - username in the url? we need auth
            this._connection = new WebSocket("ws://" + document.location.host + "/room/" + this.roomName + "/" + this.userName)

            // Display our help message first
            this._handleMessage(this._message.info("Enter /h for help").toString())

            // Pass the user to the message object to give it context
            this._message.to = this.userName

            // Register connection events

            // When closing the socket
            this._connection.onclose = (event) => {
                this._handleMessage(this._message.error(this._message._separator + "You have been disconnected from the server").toString())
            }

            // When recieving a new message
            this._connection.onmessage = (event) => {
                this._handleMessage(event.data)
            }
        }

        // handles the displaying of the message element to a target in the html
        _handleMessage = (message) => {
            const targetView = this.targetMessageElm
            var doScroll = targetView.scrollTop > targetView.scrollHeight - targetView.clientHeight - 1
            targetView.appendChild(this._message.parse(message).toElement())
            if (doScroll) {
                targetView.scrollTop = targetView.scrollHeight - targetView.clientHeight
            }
        }
    }