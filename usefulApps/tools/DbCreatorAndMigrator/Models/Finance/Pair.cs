namespace DbCreatorAndMigrator.Models.Finance
{
    public class Pair
    {
        public long Id { get; set; }
        public required string Name { get; set; }
        public double Price { get; set; }
    }
}
