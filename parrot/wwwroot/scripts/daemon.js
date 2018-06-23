const connection = new signalR.HubConnectionBuilder()
    .withUrl('/daemonhub', { 
        logger: signalR.LogLevel.Verbose 
    })
    .build();

connection.onclose(() => console.log('closed'));

connection.on("userCountUpdated", (userCount) => { 
    const userCountUi = document.getElementById("userCount");
    userCountUi.textContent = userCount;
});

connection
    .start()
    .then(() => {
        console.log('connection started');
    })
    .catch(console.error);