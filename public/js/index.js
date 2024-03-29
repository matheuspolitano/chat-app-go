var socket;
var messageInputField = document.getElementById("message-user");
var username;
document.getElementById("chatForm").onsubmit = () => {
    if (!socket) {
        console.log("WebSocket connection not established");
        return false;
    }
    if (!messageInputField.value) {
        console.log("No message to send");
        return false;
    }
    console.log(messageInputField.value);
    socket.send(messageInputField.value);
     // Clear the input field after sending the message


    const messageElement = document.createElement("div");
    messageElement.className = "flex items-start mb-4 text-sm";
    messageElement.innerHTML = `
        <div class="flex-1 p-2 rounded bg-blue-100 text-right">${messageInputField.value}</div>
    `;

    // Append the message element to the chatMessage element
    const chatMessageElement = document.getElementById("chatMessage");
    if (chatMessageElement) {
        chatMessageElement.appendChild(messageElement);
    }

    chatMessageElement.scrollTop = chatMessageElement.scrollHeight;
    messageInputField.value = '';
    return false; // Prevents form from submitting in the traditional way
};

document.getElementById('userForm').addEventListener('submit', function(event) {
    event.preventDefault();
    username = document.getElementById('username').value;
    if (username) {
        // Store the username or use it as needed
        // Hide the modal and show the main content
        document.getElementById('userModal').style.display = 'none';
        document.getElementById('mainContent').style.display = 'block';
        var userIdDiv = document.getElementById("UserId");
        var userImage = document.getElementById("UserImage");

// Change the src attribute
userImage.src = "https://github.com/"+username+".png"; // Replace with the new image URL

// Change the alt attribute
userImage.alt = username; 
        userIdDiv.innerText = username;
        load_socket(username)
    }
});








function load_socket(username){
    if(window["WebSocket"]){
        socket = new WebSocket("ws://" + document.location.host + "/ws?username="+username);

        socket.onmessage = function(event) {
            console.log(event)
            const data = JSON.parse(event.data); // The message received from the server
        
            // Create the message HTML element
            const messageElement = document.createElement("div");
            messageElement.className = "flex items-start mb-4 text-sm";
            messageElement.innerHTML = `
            <div class="flex-shrink-0 mr-2 rounded bg-green-100 flex items-center justify-center w-8 h-8" title="${data.username}">
            <img src="https://github.com/${data.username}.png" alt="User Image" class="w-full h-full rounded">
        </div>
                <div class="flex-1 p-2 rounded bg-green-100">${data.message}</div>
            `;
        
            // Append the message element to the chatMessage element
            const chatMessageElement = document.getElementById("chatMessage");
            if (chatMessageElement) {
                chatMessageElement.appendChild(messageElement);
            }

            chatMessageElement.scrollTop = chatMessageElement.scrollHeight;
        };

        socket.onerror = function(event){
            console.log("WrSocket error observed", event)
        }
        console.log("Connected");
    } else {
        console.log("Not available the WEBSOCKET");
    }
}