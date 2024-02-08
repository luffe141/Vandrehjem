import styles from './Omgivelser.module.css'
import ImageSlider from './ImageSlider';




const Omgivelser = () => {
  const images: string[]  = [
    'http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/_DAH7270.jpg',
    'http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/front2.jpg',
    'http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/slider-1.jpg',
  ];

  return (
    <div className={styles.OmgivelserDiv}>
      <h1>Image Slider App</h1>
      <ImageSlider images={images} />
    </div>
  );
};

export default Omgivelser;
