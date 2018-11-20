new Vue({
    el: '#app',

    data: {
        ws: null,
        newMessage: '',
        userId: null
    },
    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            var element = document.createElement("p")
            element.textContent = msg.Message
            document.body.appendChild(element)
        });
    },
    methods: {
        send: function () {
            if (this.newMessage != '') {
                this.ws.send(
                    JSON.stringify({
                        'UserId': "2"
                    })
                );
                this.newMessage = '';
            }
        }
    }
});