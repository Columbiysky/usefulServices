using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;
using Microsoft.Extensions.Configuration;

namespace DbCreatorAndMigrator
{
    internal class ContextFactory : IDesignTimeDbContextFactory<UsefulServicesContext>
    {
        public UsefulServicesContext CreateDbContext(string[] args)
        {

            var builder = new ConfigurationBuilder();
            // установка пути к текущему каталогу
            builder.SetBasePath(Directory.GetCurrentDirectory());
            // получаем конфигурацию из файла appsettings.json
            builder.AddJsonFile("config.json");
            // создаем конфигурацию
            var config = builder.Build();
            // получаем строку подключения
            var connectionString = config.GetConnectionString("Database");

            var optionsBuilder = new DbContextOptionsBuilder<UsefulServicesContext>();
            var options = optionsBuilder.UseNpgsql(connectionString).Options;

            return new UsefulServicesContext(optionsBuilder.Options);
        }
    }
}
