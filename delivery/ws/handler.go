package ws

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/service"
)

type Handler struct {
	upgrader    *websocket.Upgrader
	connections map[uint]*websocket.Conn
	services    *service.Service
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{
		upgrader:    &websocket.Upgrader{},
		connections: make(map[uint]*websocket.Conn),
		services:    services,
	}

	return h
}

func (h *Handler) Connect(c *gin.Context) {
	userid := c.GetUint("userid")

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	log.Println("Connected ", userid)

	h.connections[userid] = conn
	go h.readConnection(conn, userid)

	// TODO send unsent notifications
	//
	unsent, _ := h.services.Notification.GetUnsent(userid)
	for _, notif := range unsent {
		err = conn.WriteJSON(notif)
		if err != nil {
			log.Println("notif:", err)
			continue
		}
		h.services.Notification.Delete(&notif)
	}
}

func (h *Handler) readConnection(conn *websocket.Conn, userid uint) {
	for {
		log.Println("Reading from ", userid)

		var msg dto.ChatUpdate
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("read:", err)
			break
		}

		model, err := h.services.Chat.SaveMessage(msg, userid)
		if err != nil {
			log.Println("save:", err)
		}

		var recip uint
		if model.Connection.U1 == userid {
			recip = model.Connection.U2
		} else {
			recip = model.Connection.U1
		}
		err = h.SendMessage(model, recip)
		if err != nil {
			log.Println("send:", err)
		}
	}

	delete(h.connections, userid)
}

func (h *Handler) SendMessage(msg *domain.Message, recipient uint) error {
	not := domain.Notification{
		Type:      domain.TypeMessage,
		Recipient: recipient,
		Content:   msg,
	}

	recip, ok := h.connections[recipient]
	if ok {
		err := recip.WriteJSON(not)
		if err != nil {
			return err
		}
		return nil
	}

	h.services.Notification.Save(&not)

	return errors.New("recipient not connected")
}

func (h *Handler) SendFriendRequest(req *domain.MessageRequest, recipient uint) error {
	not := domain.Notification{
		Type:      domain.TypeMessage,
		Recipient: recipient,
		Content:   req,
	}

	recip, ok := h.connections[recipient]
	if ok {
		err := recip.WriteJSON(not)
		if err != nil {
			return err
		}
		return nil
	}

	h.services.Notification.Save(&not)

	return errors.New("recipient not connected")
}

// {"type":0,"content":"hiikfjhakjfh!!!!","connection":1}

//////////////// FOR TESTING PURPOSES ONLY

func Test(c *gin.Context) {
	userid64, _ := strconv.ParseUint(c.Param("userid"), 10, 32)
	userid := uint(userid64)
	_ = homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/ws/"+fmt.Sprint(userid))
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
