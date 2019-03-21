using System.Threading.Tasks;
using Famis;

namespace Accruent.Famis.Steps.Update
{
    public class Update : FamisUpsert
    {
        public override Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            return service.UpdateRecord(Endpoint, Object, IdField);
        }
    }
}