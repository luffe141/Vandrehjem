import { useState } from "react";
import Calendar from "react-calendar";
import './Calender.css';

type ValuePiece = Date | null;

type Value = ValuePiece | [ValuePiece, ValuePiece];

function Kalender() {
  const [value, onChange] = useState<Value>(new Date());

  return (
    <div >
      <Calendar onChange={onChange} value={value} />
    </div>
  );
}
export default Kalender;
