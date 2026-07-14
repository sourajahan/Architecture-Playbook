// Principal-grade implementation of a Feature Detector.
// We avoid direct usage to ensure system stability.

type Capability = 'web-share' | 'intersection-observer' | 'service-worker';

export class FeatureDetector {
  // Defensive check: Centralized registry of supported features
  private static readonly registry: Record<Capability, () => boolean> = {
    'web-share': () => typeof navigator.share === 'function',
    'intersection-observer': () => 'IntersectionObserver' in window,
    'service-worker': () => 'serviceWorker' in navigator,
  };

  /**
   * Executes code only if the capability is supported.
   * This is the "Principal" way to handle the evolving Web Standard.
   */
  static runIfSupported(capability: Capability, callback: () => void, fallback?: () => void) {
    if (this.registry[capability]()) {
      callback();
    } else if (fallback) {
      fallback();
    } else {
      console.warn(`Feature ${capability} not supported on this platform.`);
    }
  }
}

// Usage Example:
FeatureDetector.runIfSupported('web-share', 
  () => navigator.share({ title: 'System Architecture' }),
  () => console.log('Falling back to manual clipboard copy...')
);
