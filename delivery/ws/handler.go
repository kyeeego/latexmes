package ws

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/service"
)

type Handler struct {
	upgrader    *websocket.Upgrader
	connections map[uint]*websocket.Conn
	services    *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		upgrader:    &websocket.Upgrader{},
		connections: make(map[uint]*websocket.Conn),
	}
}

func (h *Handler) Connect(c *gin.Context) {
	userid := c.GetUint("userid")

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	h.connections[userid] = conn
	go h.readConnection(conn, userid)

	// TODO send unsent messages
	// unsent := h.services.Chat.GetUndeliveredMessages(userid)
	// h.services.Chat.SetMessagesAsSent(userid)
	// lol
}

func (h *Handler) readConnection(conn *websocket.Conn, userid uint) {
	for {
		var msg dto.ChatUpdate
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("read:", err)
			break
		}

		// status := domain.StatusUnsent
		recip, ok := h.connections[msg.Recipient]
		if ok {
			err := recip.WriteJSON(msg)
			if err != nil {
				log.Println("write:", err)
			} else {
				// status = domain.StatusSent
			}
		}

		// err = h.services.Chat.SaveMessage(msg, status)
		// if err != nil {
		// 	log.Println("save:", err)
		// }
	}
}

//////////////// FOR TESTING PURPOSES ONLY

func Test(c *gin.Context) {
	_ = homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/ws")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
