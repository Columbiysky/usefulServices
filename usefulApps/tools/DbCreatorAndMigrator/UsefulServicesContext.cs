using DbCreatorAndMigrator.Models.Finance;
using DbCreatorAndMigrator.Models.Sso;
using Microsoft.EntityFrameworkCore;

namespace DbCreatorAndMigrator
{
    public class UsefulServicesContext : DbContext
    {


        public UsefulServicesContext(DbContextOptions<UsefulServicesContext> options) : base(options)
        {
        }

        public virtual DbSet<Account> Accounts { get; set; }
        public virtual DbSet<Token> Tokens { get; set; }
        public virtual DbSet<AccountToken> AccountsTokens { get; set; }
        public virtual DbSet<Pair> Pair { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Account>().ToTable(nameof(Accounts), "sso");
            modelBuilder.Entity<Token>().ToTable(nameof(Tokens), "sso");
            modelBuilder.Entity<AccountToken>().ToTable(nameof(AccountsTokens), "sso");
            modelBuilder.Entity<AccountToken>().HasOne(a => a.Account).WithMany();
            modelBuilder.Entity<AccountToken>().HasOne(a => a.Token).WithMany();

            modelBuilder.Entity<Pair>().ToTable(nameof(Pair), "finance");
        }
    }
}
