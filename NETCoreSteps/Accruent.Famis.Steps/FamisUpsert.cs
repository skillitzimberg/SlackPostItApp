using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using StepCore;

namespace Accruent.Famis.Steps
{
    public abstract class FamisUpsert : ServiceStep
    {
        [Input(Description = "a json formatted object that you would like to create or update")]
        public Dictionary<string, object> Object { get; set; }

        [Input(Description = "the endpoint of the entity you are trying to update")]
        public string Endpoint { get; set; }

        [Input(Description = "the endpoint of the entity you are trying to update")]
        public string IdField { get; set; }
    }
}