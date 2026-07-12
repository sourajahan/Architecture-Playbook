public record Order(String id, double amount) {}

public class OrderProcessor {
    public void process(Order order) {
        // Java 21+ Virtual Threads: Lightweight concurrency
        Thread.ofVirtual().start(() -> {
            double converted = order.amount() * 1.1; 
            System.out.println("Java processed: " + order.id() + " -> " + converted);
        });
    }
}
