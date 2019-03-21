using System.Data;
using System.Threading.Tasks;
using Famis;

namespace Accruent.Famis.Steps.Create
{
    public class Create : FamisUpsert
    {
        public override Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            return service.CreateRecord(Endpoint, Object);
        }
    }
}