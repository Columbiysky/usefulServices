namespace DbCreatorAndMigrator.Models.Sso
{
    public class AccountToken
    {
        public long Id { get; set; }
        public required Account Account { get; set; }
        public required Token Token { get; set; }
    }
}
