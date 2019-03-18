using System;
using System.Threading.Tasks;
using Famis;
using Famis.Model;
using StepCore;

namespace Accruent.Famis.Steps.Fetch {
    
    [StepDescription("fetch_user", Description = "Fetches a user from famis")]
    public class FetchUser : ServiceStep {
        [Input(Description = "The filter to use when fetching a user")]
        public string Filter { get; set; }
        
        [Output(Description = "The user that was fetched from FAMIS")]
        public User User { get; set; }

        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            var response = await service.GetUsers("ExternalId eq '1'", limit: 1);
            if (response.ResultCount == 1) {
                User = response.Current;
            }
        }
    }
}