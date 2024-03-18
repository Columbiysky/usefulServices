namespace DbCreatorAndMigrator.BuiltinDataFiller
{
    public interface IBuiltInDataFillerBase
    {
        public abstract void Fill(UsefulServicesContext ops);
    }
}
