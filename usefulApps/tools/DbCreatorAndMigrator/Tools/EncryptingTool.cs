using Org.BouncyCastle.Crypto.Digests;
using System.Runtime.Serialization;

namespace DbCreatorAndMigrator.Helpers
{
    public static class EncryptingTool
    {
        public static string Encrypt(object obj)
        {
            var bytes = ToBytes(obj);

            if (bytes is null)
            {
                return string.Empty;
            }

            var keccak = new KeccakDigest(256);
            keccak.BlockUpdate(bytes);
            var hash = new byte[keccak.GetByteLength()];
            keccak.DoFinal(hash);

            var result = BitConverter.ToString(hash, 0, 32).Replace("-", "").ToLower();
            return result;
        }

        private static byte[]? ToBytes(object obj)
        {
            if (obj is null)
            {
                return null;
            }

            var bf = new DataContractSerializer(typeof(Object));

            using (var ms = new MemoryStream())
            {
                bf.WriteObject(ms, obj);
                return ms.ToArray();
            }
        }
    }
}
