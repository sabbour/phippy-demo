﻿using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using Newtonsoft.Json;
using parrot;
using parrot.Models;

namespace api.Controllers
{
    [Route("api/[controller]")]
    public class ClusterStatusController : Controller
    {   

        public ClusterStatusController(ILogger<ClusterStatusController> logger, DaemonHub hub)
        {
            _hub = hub;
            _logger = logger;
        }

        private DaemonHub _hub;
        private readonly ILogger _logger;

        [HttpGet]
        public ActionResult Get()
        {
            return new OkResult();
        }

        [HttpPost]
        public ActionResult Post([FromBody]Pod pod)
        {
            _logger.LogDebug("Incoming Cluster Update", pod);
            _hub.updateClusterView(pod);

            return new OkResult();
        }
    }
}
