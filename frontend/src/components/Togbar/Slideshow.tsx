import { useState } from 'react'
import styles from './Togbar.module.css'

const ImageSliderTre = ({ images }) => {
    const [currentImageIndex, setCurrentImageIndex] = useState(0);
  
    const nextSlide = () => {
      setCurrentImageIndex((prevIndex) => (prevIndex + 1) % images.length);
    };
  
    const prevSlide = () => {
      setCurrentImageIndex(
        (prevIndex) => (prevIndex - 1 + images.length) % images.length
      );
    };
  
    return (
      <div className={styles.imageSliderDiv}>
        <button className={styles.galleryBtn} onClick={prevSlide}>Previous</button>
        <div className={styles.galleryImgContainer}><img src={images[currentImageIndex]} alt={`Slide ${currentImageIndex + 1}`} /></div>
        <button className={styles.galleryBtn} onClick={nextSlide}>Next</button>
      </div>
    );
  };

export default ImageSliderTre