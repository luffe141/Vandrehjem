import React from "react";
import { useEffect } from 'react';


const page = () => {

    useEffect(() => {
        const initMap = () => {
          new google.maps.Map(document.getElementById('map'), {
            center: {lat: -34.397, lng: 150.644},
            zoom: 8
          });
        };
        const script = document.createElement('script');
        script.src = `https://maps.googleapis.com/maps/api/js?key=YOUR_API_KEY&callback=initMap`;
        script.async = true;
        script.defer = true;
        document.body.appendChild(script);


  return (
    <div>
      <h1>Kontakt</h1>
      <p>
        Hvis du har et spørgsmål eller kan vi på nogen anden måde kan hjælpe
        dig, så tøv ikke med at tage fat i os, i formularen her under kan du
        indtaste dine oplysninger og så vil vi svare dig hurtigst muligt, du er
        også altid velkommen til at ringe på 40224199
      </p>
      <div>
        <input type="text" name="" id="" placeholder="Navn" />
        <input type="text" name="" id="" placeholder="Email" />
        <input type="text" name="" id="" placeholder="Tlf" />
        <input type="text" name="" id="" placeholder="Emne" />
        <textarea name="" id="" cols="30" rows="10" placeholder="Besked"></textarea>
        <button>Send</button>
      </div>
      <div>
       <p>maps</p>
        
      </div>
    </div>
  );
};

export default page;
