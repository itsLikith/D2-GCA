// Default coordinate limits for the map canvas
export const MAP_BOUND = 1000; // -1000 to +1000 NM
export const CANVAS_SIZE = 500; // 500x500 px

export const toCanvasX = (nmX: number) => {
  return CANVAS_SIZE / 2 + (nmX / MAP_BOUND) * (CANVAS_SIZE / 2);
};

export const toCanvasY = (nmY: number) => {
  // Invert Y axis since canvas origin (0,0) is at top-left
  return CANVAS_SIZE / 2 - (nmY / MAP_BOUND) * (CANVAS_SIZE / 2);
};

export const fromCanvasX = (pixelX: number) => {
  return ((pixelX - CANVAS_SIZE / 2) / (CANVAS_SIZE / 2)) * MAP_BOUND;
};

export const fromCanvasY = (pixelY: number) => {
  return -((pixelY - CANVAS_SIZE / 2) / (CANVAS_SIZE / 2)) * MAP_BOUND;
};
