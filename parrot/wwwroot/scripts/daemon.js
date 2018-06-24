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

connection.on('clusterViewUpdated', (pod) => {
    console.log(pod);
    var li = document.createElement('li');
    var textNode = document.createTextNode(pod.name + ' namespace: (' + pod.nameSpace + ') ' + pod.action);
    li.appendChild(textNode);
    document.getElementById('pods').appendChild(li);
});

connection
    .start()
    .then(() => {
        console.log('connection started');
    })
    .catch(console.error);