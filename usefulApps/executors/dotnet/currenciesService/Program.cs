using System.Net;
using System.Net.Http.Headers;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

var app = builder.Build();

// Configure the HTTP request pipeline.

//app.UseHttpsRedirection();

app.MapGet("/", () => new { Test = "Test" });


app.Use(async (context, next) =>
{
    var tokenArr = context.Request.Headers.Authorization.ToString()?.Split(" ");
    if (tokenArr is null || tokenArr.Length < 2)
    {
        context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
        await context.Response.CompleteAsync();
    }
    else
    {
        var token = tokenArr[1];

        if (token != "CI_TEST")
        {
            var client = new HttpClient();
            var req = new HttpRequestMessage(HttpMethod.Get, "http://127.0.0.1:8081/checkToken");
            req.Headers.Authorization = new AuthenticationHeaderValue("Bearer", token);
            var response = client.Send(req);
            if (response is not null && response.StatusCode == HttpStatusCode.OK)
            {
                await next.Invoke();
            }
            else
            {
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                await context.Response.CompleteAsync();
            }
        }
    }
});


app.Run();
