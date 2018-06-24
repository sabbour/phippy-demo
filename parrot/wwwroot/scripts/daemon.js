const pods = [];

const connection = new signalR.HubConnectionBuilder()
    .withUrl('/daemonhub', { 
        logger: signalR.LogLevel.Verbose 
    })
    .build();

connection.onclose(() => console.log('closed'));

connection.on('clusterViewUpdated', (pod) => {
    console.log(pod);
    var li = document.createElement('li');
    var textNode = document.createTextNode(pod.name + ' namespace: (' + pod.nameSpace + ') ' + pod.action);
    li.appendChild(textNode);
    document.getElementById('pods').appendChild(li);

    var source = $('#podCardTemplate').html();
    var template = Handlebars.compile(source);
    var result = template(pod);
    console.log(result);
    $('#podcards').append(result);
});

connection
    .start()
    .then(() => {
        console.log('connection started');
    })
    .catch(console.error);