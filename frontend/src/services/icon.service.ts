export class IconService {
  // Store the loaded icons in a cache
  private iconCache: Record<string, any> = {};
  // Store promise map for pending icon loads to prevent duplicate requests
  private loadingIcons: Record<string, Promise<any>> = {};

  // Function to load an icon and cache it
  loadIcon = async (iconName: string): Promise<any> => {
    if (!iconName) return null;

  // If this icon is already being loaded, wait for that promise
  if (this.loadingIcons[iconName]) {
    return this.loadingIcons[iconName];
  }

  // If the icon is in cache, return it immediately
  if (this.iconCache[iconName]) {
    return this.iconCache[iconName];
  }

  // Otherwise, load the icon and cache it
  const iconPromise = import(`ionicons/icons/`)
    .then((module) => {
      const icon = module[iconName];
      this.iconCache[iconName] = icon;
      delete this.loadingIcons[iconName];
      return icon;
    })
    .catch(async (error) => {
      console.error(`Failed to load icon: ${iconName}`, error);
      // Load a default fallback icon
      return import("ionicons/icons").then((module) => {
        const icon = module.help;
        this.iconCache[iconName] = icon;
        delete this.loadingIcons[iconName];
        return icon;
      });
    });

  // Store the promise in the loading map
  this.loadingIcons[iconName] = iconPromise;
  return iconPromise;
};

  // Function to get an icon (will return immediately if cached, or null during loading)
  getIcon = async (iconName: string): Promise<any> => {
    // Return cached icon if available
    if (this.iconCache[iconName]) {
      return this.iconCache[iconName];
    }

    // If already loading, wait for that promise
    const existingPromise = this.loadingIcons[iconName];
    if (existingPromise) {
      return existingPromise;
    }

    // Start loading if not already loading
    return this.loadIcon(iconName);
  };
}