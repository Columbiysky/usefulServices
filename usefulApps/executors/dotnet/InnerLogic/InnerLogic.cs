namespace InnerLogic
{
    public class InnerLogicService
    {
        public InnerLogicService() { }

        public void GetInfo(ServiceRequest serviceRequest)
        {
            Console.WriteLine("asfd");
        }
    }

    public record ServiceRequest
    {
        public required string CurrOne { get; set; }
        public required string CurrTwo { get; set; }
        public string? TimeFrame { get; set; }
        public string? Provider { get; set; }
    }
}
