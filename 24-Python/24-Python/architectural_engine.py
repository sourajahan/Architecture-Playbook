import asyncio
import signal
import logging
import uuid
from abc import ABC, abstractmethod
from typing import List

# Setup logging for observability
logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] [%(name)s] %(message)s")
logger = logging.getLogger("Engine")

# Interface pattern for interchangeable architecture (Dependency Inversion)
class DataProcessor(ABC):
    @abstractmethod
    async def run(self, task_id: str) -> None:
        pass

# Business logic implementation
class FinancialProcessor(DataProcessor):
    async def run(self, task_id: str) -> None:
        logger.info(f"Processing financial task: {task_id}")
        await asyncio.sleep(1) # Simulate heavy workload
        logger.info(f"Task {task_id} completed.")

# Main orchestrator (The system heart)
class Engine:
    def __init__(self, processor: DataProcessor):
        self._processor = processor
        self._running = True

    async def start(self):
        logger.info("Engine started. Waiting for tasks...")
        while self._running:
            task_id = str(uuid.uuid4())[:8]
            await self._processor.run(task_id)
            await asyncio.sleep(2)

    def stop(self):
        logger.info("Shutdown signal received. Cleaning up...")
        self._running = False

# Lifecycle Management
async def main():
    processor = FinancialProcessor()
    engine = Engine(processor)
    
    # Setup system signals for graceful shutdown
    loop = asyncio.get_running_loop()
    for sig in (signal.SIGINT, signal.SIGTERM):
        loop.add_signal_handler(sig, engine.stop)

    try:
        await engine.start()
    except Exception as e:
        logger.error(f"Engine crash: {e}")
    finally:
        logger.info("Engine exited gracefully.")

if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        pass
