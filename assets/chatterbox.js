window.onload = () => {

    // Sanity Check! If we don't support websockets no need to go to any further
    if (!window["WebSocket"]) {
        document.getElementsByTagName('body')[0].innerHTML = `
    <div align='center'>
        <h2>YOU SHALL NOT PASS!</h2>
        <h4>(Because your web browser doesn't support web sockets)</h4>
    </div>`
        throw new Error("Browser doesn't support websockets, you should probably upgrade . . . ")
    }
    
    // Makes things a bit more readable
    const entryForm = document.getElementById("entry_form")
    const registrationForm = document.getElementById("registration_form")
    const messageInput = document.getElementById("message")
    const displayNameInput = document.getElementById("display_name")
    const modal = document.getElementById("modal")
    const messageView = document.getElementById("chat_view")

    // Chatterbox
    const chatterbox = new ChatterBoxClient
    
    // Handle send message event
    entryForm.onsubmit = (event) => {
        if(messageInput.value) {
            chatterbox.send(messageInput.value)
            messageInput.value = ""
            messageInput.focus()
        }
        event.preventDefault()
    }

    // Handle display name event
    registrationForm.onsubmit = (event) => {
        if(displayNameInput.value) {
            chatterbox.init(displayNameInput.value, "Nexus", messageView)
            modal.style.display = "none"
            messageInput.focus()
        }
        event.preventDefault()
    }

    // One and done
    displayNameInput.focus()   // Focus on the display name input
}
