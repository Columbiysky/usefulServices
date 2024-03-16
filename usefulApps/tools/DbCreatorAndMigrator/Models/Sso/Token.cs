namespace DbCreatorAndMigrator.Models.Sso
{
    public class Token
    {
        public long Id { get; set; }
        public required string TokenValue { get; set; }
        public DateTimeOffset LastActivityTime { get; set; }
    }
}
