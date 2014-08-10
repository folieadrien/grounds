// FIXME: Alert if browser doesn't support websockets
// FIXME: Do something if connection to websockets is impossible
// FIXME: readyState3
function Client(endpoint) {
  this.endpoint = endpoint;
  this.socket = null;
}

Client.prototype.connect = function() {
  this.socket = new WebSocket(this.endpoint);
  this.bindEvents();
}

Client.prototype.bindEvents = function() {
  this.socket.onmessage = function(event) {
    if (event.data.length) {
      response = JSON.parse(event.data);
      if (response.stream === 'status') {
        response.chunk = "\n[Program exited with status: " + response.chunk + "]";
      }
      if (response.stream === 'executing') {
        $("#console").find("span").each(function() {
          this.remove();
        });
      }
      $("#console").append($('<span class="'+ response.stream +'">').text(response.chunk));
    }
  };
  var that = this;
  this.socket.onclose = function() {
    that.socket = null;
  };
}

Client.prototype.send = function(data) {
  if (this.socket === null) {
    this.connect();
  }
  var that = this;
  setTimeout(function(){
    if (that.socket !== null && that.socket.readyState === 1) {
      that.socket.send(data);
      return;
    } else {
      that.send(data); 
    }
  }, 1);
};
