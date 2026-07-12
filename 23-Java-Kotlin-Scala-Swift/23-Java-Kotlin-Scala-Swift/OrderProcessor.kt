data class Order(val id: String, val amount: Double)

class OrderProcessor {
    // Kotlin Coroutines: Suspension-based concurrency
    suspend fun processOrder(order: Order) {
        val converted = order.amount * 1.1
        println("Kotlin processed: ${order.id} -> $converted")
    }
}
