using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using parrot.Models;

namespace parrot
{
    public class DaemonHub : Hub
    {
        static int _counter = 0;

        public override Task OnConnectedAsync()
        {
            _counter += 1;
            updateUserCount();
            return base.OnConnectedAsync();
        }

        public override Task OnDisconnectedAsync(System.Exception exception)
        {
            _counter -= 1;
            updateUserCount();
            return base.OnDisconnectedAsync(exception);
        }

        public async void updateUserCount()
        {
            await Clients.All.SendAsync("userCountUpdated", _counter);
        }

        public async void updateClusterView(Pod pod)
        {
            await Clients.All.SendAsync("clusterViewUpdated", pod);
        }
    }
}