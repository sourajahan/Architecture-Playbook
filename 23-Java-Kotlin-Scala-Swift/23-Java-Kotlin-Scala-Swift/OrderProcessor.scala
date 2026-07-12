case class Order(id: String, amount: Double)

object OrderProcessor {
  // Scala IO: Referential transparency & Functional purity
  def processOrder(order: Order): IO[Unit] = 
    IO {
      val converted = order.amount * 1.1
      println(s"Scala processed: ${order.id} -> $converted")
    }
}
