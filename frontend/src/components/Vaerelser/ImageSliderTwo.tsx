import { useState } from 'react'
import styles from './Vaerelser.module.css'

const ImageSliderTwo = ({ images }) => {
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
        <div className={styles.galleryImgContainer}>
        <img className={styles.galleryBtn} onClick={nextSlide} src="Images/down-arrow.png" alt="arrow" />
        <img id={styles.galleryImg} src={images[currentImageIndex]} alt={`Slide ${currentImageIndex + 1}`} />
        <img className={styles.galleryBtn} onClick={prevSlide} src="Images/down-arrow.png" alt="arrow" />
        </div>
      </div>
    );
  };

export default ImageSliderTwo