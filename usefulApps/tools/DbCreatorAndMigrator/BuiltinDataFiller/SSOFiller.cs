namespace DbCreatorAndMigrator.BuiltinDataFiller
{
    public class SSOFiller : IBuiltInDataFillerBase
    {
        public void Fill()
        {
            // todo: fix getting config to db
            //using (var context = new UsefulServicesContext())
            //{
            //    if (context.Accounts.FirstOrDefault(x => x.Id == 1) is null)
            //    {
            //        context.Accounts.Add(new Account { Id = 0, Login = "root", Password = "asd" });
            //        context.SaveChanges();
            //    }
            //}
        }
    }
}
