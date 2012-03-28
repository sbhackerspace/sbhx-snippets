function WebSocketTest() {
    alert("WebSocketTest is running...");
    if ("WebSocket" in window) {
        // Google example code
        var ws = new WebSocket("ws://newsweetness.prototypemagic.com:8080/echo");
        ws.onopen = function() {
            // Web Socket is connected. You can send data by send() method
            alert("Let the streaming begin!!!");
        };
        ws.onmessage = function(evt) {
            var received_msg = evt.data;
            // alert(received_msg);
            var txt = document.createTextNode(received_msg);
            document.getElementById('ws_div').appendChild(txt);
        };
    } else {
        // the browser doesn't support WebSocket
        alert("No websocket support detected :-( Try Chrome 16 or 17?");
    }
}