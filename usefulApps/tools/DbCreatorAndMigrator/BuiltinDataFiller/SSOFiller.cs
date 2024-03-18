using DbCreatorAndMigrator.Helpers;
using DbCreatorAndMigrator.Models.Sso;

namespace DbCreatorAndMigrator.BuiltinDataFiller
{
    public class SSOFiller : IBuiltInDataFillerBase
    {
        public void Fill(UsefulServicesContext context)
        {
            var pass = EncryptingTool.Encrypt("1adminP@ssword!");
            if (context.Accounts.FirstOrDefault(x => x.Id == 1) is null)
            {
                context.Accounts.Add(new Account { Id = 0, Login = "root", Password = pass });
                context.SaveChanges();
            }
        }
    }
}
