import { useState } from 'react';
import styles from './Omgivelser.module.css'
import ImageSlider from './ImageSlider';




const Omgivelser = () => {
  const images = [
    'https://picsum.photos/id/237/300/300',
    'https://picsum.photos/id/49/300/300',
    'https://picsum.photos/id/159/300/300',
  ];

  return (
    <div className={styles.OmgivelserDiv}>
      <h1>Image Slider App</h1>
      <ImageSlider images={images} />
    </div>
  );
};

export default Omgivelser;
