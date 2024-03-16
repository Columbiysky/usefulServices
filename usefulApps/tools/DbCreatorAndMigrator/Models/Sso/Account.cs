namespace DbCreatorAndMigrator.Models.Sso
{
    public class Account
    {
        public long Id { get; set; }
        public required string Login { get; set; }
        public required string Password { get; set; }
    }
}
