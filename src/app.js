import P5 from 'p5';
import sketch from './js/sketch';
import './styles/styles.scss';
import VisualizationManager from './js/modules/visualization/VisualizationManager';

// Initialize sketch
new P5(sketch); //eslint-disable-line

document.getElementById('tron').onclick = () => {
  VisualizationManager.next();
  console.log(VisualizationManager.currentIndex);
};

// TODO this doesn't work
document.getElementById('tron').onclick = () => {
  VisualizationManager.next();
  console.log(VisualizationManager.currentIndex);
};
