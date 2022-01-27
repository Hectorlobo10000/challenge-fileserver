const socket = io()

new Vue({
    el: "#app",
    vuetify: new Vuetify({
        icons: {
            iconfont: 'mdi'
        }
    }),
    data: {
        headers: [
            {text: 'Id', value: 'id', align: 'center'},
            {text: 'User', value: 'user', align: 'center'},
            {text: 'Command', value: 'command', align: 'center'},
            {text: 'Filename', value: 'filename', align: 'center'},
            {text: 'Extension', value: 'extension', align: 'center'},
        ],
        messages: [],
        files: [],
        model: 1
    },
    created() {
        socket.on('message', (msg) => {
            let content = JSON.parse(msg)

            this.messages.push({
                id: content.id,
                user: content.user,
                command: content.command,
                filename: content.filename,
                extension: content.extension
            })

            this.files.push({
                id: content.id,
                user: content.user,
                extension: content.extension,
                contentType: content.contentType,
                filename: content.filename,
                file: content.file
            })
        });
    },
    methods: {
        downloadFile: function (id, event) {
            event.preventDefault()

            let vm = this
            let object = vm.files.find(e => e.id === id)

            try {
                var blob = b64toBlob(object.file, object.contentType);
                var blobUrl = URL.createObjectURL(blob);
                saveAs(blobUrl, object.filename)
            } catch (e) {
                console.log(e)
            }

        }
    }
})