// Event Bus: A decoupled communication pattern
type EventCallback<T> = (data: T) => void | Promise<void>;

class EventBus {
    private listeners: Map<string, EventCallback<any>[]> = new Map();

    // Subscribe to an event with Type Safety
    public subscribe<T>(event: string, callback: EventCallback<T>): void {
        const currentListeners = this.listeners.get(event) || [];
        this.listeners.set(event, [...currentListeners, callback]);
    }

    // Publish an event
    public async publish<T>(event: string, data: T): Promise<void> {
        const callbacks = this.listeners.get(event);
        if (!callbacks) return;

        // Execute all subscribers concurrently
        await Promise.all(callbacks.map(cb => cb(data)));
    }
}

// Usage Example
interface UserCreatedEvent {
    userId: string;
    email: string;
}

const bus = new EventBus();

bus.subscribe<UserCreatedEvent>("user.created", (data) => {
    console.log(`Sending welcome email to: ${data.email}`);
});

// Business Logic
bus.publish<UserCreatedEvent>("user.created", { 
    userId: "123", 
    email: "architect@example.com" 
});
