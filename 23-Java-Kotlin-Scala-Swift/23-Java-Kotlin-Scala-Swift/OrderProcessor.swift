struct Order { let id: String; let amount: Double }

actor OrderProcessor {
    // Swift Actors: Memory safety without GC pauses
    func process(order: Order) async {
        let converted = order.amount * 1.1
        print("Swift processed: \(order.id) -> \(converted)")
    }
}
