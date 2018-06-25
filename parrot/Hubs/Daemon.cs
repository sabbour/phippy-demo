using System.Threading.Tasks;
using System.Collections;
using System.Linq;
using System.Collections.Generic;
using Microsoft.AspNetCore.SignalR;
using parrot.Models;

namespace parrot
{
    public class DaemonHub : Hub
    {
        static List<Pod> Pods { get; set; }
        static List<string> DeletedPods { get; set; }

        static DaemonHub()
        {
            Pods = new List<Pod>();
            DeletedPods = new List<string>();
        }

        const string POD_DELETED_STATUS = "Deleted";

        public override Task OnConnectedAsync()
        {
            Clients.All.SendAsync("clusterViewUpdated", Pods);
            return base.OnConnectedAsync();
        }

        public void AddPod(Pod pod)
        {
            if(!DeletedPods.Contains(pod.Name))
            {
                Pods.Add(pod);
            }
        }

        public void RemovePod(Pod pod)
        {
            Pods.Remove(Pods.First(x => x.Container == pod.Container));
            DeletedPods.Add(pod.Name);
        }

        public void UpdatePod(Pod pod)
        {
            Pods.First(x => x.Container == pod.Container).Name = pod.Name;
            Pods.First(x => x.Container == pod.Container).NameSpace = pod.NameSpace;
            Pods.First(x => x.Container == pod.Container).Status = pod.Status;
        }

        public void updateClusterView(Pod pod)
        {
            pod.ContainerImage = pod.ContainerImage.Substring(0, pod.ContainerImage.IndexOf(':'));

            if (Pods.Any(x => x.Container == pod.Container))
                if (pod.Action == POD_DELETED_STATUS)
                    RemovePod(pod);
                else
                    UpdatePod(pod);
            else
                AddPod(pod);

            Clients.All.SendAsync("clusterViewUpdated", Pods);
        }
    }
}