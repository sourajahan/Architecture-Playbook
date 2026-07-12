using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;

namespace Enterprise.Architecture;

// 1. The Contract (Interface Segregation)
public interface IOrderService
{
    Task ProcessOrderAsync(Guid orderId);
}

// 2. The Implementation (Domain Logic)
public class OrderService : IOrderService
{
    private readonly ILogger<OrderService> _logger;
    
    public OrderService(ILogger<OrderService> logger)
    {
        _logger = logger;
    }

    public async Task ProcessOrderAsync(Guid orderId)
    {
        _logger.LogInformation("Processing order: {OrderId}", orderId);
        await Task.Delay(100); // Simulate processing
    }
}

// 3. DI Container Setup (Infrastructure/Startup)
public class Startup
{
    public void ConfigureServices(IServiceCollection services)
    {
        // Singleton, Scoped, or Transient depending on the architectural lifecycle
        services.AddScoped<IOrderService, OrderService>();
        services.AddLogging();
    }
}
