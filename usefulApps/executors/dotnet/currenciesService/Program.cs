var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

var app = builder.Build();

// Configure the HTTP request pipeline.

//app.UseHttpsRedirection();

app.MapGet("/", () => new { Test = "Test" });

app.Use(async (context, next) =>
{
    await next.Invoke();
});


app.Run();
