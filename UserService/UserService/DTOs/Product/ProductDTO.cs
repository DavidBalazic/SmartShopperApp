namespace UserService.DTOs.Product
{
    public class ProductDTO
    {
        public string Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public double Quantity { get; set; }
        public double Price { get; set; }
        public string Store { get; set; }
        public string Unit { get; set; }
        public double PricePerUnit { get; set; }
    }
}
