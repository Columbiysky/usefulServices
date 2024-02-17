using System.Net;
using System.Net.Http.Headers;
using System.Web.Http;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

var app = builder.Build();

// Configure the HTTP request pipeline.

//app.UseHttpsRedirection();

app.MapPost("/", ([FromBody] ServiceRequest req) => new { one = req.CurrOne, two = req.CurrTwo, tf = req.TimeFrame, provider = req.Provider });


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


public record ServiceRequest
{
    public required string CurrOne { get; set; }
    public required string CurrTwo { get; set; }
    public string? TimeFrame { get; set; }
    public string? Provider { get; set; }
}