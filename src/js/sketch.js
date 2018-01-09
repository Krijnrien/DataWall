/* eslint-disable no-param-reassign */

import 'p5/lib/addons/p5.sound';
import 'p5/lib/addons/p5.dom';

import { initializeService } from './modules/service/DeviceService';
import VisualizationManager from './modules/visualization/VisualizationManager';
import Utillities from './modules/util/Utillities';

const manager = new VisualizationManager();

window.manager = manager; // for debugging purposes.

initializeService((sender, args) => {
  manager.currentData = Utillities.normalizeData(args.data);
  manager.previousData = Utillities.normalizeData(args.previousData);
});

function setupEventListeners() {
  const animations = ['tron', 'dots', 'drops', 'vertex'];

  animations.forEach((animation) => {
    document.getElementById(animation).onclick = () => manager.find(animation);
  });
}

const sketch = (p5) => {
  window.p5 = p5;

  p5.setup = () => {
    const canvas = p5.createCanvas(p5.windowWidth, p5.windowHeight);
    canvas.parent('sketch');
    p5.frameRate(24);

    setupEventListeners();
  };

  p5.draw = () => {
    manager.current.update();
    manager.current.show();
  };
};

export default sketch;
