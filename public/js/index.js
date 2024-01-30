
Window.onload = ()=>{
    var conn
var messageInputField = document.getElementById("message-user")
document.getElementById("chatForm").onsubmit = ()=>{
    if (!conn){
        return false
    }if (!messageInputField.value){
        return false
    }
    conn.send(messageInputField.value)
}
    if(window["WebSocket"]){
        console.log("WebSocket Started")
        conn =  new WebSocket(document.location.host + "/ws");
        

    }else{
        console.log("Not available the WEBSOCKET")
    }
}