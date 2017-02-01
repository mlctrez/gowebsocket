define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {


		var loc = window.location, new_uri;
		if (loc.protocol === "https:") {
		    new_uri = "wss:";
		} else {
		    new_uri = "ws:";
		}
		new_uri += "//" + loc.host + "/entry";
		
		var ws = new WebSocket(new_uri);
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);
