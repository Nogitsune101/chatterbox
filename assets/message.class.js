
// Message type decoder ring
const CBMTYPES = {
    SELF: "0",
    ERROR: "1",
    WARNING: "2",
    NOTICE: "3",
    INFO: "4",
    WHISPER: "5",
    EMOTE: "6",
    SAY: "7",
}

/**
 * Chatterbox Message Class
 * 
 * This component handles reading and writing individual ws messages from the server and the client
 *  
 */
class ChatterBoxMessage {
    // Params (Initialized)
    type = (CBMTYPES.SAY)
    to = ""
    from = ""
    message = ""

    // I figure that using character code 0 is safe to use as a separator character since it is not keyboard friendly
    _separator = String.fromCharCode(0)
    _who = "■" // Alt + 254 (old skool)
    _replyto = ""

    // NOTE - Used method chaining style for these

    // Sets object to be an error
    error = (message) => {
        this.type = CBMTYPES.ERROR
        this.message = message
        return this
    }

    // Sets object to be a warning
    warning = (message) => {
        this.type = CBMTYPES.WARNING
        this.message = message
        return this
    }

    // Sets object to be a notice
    notice = (message) => {
        this.type = CBMTYPES.NOTICE
        this.message = message
        return this
    }

    // Sets object to be a notice (This one isn't broadcast which is handled in the client)
    info = (message) => {
        this.type = CBMTYPES.INFO
        this.message = this._separator+message
        return this
    }

    // Sets object to send private messages
    whisper = (message) => {
        this.type = CBMTYPES.WHISPER

        const messageParts = message.match(/^[\s]?([^\s]*)[\s]+(.*)$/i)
        // TODO - Something seems wrong here
        this.message = messageParts[1] + this._separator + "[To &#91;" + messageParts[1] + "&#93;|&#91;" + this.to + "&#93; whispers]: " + messageParts[2]

        return this
    }

    // Sets object to be an emote
    emote = (message) => {
        this.type = CBMTYPES.EMOTE
        this.message = this._who + " " + message
        return this
    }

    // Default chat method
    say = (message) => {
        this.type = CBMTYPES.SAY
        this.message = this._who + " say[|s]: " + message
        return this
    }

    ee1 = () => {
        document.body.classList.add("ee1")
        this.message = this._who + " [|is spinning around in circles, seemingly amused]"
        this.type = CBMTYPES.EMOTE
        return this
    }

    ee2 = () => {
        document.getElementsByClassName("content")[0].classList.add("ee2")
        this.message = this._separator+"You are standing in an open field west of a white house, with a boarded front door. There is a small mailbox here."
        this.type = CBMTYPES.SELF
        return this
    }

    // Commands are a chat sub system for performing predefined messages
    command = (message) => {
        // Parse our command and seperate it from the payload
        const matches = message.match(/^[\/]([^\s]*)[\s]?/i)
        const command = (matches ? matches[1] : null)
        const getPayloadRegExp = new RegExp("^[\/]"+command+"[\s]?")
        const payload = message.replace(getPayloadRegExp, "")

        // Fun stuff goes here (handles/evaluate commands)
        switch(command) {
            case 'do':
                switch(payload.toLowerCase()) {
                    case " a barrel roll":
                        return this.ee1()
                        break
                }
                break
            case 'e':
            case 'emote':
                return this.emote(payload)
                break
            case 'fart':
                return this.emote("[fart loudly|farts in your general direction]")
                break
            case 'hello':
                return this.emote("[greet everyone magnanimously!|greets everyone with a friendly hello!]")
                break
            case 'h':
            case 'help':
                return this.info(
`/e or /emote {emote text}         (Sends a custom emote)
/h or /help                       (Displays help text)
/r or /reply {message}            (Replies to the last whisper you received)
/s or /shout {message}            (Shouts a message to the others in the room)
/w or /whisper {user} {message}   (Sends a private message to another user)
/fart                             (Fart emote)
/hello                            (Hello emote)
/wave                             (Wave emote)
/do a barrel roll                 (The original!)
/look around                      (The mailbox is a lie!)`)

                break
            case 'look':
                switch(payload.toLowerCase()) {
                    case " around":
                        return this.ee2()
                        break
                }
                break
            case 'r':
            case 'reply':
                if(!this._replyto) {
                    return this.info("Nobody has talked to you yet . . . so lonely")
                } else {
                    return this.whisper(" " + this._replyto + " " + payload)
                }
                break
            case 's':
            case 'shout':
                    return this.error(this._who + " shout[|s]: " + payload)
                    break
            case 'wave':
                return this.emote("[wave hello|greets everyone with a friendly wave]")
                break
            case 'w':
            case 'whisper':
                return this.whisper(payload)
                break
            default:
                return this.say(payload)
        }
    }

    // This is used to parse incoming messages
    parse = (message) => {
        const messageParts = message.split(this._separator)
        // TODO - The more I look at this, it would be better to make the entire string format 4 elements instead of 3 or 4 elements, would be less fragile
        if(messageParts.length >= 3) {
            this.type = messageParts[0]
            this.from = messageParts[1]
            this.message = messageParts[messageParts.length-1]

            if(this.type == CBMTYPES.WHISPER && this.from.toLowerCase() != this.to.toLowerCase()) {
                this._replyto = messageParts[1]
            }
        } else {
            this.type = CBMTYPES.ERROR
            this.message = "Unable to parse message [" + message + "]"
        }
        return this
    }

    // This is used to format the message for broadcast
    toString = () => {
        return (this.message ? this.type + this._separator + this.message : '')
    }

    // This is used to change the message into a html element
    toElement = () => {
        const typeKeys = Object.keys(CBMTYPES)
        const className = typeKeys[(this.type)].toLowerCase()

        if(this.from.toLowerCase() == this.to.toLowerCase()) {
            this.message = this.message
                .replace(/\[([^|]*)[|]([^\]]*)\]/, "$1")
                .replace(this._who, "You")
                .replace("\n", "<br />", -1)
        } else {
            this.message = this.message
                .replace(/\[([^|]*)[|]([^\]]*)\]/, "$2")
                .replace(this._who, "[" + this.from + "]")
                .replace(/[\n\r]+/g, "<br />", -1)
        }

        const elm = document.createElement("div")
            elm.classList.add(className)
            elm.innerHTML = this.message
        
        return elm
    }
}