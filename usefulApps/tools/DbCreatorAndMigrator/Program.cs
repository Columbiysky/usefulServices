using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;

namespace DbCreatorAndMigrator
{
    internal class Program
    {
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
            }
        }
    }
}