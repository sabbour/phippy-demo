using System;

namespace parrot.Models
{
    public class Pod
    {
        public string Name { get; set; }
        public string Container { get; set; }
        public string NameSpace { get; set; }
        public string ContainerImage { get; set; }
        public string Status { get; set; }
        public string Action { get; set; }
    }
}