using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Newtonsoft.Json;
using parrot.Models;

namespace api.Controllers
{
    [Route("api/[controller]")]
    public class ClusterStatusController : Controller
    {   

        public ClusterStatusController(ILogger<ClusterStatusController> logger)
        {
            _logger = logger;
        }

        private readonly ILogger _logger;

        [HttpGet]
        public ActionResult Get()
        {
            return new OkResult();
        }

        [HttpPost]
        public ActionResult Post([FromBody]dynamic metadata)
        {
            string json = metadata.ToString();
            _logger.LogDebug("Incoming Cluster Update", json);

            Pod pod = new Pod();
            pod.Container = metadata.spec.containers[0].name;
            pod.ContainerImage = metadata.spec.containers[0].image;
            pod.CreationTimestamp = DateTime.Parse(metadata.creationTimestamp);
            pod.Name = metadata.name;
            pod.NameSpace = metadata.Namespace;
            pod.Status = metadata.status.phase;

            return new OkResult();
        }
    }
}
