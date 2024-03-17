using DbCreatorAndMigrator.BuiltinDataFiller;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace DbCreatorAndMigrator
{
    internal class Program
    {
        private static ServiceCollection Services { get => InitServices(); }

        static void Main(string[] args)
        {

            var builder = new ConfigurationBuilder();
            builder.SetBasePath(Directory.GetCurrentDirectory());
            builder.AddJsonFile("config.json");
            var config = builder.Build();
            var connectionString = config.GetConnectionString("Database");

            var optionsBuilder = new DbContextOptionsBuilder<UsefulServicesContext>();
            var options = optionsBuilder.UseNpgsql(connectionString).Options;

            using (var context = new UsefulServicesContext(options))
            {
                Console.WriteLine("Migrating started...");
                context.Database.Migrate();
                Console.WriteLine("Migrating completed!");

                Console.WriteLine(string.Empty);

                Console.WriteLine("Fillin' buiiltin started...");
                FillBuiltInData();
                Console.WriteLine("Fillin' buiiltin completed!");
            }
        }

        static private void FillBuiltInData()
        {
            Services.Scan(s => s.FromAssemblyOf<Program>()
                .AddClasses(c => c.AssignableTo(typeof(IBuiltInDataFillerBase)))
                .AsImplementedInterfaces()
                .WithSingletonLifetime());

            var serviceProvider = Services.BuildServiceProvider();
            using var scope = serviceProvider.CreateScope();
            var endpoints = scope.ServiceProvider.GetServices<IBuiltInDataFillerBase>().ToList();
            endpoints.ForEach(e => e.Fill());
        }

        private static ServiceCollection InitServices()
        {
            var services = new ServiceCollection();
            services.AddSingleton<IBuiltInDataFillerBase, SSOFiller>();

            return services;
        }
    }
}