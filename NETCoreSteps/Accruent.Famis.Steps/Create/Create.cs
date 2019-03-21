using System.Threading.Tasks;
using Famis;

namespace Accruent.Famis.Steps.Create
{
    public class Create : FamisUpsert
    {
        public override Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            
        }
    }
}