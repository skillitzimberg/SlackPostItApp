using System;
using System.Threading.Tasks;
using Famis;
using Famis.Model;
using Newtonsoft.Json;
using StepCore;

namespace Accruent.Famis.Steps.Create {
    public abstract class UpsertUserBase : ServiceStep {
        [Input(Description = "User to create or update")]
        public User User { get; set; }
        
        [Output(Description = "True if the user was created successfully")]
        public bool Success { get; set; }
        [Output(Description = "The newly created user. Will contain the ID of the user.", Key = "User")]
        public User UserResult { get; set; }
        [Output(Description = "Contains the response message from the FAMIS services")]
        public string Message { get; set; }
    }
    

    [StepDescription("create_user")]
    public class CreateUser : UpsertUserBase {
       
        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            var result = await service.PostUser(User);
            Message = result.Message;
            UserResult = result.Object;
            Success = result.Success;
        }
    }

    [StepDescription("create_user_mock")]
    public class CreateUserMock : UpsertUserBase {
        public override Task ExecuteAsync() {
            Success = true;
            UserResult = User;
            Console.WriteLine(JsonConvert.SerializeObject(User));
            User.Id = 12345;
            Message = "User created successfully";
            return Task.CompletedTask;
        }
    }
}